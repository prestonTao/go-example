package queue

import (
	"fmt"
	"github.com/streadway/amqp"
	// "github.com/virushuo/Go-Apns"
)

var (
	queueName   = "test-queue"
	consumerTag = "simple-consumer"
)

type Worker struct {
	channel *amqp.Channel
}

func (this *Worker) Connect() {
	conn, err := amqp.Dial(amqpUri)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	this.channel, err = conn.Channel()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = this.channel.ExchangeDeclare(
		exchangeName, // name of the exchange
		exchangeType, // type
		true,         // durable
		false,        // delete when complete
		false,        // internal
		false,        // noWait
		nil,          // arguments
	)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	queue, err := this.channel.QueueDeclare(
		queueName, // name of the queue
		true,      // durable
		false,     // delete when usused
		false,     // exclusive
		false,     // noWait
		nil,       // arguments
	)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = this.channel.QueueBind(
		queue.Name,   // name of the queue
		routingKey,   // bindingKey
		exchangeName, // sourceExchange
		false,        // noWait
		nil,          // arguments
	)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	deliveries, err := this.channel.Consume(
		queue.Name,  // name
		consumerTag, // consumerTag,
		false,       // noAck
		false,       // exclusive
		false,       // noLocal
		false,       // noWait
		nil,         // arguments
	)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for d := range deliveries {
		fmt.Println("recv : ", string(d.Body))
		d.Ack(false)
	}

}

// func pushApns() {
// 	apn, err := apns.New("apns_cert.pem", "apns_key_nosecure.pem", "gateway.sandbox.push.apple.com:2195", 1*time.Second)
// 	if err != nil {
// 		fmt.Printf("connect error: %s\n", err.Error())
// 		os.Exit(1)
// 	}

// }
