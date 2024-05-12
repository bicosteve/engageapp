package utils

import (
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Contains RabbitMQ helpers

// Creates Connection to RabbitMQ
func ConnectQueue() *amqp.Connection {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		Log("ERROR", "queue", "problem connecting to rabbitmq because of %v \n", err)
		os.Exit(1)
	}

	Log("INFO", "queue", "Connected to rabbitmq successfully")
	return conn
}

// Creates a channel on rabbitmq
func CreateChannel(q *amqp.Connection) *amqp.Channel {
	// Log("INFO", "queue", "Creating a channel")
	channel, err := q.Channel()
	if err != nil {
		Log("ERROR", "quechan", "error creating channel because of %v ", err)
		os.Exit(1)
	}

	// Declare Queue to publish to
	_, err = channel.QueueDeclare(
		"Test", // queue name
		true,   // durable
		false,  // auto delete
		false,  // exclusive
		false,  // no wait
		nil,    // arguments
	)

	if err != nil {
		Log("ERROR", "queuechan", "error declaring queue because of %v ", err)
		os.Exit(1)
	}

	// defer channel.Close()

	Log("INFO", "queue", "Channel created")

	return channel
}

func PublishToQueue(chn *amqp.Channel, quename string, payload interface{}) error {
	bytes, ok := payload.([]byte)
	_ = ok

	msg := amqp.Publishing{
		ContentType: "application/json",
		Body:        []byte(bytes),
	}

	err := chn.Publish(
		"", // Exchange
		quename,
		false,
		false,
		msg,
	)
	if err != nil {
		return err
	}

	return nil
}
