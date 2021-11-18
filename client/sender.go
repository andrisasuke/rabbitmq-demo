package client

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"rabbitmq-demo/configs"
	"rabbitmq-demo/model"
	"rabbitmq-demo/util"

	"github.com/streadway/amqp"
)

func Send() {
	conn, err := amqp.Dial(configs.Config.RabbitMQUri)
	util.FailOnError(err, "Failed to connect to rabbit_mq")
	defer conn.Close()

	ch, err := conn.Channel()
	util.FailOnError(err, "Failed to open rabbit_mq channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		configs.Config.EmailQueue, // name
		true,                      // durable
		false,                     // delete when unused
		false,                     // exclusive
		false,                     // no-wait
		nil,                       // arguments
	)
	util.FailOnError(err, "Failed to declare a email queue")
	for n := 1; n <= 20; n++ {
		email := &model.Email{
			To:      fmt.Sprintf("receiver_test.%d@gmail.com", n),
			Subject: "Your_subject",
			Body:    fmt.Sprintf("this is the message body %d", rand.Int()),
		}
		body := `{ "to": "receiver_test@gmail.com", "subject": "Your_subject" , "body" : "this is the message body"}`
		b, _ := json.Marshal(email)
		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "application/json",
				Body:        b,
			})
		log.Printf(" [x] Sent %s", body)
		util.FailOnError(err, "Failed to publish a message")
	}

}
