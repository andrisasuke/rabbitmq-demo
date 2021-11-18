package configs

import (
	"log"
	"strconv"

	"github.com/spf13/viper"
)

type AppConfig struct {
	RabbitMQUri       string
	SmtpAddress       string
	EmailQueue        string
	EmailSenderWorker int
}

var Config *AppConfig

func ReadConfig() {
	log.Println("Reading configuration")
	viper.SetDefault("env", "development")
	viper.BindEnv("env")
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/rabbitmq_consumer/")
	viper.AddConfigPath("$HOME/.rabbitmq_consumer")
	viper.AddConfigPath("./config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/Data/")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	config := viper.GetStringMapString(viper.GetString("env"))

	mailSenderWorker, e := strconv.Atoi(config["email_sender_worker"])
	if e != nil {
		panic(e)
	}

	Config = &AppConfig{
		RabbitMQUri:       config["rabbitmq_uri"],
		SmtpAddress:       config["smtp_host"],
		EmailQueue:        config["email_queue"],
		EmailSenderWorker: mailSenderWorker,
	}
}
