package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

func PublishMsg(dialUrl, exchangeName, queueName string, publishMsg amqp.Publishing) {
	conn, err := amqp.Dial(dialUrl) //获取mq连接
	if err != nil {
		log.Fatalf("connection.open: %s", err)
	}

	defer conn.Close()

	c, err := conn.Channel() //声明通道
	if err != nil {
		log.Fatalf("channel.open: %s", err)
	}

	err = c.ExchangeDeclare(exchangeName, "topic", true, false, false, false, nil) //声明交换器
	if err != nil {
		log.Fatalf("exchange.declare: %v", err)
	}

	// Prepare this message to be persistent.  Your publishing requirements may
	// be different.
	msg := amqp.Publishing{
		DeliveryMode: publishMsg.DeliveryMode,
		Timestamp:    publishMsg.Timestamp,
		ContentType:  publishMsg.ContentType,
		Body:         publishMsg.Body,
	}

	// This is not a mandatory delivery, so it will be dropped if there are no
	// queues bound to the logs exchange.
	err = c.Publish(exchangeName, queueName, false, false, msg)
	if err != nil {
		// Since publish is asynchronous this can happen if the network connection
		// is reset or if the server has run out of resources.
		log.Fatalf("basic.publish: %v", err)
	}
}
