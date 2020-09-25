package main

import (
	"log"
	"rabbitmqdemo/rabbitmq"
)

func main() {
	dialUrl := "amqp://mq:123456@192.168.56.106:5672/"
	exchangeName := "demo_exchage"
	bindings := []rabbitmq.Queuebind{
		{Queue: "demo_queue", Key: "demo_queue"},
	}
	c, err := rabbitmq.NewConsumer(dialUrl, exchangeName, bindings, 1)
	demoMsgs, err := c.Consume("demo_queue", "", false, false, false, false, nil)
	if err != nil {
		log.Fatalf("basic.consume: %v", err)
	}
	for demoMsg := range demoMsgs {
		log.Println(string(demoMsg.Body))
		demoMsg.Ack(false)
	}
}
