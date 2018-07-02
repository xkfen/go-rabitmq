package util

import (
	"log"
	"fmt"
	"github.com/streadway/amqp"
)

// 错误处理
func FailOnErr(err error, msg string){
	if err != nil {
		log.Fatalf("%s: %s", msg, err.Error())
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

// 得到rabitMQ连接
func GetRabitMqConn()(*amqp.Connection, error){
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		FailOnErr(err, "Failed to connect to RabitMQ")
		return nil, err
	}

	return conn, nil
}

// 创建通道channel
func CreateChannel(conn *amqp.Connection)(*amqp.Channel, error){
	ch, err := conn.Channel()
	if err != nil {
		FailOnErr(err, "Failed to open a channel")
		return nil, err
	}
	return ch, nil
}

// 声明队列
func DeclareQueue(channel *amqp.Channel, name string, durable, autoDelete, exclusive, noWait bool, args amqp.Table)(amqp.Queue, error){
	q, err := channel.QueueDeclare(
		name,// queue name
		durable, // durable
		autoDelete, // delete when unused
		exclusive, // exclusive
		noWait, // no -wait
		args, // arguments
	)
	if err != nil {
		FailOnErr(err, "Failed to declare a queue")
		return amqp.Queue{}, err
	}
	return q, nil
}