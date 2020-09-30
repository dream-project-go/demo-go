package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

func publishDemo(topic string, msg string) {

}

func consumerDemo() {
	fmt.Printf("consumer_test")

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Version = sarama.V0_11_0_2

	// consumer
	consumer, err := sarama.NewConsumer([]string{"192.168.56.110:9092"}, config)
	if err != nil {
		fmt.Printf("consumer_test create consumer error %s\n", err.Error())
		return
	}

	defer consumer.Close()

	partition_consumer, err := consumer.ConsumePartition("test0", 3, sarama.OffsetOldest)
	if err != nil {
		fmt.Printf("try create partition_consumer error %s\n", err.Error())
		return
	}
	defer partition_consumer.Close()

	for {
		select {
		case msg := <-partition_consumer.Messages():
			fmt.Printf("msg offset: %d, partition: %d, timestamp: %s, value: %s\n",
				msg.Offset, msg.Partition, msg.Timestamp.String(), string(msg.Value))
		case err := <-partition_consumer.Errors():
			fmt.Printf("err :%s\n", err.Error())
		}
	}

}

func main() {
	// producer, err := sarama.NewSyncProducer([]string{"192.168.56.110:9092"}, nil)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// defer func() {
	// 	if err := producer.Close(); err != nil {
	// 		log.Fatalln(err)
	// 	}
	// }()
	// msg := &sarama.ProducerMessage{Topic: "test0", Value: sarama.StringEncoder("testing 12388888")}
	// partition, offset, err := producer.SendMessage(msg)
	// if err != nil {
	// 	log.Printf("FAILED to send message: %s\n", err)
	// } else {
	// 	log.Printf("> message sent to partition %d at offset %d\n", partition, offset)
	// }

	//消费demo
	consumerDemo()

}
