package main

import (
	"log"

	"github.com/haneure/eventdrivenrabbit/internal"
)

func main() {
	conn, err := internal.ConnectRabbitMQ("admin", "12345678", "localhost:5672", "customers")

	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client, err := internal.NewRabbitMQClient(conn)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	messageBus, err := client.Consume("customers_created", "email-service", false)
	if err != nil {
		panic(err)
	}

	var blocking chan struct{}

	go func() {
		for message := range messageBus {
			log.Printf("New Message: %v", message)
			// if err := message.Ack(false); err != nil {
			// 	log.Println("Acknowledge message failed")
			// 	continue
			// }

			if !message.Redelivered {
				message.Nack(false, true)
				continue
			}
			if err := message.Ack(false); err != nil {
				log.Println("Failed to ack message")
				continue
			}
			
			log.Printf("Acknowledge message %s\n", message.MessageId)
		}
	}()

	log.Println("Consuming, to close the program press CTRL+C")

	<-blocking
}
