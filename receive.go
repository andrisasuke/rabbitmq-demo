package main

import (
	"log"
	"rabbitmq-demo/util"
	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://admin:admin@localhost:5672/")
	util.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	util.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"email_q", // name
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
			log.Printf("Sending a email message: %s", d)
			//todo send email action
		}
	}()

	log.Printf(" [*] Waiting for messages. To shutdown email receiver press CTRL+C")
	<-receiver
	<-mailSender
}

