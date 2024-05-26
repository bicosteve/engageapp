package controllers

import (
	"os"

	"github.com/engageapp/pkg/utils"
	amqp "github.com/rabbitmq/amqp091-go"
)

// Contains the code for RabbitMQ Consumer
func (b *Base) Consume(c *amqp.Channel, queueName string) {
	messages, err := c.Consume(
		queueName, // queue name
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no local
		false,     // no wait
		nil,       // arguments
	)

	if err != nil {
		utils.Log("ERROR", "queconsume", "error consuming because of %v ", err)
		os.Exit(1)
	}

	// utils.Log("INFO", "queconsume", "Waiting for messages")

	// Channel message infinite loop
	forever := make(chan bool)

	go func() {
		for message := range messages {
			utils.Log("INFO", "queconsume", "message %v", message.Body)

		}
	}()

	<-forever

}
