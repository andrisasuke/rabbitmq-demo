package client

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"rabbitmq-demo/configs"
	"rabbitmq-demo/model"
	"rabbitmq-demo/util"
	"time"

	"github.com/streadway/amqp"
)

type amqConsumer struct {
	queueName  string
	consumerID string
	connection *amqp.Connection
	channel    *amqp.Channel
}

func NewAMQConsumer() {
	connection, e := amqp.Dial(configs.Config.RabbitMQUri)
	if e != nil {
		panic(e)
	}
	amqConsumer := &amqConsumer{
		connection: connection,
		queueName:  configs.Config.EmailQueue,
		consumerID: getConsumerID(configs.Config.EmailQueue),
	}
	doneChannel := make(chan bool)
	errorChan := amqConsumer.connection.NotifyClose(make(chan *amqp.Error))
	amqConsumer.WatchConnection(errorChan)
	if e := amqConsumer.Consume(); e != nil {
		fmt.Println("unable to start consumer")
		return
	}
	<-doneChannel
}

func (a *amqConsumer) WatchConnection(errorChan chan *amqp.Error) {
	// notify if connection is closed then try to reconnect
	go func() {
		for {
			reason, ok := <-errorChan
			if !ok {
				break
			}
			fmt.Printf("connection is closed :%d reason: %s\n", reason.Code, reason.Reason)
			for {
				// wait 5s for reconnect
				time.Sleep(5 * time.Second)
				connection, e := amqp.Dial(configs.Config.RabbitMQUri)
				if e != nil {
					fmt.Printf("unable reconnect to rabbitmq :%s\n", e.Error())
				} else {
					a.connection = connection
					a.Reconnect(a.connection)
					fmt.Println("reconnect success")
					break
				}
			}
		}
	}()
}

func (a *amqConsumer) Consume() error {
	fmt.Println("Starting amq consumer")
	ch, e := a.connection.Channel()
	if e != nil {
		fmt.Printf("unable to create amq channel :%s\n", e.Error())
		return e
	}

	a.channel = ch

	// Create the queue if not exist
	if _, e := ch.QueueDeclare(a.queueName, true, false, false, false, nil); e != nil {
		fmt.Printf("unable to create queue %s :%s\n", a.queueName, e.Error())
		return e
	}

	// notify if channel is closed, then recreate it
	go func() {
		for {
			reason, ok := <-ch.NotifyClose(make(chan *amqp.Error))
			if !ok {
				if ch != nil {
					ch.Close()
				}
				break
			}
			fmt.Printf("channel is closed :%d reason: %s\n", reason.Code, reason.Reason)
			for {
				// wait 3s for connection reconnect
				time.Sleep(3 * time.Second)
				newChannel, err := a.connection.Channel()
				if err == nil {
					fmt.Println("channel recreate success")
					a.channel = newChannel
					err = a.CreateConsumer()
					if err == nil {
						break
					}
				}
				fmt.Printf("unable recreate channel reason: %s\n", err.Error())
			}
		}

	}()
	return a.CreateConsumer()
}

func (a *amqConsumer) Reconnect(conn *amqp.Connection) {
	fmt.Println("create new connection for consumer channel")
	a.connection = conn
}

func (a *amqConsumer) CreateConsumer() error {
	msgs, err := a.channel.Consume(
		configs.Config.EmailQueue, // queue
		"",                        // consumer
		false,                     // auto-ack
		false,                     // exclusive
		false,                     // no-local
		false,                     // no-wait
		nil,                       // args
	)
	if err != nil {
		fmt.Printf("unable to consume messages from channel :%s\n", err.Error())
		return err
	}

	go func(msgs <-chan amqp.Delivery) {
		for msg := range msgs {
			fmt.Println("consuming message email")
			defer func(message *amqp.Delivery) {
				if p := recover(); p != nil {
					fmt.Printf("requeuing message due to recover() panic, %v\n", p)
					message.Reject(true)
				}
			}(&msg)

			var email model.Email
			err := json.Unmarshal([]byte(msg.Body), &email)
			if err != nil {
				fmt.Println("Error while decoding message ", err)
			} else {
				e := util.Send(email.To, email.Subject, email.Body, "sender@example.com")
				if e != nil {
					fmt.Printf("Message is failed to sent to: %s, %s\n", email.To, err.Error())
				} else {
					// doing actual send email
					fmt.Printf("Message has been sent to: %s with body %s\n", email.To, email.Body)
				}
			}
			msg.Ack(false)
		}
	}(msgs)
	return nil
}

func getConsumerID(queue string) string {
	var hostname string
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "localhost"
	}
	pid := os.Getpid()
	return fmt.Sprintf("%s#%d@%s/%d", queue, pid, hostname, rand.Uint32())
}
