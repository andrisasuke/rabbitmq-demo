package main

import (
	"log"
	"fmt"
	"github.com/spf13/viper"
)

type AppConfig struct {
	RabbitMQUri   	    string
	SmtpUserAccount     string
	SmtpPasswordAccount string
	SmtpHost	    string
	SmtpAddress	    string
	EmailQueue  	    string
}

var Config *AppConfig

func ReadConfig()  {
	log.Println("Reading configuration")
	viper.SetDefault("env", "development")
	viper.BindEnv("env")
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/rabbitmq_consumer/")
	viper.AddConfigPath("$HOME/.rabbitmq_consumer")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("/Data/")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	config := viper.GetStringMapString(viper.GetString("env"))

	Config = &AppConfig{
		RabbitMQUri:    config["rabbitmq_uri"],
		SmtpUserAccount:   config["smtp_user_account"],
		SmtpPasswordAccount: config["smtp_user_password"],
		SmtpHost: config["smtp_host"],
		SmtpAddress: config["smtp_address"],
		EmailQueue:  config["email_queue"],
	}
}