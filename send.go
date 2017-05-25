package main

import (
	"log"
	"github.com/streadway/amqp"
	"rabbitmq-demo/util"
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

	body := "{ 'sender': 'receiver@gmail.com', 'subject': 'Your_subject'  'body' : 'this is the message body'}"
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(body),
		})
	log.Printf(" [x] Sent %s", body)
	util.FailOnError(err, "Failed to publish a message")
}