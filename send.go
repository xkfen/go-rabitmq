package main

import (
	"github.com/streadway/amqp"
	"go-rabitmq/util"
	"log"
	"fmt"
)

func main(){
	// 连到rabitMQ服务器
	conn, cerr := util.GetRabitMqConn()
	if cerr != nil {
		fmt.Printf("%s", cerr.Error())
	}
	/**
	which will be executed when the function return. That is to say, whenever init finished, your connection 	will be closed, which causes unopen connection.
	You need to defer the connection closing in main, or somewhere you want it to be closed.
	 */
	// 这个不能放在util的方法里面，因为随着GetRabitMqConn的return，conn连接就会被关闭
	defer conn.Close()
	// 创建通道
	ch, cherr := util.CreateChannel(conn)
	if cerr != nil {
		fmt.Printf("%s", cherr.Error())
	}
	// 这个不能放在util的方法里面，因为随着CreateChannel的return，conn/channel连接就会被关闭
	defer ch.Close()
	// 发送消息之前，必须创建消息队列，再将消息存入消息队列
	q, qErr := util.DeclareQueue(ch, "first", false, false, false, false, nil)
	if qErr != nil {
		fmt.Printf("%s", qErr.Error())
	}
	body := "hello"
	err := ch.Publish(
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