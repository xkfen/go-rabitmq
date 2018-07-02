package main

import (
	"go-rabitmq/util"
	"fmt"
	"github.com/streadway/amqp"
	"log"
)
func main(){
	conn, cerr := util.GetRabitMqConn()
	if cerr != nil {
		fmt.Printf("%s", cerr.Error())
	}
	defer conn.Close()
	ch, chErr := util.CreateChannel(conn)
	if chErr != nil {
		fmt.Printf("%s", chErr.Error())
	}
	defer ch.Close()

	// 发送消息之前，必须创建消息队列，再将消息存入消息队列
	q, qErr := util.DeclareQueue(ch, "second", false, false, false, false, nil)
	if qErr != nil {
		fmt.Printf("%s", qErr.Error())
	}

	body := "hello world"


	err := ch.Publish("", q.Name, false, false, amqp.Publishing{
		DeliveryMode:amqp.Persistent,
		ContentType:"text/plain",
		Body:[]byte(body),
	})
	if err != nil {
		util.FailOnErr(err, "Failed to publish a message")
	}

	log.Printf("[x] sent %s", body)
}
