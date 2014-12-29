package queue

import (
	"fmt"
	"github.com/streadway/amqp"
)

var (
	amqpUri      = "amqp://guest:guest@localhost:5672/"
	exchangeName = "test-exchange"
	exchangeType = "direct"
	routingKey   = "test-key"
)

type MsgQueue struct {
	channel *amqp.Channel
}

func (this *MsgQueue) StartUP() {
	connection, err := amqp.Dial(amqpUri)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	this.channel, err = connection.Channel()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if err = this.channel.ExchangeDeclare(
		exchangeName, // name
		exchangeType, // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // noWait
		nil,          // arguments
	); err != nil {
		fmt.Println(err.Error())
		// fmt.Errorf("Exchange Declare: %s", err)
		return
	}

}

func (this *MsgQueue) Send(msg string) {
	if err := this.channel.Publish(
		exchangeName, // publish to an exchange
		routingKey,   // routing to 0 or more queues
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			Headers:         amqp.Table{},
			ContentType:     "text/plain",
			ContentEncoding: "",
			Body:            []byte(msg),
			DeliveryMode:    amqp.Transient, // 1=non-persistent, 2=persistent
			Priority:        0,              // 0-9
			// a bunch of application/implementation-specific fields
		},
	); err != nil {
		// fmt.Errorf("Exchange Publish: %s", err)
		fmt.Println(err.Error())
		return
	}
}

// func NewQueue() *MsgQueue {
// 	msgQueue := MsgQueue{
// 		exchangeName: "test-exchange",
// 		exchangeType: "direct",
// 		routingKey:   "test-key",
// 		msg:          "foobar",
// 		reliable:     false,
// 	}
// 	return msgQueue
// }
