package main

import (
	"github.com/streadway/amqp"
	"go-rabitmq/util"
	"log"
	)

func main(){
	// open a connection
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	util.FailOnErr(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	// open a channel
	ch, err := conn.Channel()
	util.FailOnErr(err, "Failed to open a channel")
	defer ch.Close()
	// declare the queue from which we're going to consume
	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when usused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	util.FailOnErr(err, "Failed to declare a queue")
	//  make sure the queue exists before we try to consume messages
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	util.FailOnErr(err, "Failed to register a consumer")

	forever := make(chan bool)
	//  tell the server to deliver us the messages from the queue. Since it will push us messages asynchronously, we will read the messages from a channel in a goroutine
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}