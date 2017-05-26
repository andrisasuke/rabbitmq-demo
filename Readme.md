RabbitMQ golang consumer
Study case : put string to rabbit Queue and sends it as email (SMTP)
1. Download and install RabbitMQ https://www.rabbitmq.com/#getstarted
2. Install Golang (mine is v1.6)
3. Checkout into your $GOPATH
4. Put to queue : $ go run send.go
5. Consume from queue : $ go run receive.go
