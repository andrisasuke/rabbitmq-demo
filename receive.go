package main

import (
	"log"
	"rabbitmq-demo/util"
	"rabbitmq-demo/model"
	"github.com/streadway/amqp"
	"encoding/json"
)


func main() {
	ReadConfig()
	conn, err := amqp.Dial(Config.RabbitMQUri)
	util.FailOnError(err, "Failed to connect to rabbit_mq")
	defer conn.Close()

	ch, err := conn.Channel()
	util.FailOnError(err, "Failed to open rabbit_mq channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		Config.EmailQueue, // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	util.FailOnError(err, "Failed to declare a email queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	util.FailOnError(err, "Failed to register a consumer")

	receiver := make(chan bool, 10)
	mailSender := make(chan string, 10)
	go func() {
		for d := range msgs {
			log.Printf("Received a message from rabbit_mq: %s", string(d.Body))
			mailSender <- string(d.Body)
		}
	}()

	go func() {
		for d := range mailSender {
			log.Println("consume message ", d)
			var email model.Email
			err := json.Unmarshal([]byte(d), &email)
			if err != nil {
				log.Println("Error while decoding message ", err)
			} else {
				log.Printf("Sending a email message: %s", email)
				er, _ := util.Send(email.To, email.Subject, email.Body, Config.SmtpUserAccount,
				 		   Config.SmtpPasswordAccount, Config.SmtpHost, Config.SmtpAddress )
				if er != nil {
					log.Printf("Failed to sending email message: %s", er)
				} else {
					log.Printf("Message has been sent to: %s", email.To)
				}
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To shutdown email receiver press CTRL+C")
	<-receiver
}

