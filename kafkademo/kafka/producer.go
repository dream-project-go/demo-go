package kafka

//import (
//	"log"
//
//	"github.com/Shopify/sarama"
//)
//
//func publishMsg(topic){
//	producer, err := sarama.NewSyncProducer([]string{"192.168.56.110:9092"}, nil)
//	if err != nil {
//		log.Fatalln(err)
//	}
//	defer func() {
//		if err := producer.Close(); err != nil {
//			log.Fatalln(err)
//		}
//	}()
//	msg := &sarama.ProducerMessage{Topic: "test0", Value: sarama.StringEncoder("testing 123")}
//	partition, offset, err := producer.SendMessage(msg)
//	if err != nil {
//		log.Printf("FAILED to send message: %s\n", err)
//	} else {
//		log.Printf("> message sent to partition %d at offset %d\n", partition, offset)
//	}
//}