RabbitMQ golang consumer demo
- Study case : put string to rabbit Queue and sends it as email (SMTP)
- Support retry connection
1. Download and install RabbitMQ `https://www.rabbitmq.com/#getstarted`
2. Install Golang (mine is v1.14)
3. Checkout repository
4. Install dependency amq : `$ go get github.com/streadway/amqp`
   and `$ go get get github.com/spf13/viper`
5. Put to queue : `$ go run main.go send`
6. Consume from queue : `$ go run main.go receive.go`

How to test Reconnect
1. Run receiver `$go run main.go receive`
2. Stop rabbitmq server
3. Start rabbitmq server

Output will be
```
$ go run main.go receive  
2021/11/19 06:27:02 Reading configuration
Run agument is receive
Starting amq consumer

connection is closed :320 reason: CONNECTION_FORCED - broker forced connection closure with reason 'shutdown'
channel is closed :320 reason: CONNECTION_FORCED - broker forced connection closure with reason 'shutdown'
unable recreate channel reason: Exception (504) Reason: "channel/connection is not open"
unable reconnect to rabbitmq :dial tcp [::1]:5672: connect: connection refused
unable recreate channel reason: Exception (504) Reason: "channel/connection is not open"
unable recreate channel reason: Exception (504) Reason: "channel/connection is not open"
unable reconnect to rabbitmq :dial tcp [::1]:5672: connect: connection refused
unable recreate channel reason: Exception (504) Reason: "channel/connection is not open"
unable recreate channel reason: Exception (504) Reason: "channel/connection is not open"
unable reconnect to rabbitmq :Exception (501) Reason: "EOF"
unable recreate channel reason: Exception (504) Reason: "channel/connection is not open"
unable reconnect to rabbitmq :Exception (501) Reason: "EOF"
unable recreate channel reason: Exception (504) Reason: "channel/connection is not open"
unable recreate channel reason: Exception (504) Reason: "channel/connection is not open"
unable reconnect to rabbitmq :Exception (501) Reason: "EOF"
unable recreate channel reason: Exception (504) Reason: "channel/connection is not open"
unable recreate channel reason: Exception (504) Reason: "channel/connection is not open"
create new connection for consumer channel
reconnect success
channel recreate success
```