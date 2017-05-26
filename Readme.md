RabbitMQ golang consumer
- Study case : put string to rabbit Queue and sends it as email (SMTP)
1. Download and install RabbitMQ `https://www.rabbitmq.com/#getstarted`
2. Install Golang (mine is v1.6)
3. Checkout repository into your $GOPATH
4. Install dependency amq : `$ go get github.com/streadway/amqp`
   and `$ go get get github.com/spf13/viper`
5. Put to queue : `$ go run send.go`
6. Consume from queue : `$ go run receive.go`
7. Or Run as binary `$ go install` and `$ ./rabbitmq-demo`
