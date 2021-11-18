package main

import (
	"os"
	"rabbitmq-demo/client"
	"rabbitmq-demo/configs"
)

func main() {
	if len(os.Args) <= 1 {
		println("need run argument receive or send")
		os.Exit(1)
	}

	if arg := os.Args[1]; arg != "receive" && arg != "send" {
		println("need run argument receive or send")
		os.Exit(1)
	}

	configs.ReadConfig()
	println("Run agument is", os.Args[1])

	if os.Args[1] == "receive" {
		client.NewAMQConsumer()
	} else {
		client.Send()
	}
}
