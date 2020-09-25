package main

import (
	"rabbitmq-demo/rabbitmq"
	"time"

	"github.com/streadway/amqp"
)

func main() {
	dialUrl := "amqp://mq:123456@192.168.56.106:5672/"
	exchangeName := "demo_exchage"
	queueName := "demo_queue"
	publishMsg := amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		Timestamp:    time.Now(),
		ContentType:  "application/json",
		Body:         []byte("hello rabbitmq"),
	}
	rabbitmq.PublishMsg(dialUrl, exchangeName, queueName, publishMsg)
}
