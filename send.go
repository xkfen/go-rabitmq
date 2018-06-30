package main

import (
	"github.com/streadway/amqp"
	"go-rabitmq/util"
	"log"
)

func main(){
	// 连到rabitMQ服务器
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	util.FailOnErr(err, "Failed to connect to RabitMQ")
	defer conn.Close()
	// 创建通道
	ch, err := conn.Channel()
	util.FailOnErr(err, "Failed to open a channel")
	defer ch.Close()
	// 发送消息之前，必须创建消息队列，再将消息存入消息队列
	q, err := ch.QueueDeclare(
		"hello",// queue name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no -wait
		nil, // arguments
		)
	util.FailOnErr(err, "Failed to declare a queue")
	body := "hello"
	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType:"text/plain",
			Body:[]byte(body),
		})
	log.Printf(" [x] Sent %s", body)
	util.FailOnErr(err, "Failed too publish a message")
}