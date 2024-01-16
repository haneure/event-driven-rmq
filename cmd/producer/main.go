package main

import (
	"log"
	"time"

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

	time.Sleep(60 * time.Second)

	log.Println(client)
}
