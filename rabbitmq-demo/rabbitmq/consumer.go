package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

type Queuebind struct {
	Queue string
	Key   string
}

func NewConsumer(dialUrl, exchangeName string, bindings []Queuebind, QosPrefetchCount int) (*amqp.Channel, error) {
	conn, err := amqp.Dial(dialUrl)
	if err != nil {
		log.Fatalf("connection.open: %s", err)
	}
	// defer conn.Close()
	c, err := conn.Channel()
	if err != nil {
		log.Fatalf("channel.open: %s", err)
	}
	err = c.ExchangeDeclare(exchangeName, "topic", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("exchange.declare: %s", err)
	}
	for _, b := range bindings {
		_, err = c.QueueDeclare(b.Queue, true, false, false, false, nil)
		if err != nil {
			log.Fatalf("queue.declare: %v", err)
		}
		log.Println("queue", b.Queue, "Key", b.Key, "exchangeName", exchangeName)

		err = c.QueueBind(b.Queue, b.Key, exchangeName, false, nil)
		if err != nil {
			log.Fatalf("queue.bind: %v", err)
		}
	}
	err = c.Qos(QosPrefetchCount, 0, false)
	if err != nil {
		log.Fatalf("basic.qos: %v", err)
	}
	return c, err
}
