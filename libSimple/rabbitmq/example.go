package main

import (
	"fmt"
	"github.com/streadway/amqp"
)

func main() {
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println("连接服务器失败 ", err.Error())
	}
	defer connection.Close()

	//获取管道
	channel, err := connection.Channel()
	if err != nil {
		fmt.Errorf("Channel: %s", err)
	}

	if err := channel.ExchangeDeclare(
		"test",   // name
		"direct", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // noWait
		nil,      // arguments
	); err != nil {
		fmt.Errorf("Exchange Declare: %s", err)
	}

	if err = channel.Publish(
		"direct", // publish to an exchange
		"test",   // routing to 0 or more queues
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			Headers:         amqp.Table{},
			ContentType:     "text/plain",
			ContentEncoding: "",
			Body:            []byte("nihao"),
			DeliveryMode:    amqp.Transient, // 1=non-persistent, 2=persistent
			Priority:        0,              // 0-9
			// a bunch of application/implementation-specific fields
		},
	); err != nil {
		fmt.Errorf("Exchange Publish: %s", err)
	}

}
