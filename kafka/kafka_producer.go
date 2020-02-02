package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)
/*
	初始化NewConfig配置 sarama.NewConfig
	创建生产者sarama.NewSyncProducer
	创建消息sarama.ProducerMessage
	发送消息client.SendMessage
*/
func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	msg := &sarama.ProducerMessage{}
	msg.Topic = "test3"
	msg.Value = sarama.StringEncoder("this is a good test, my message is good")

	client, err := sarama.NewSyncProducer([]string{"localhost:9092","localhost:9090","localhost:9091"}, config)
	if err != nil {
		fmt.Println("producer close, err:", err)
		return
	}

	defer client.Close()

	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send message failed,", err)
		return
	}

	fmt.Printf("pid:%v offset:%v\n", pid, offset)
}


//package main
//
//import (
//	"bufio"
//	"fmt"
//	"os"
//	"github.com/Shopify/sarama"
//	"strings"
//)
//
//func main() {
//	config := sarama.NewConfig()
//	config.Producer.Return.Successes = true
//	config.Producer.RequiredAcks = sarama.WaitForAll
//	config.Producer.Partitioner = sarama.NewRandomPartitioner
//
//	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
//
//	if err != nil {
//		panic(err)
//	}
//	defer producer.Close()
//
//	msg := &sarama.ProducerMessage{
//		Topic:     "test2",
//		Partition: int32(1),
//		Key:       sarama.StringEncoder("key"),
//	}
//
//	var value string
//	for {
//		// 生产消息
//		inputReader := bufio.NewReader(os.Stdin)
//		value, err = inputReader.ReadString('\n')
//		if err != nil {
//			panic(err)
//		}
//		value = strings.Replace(value, "\n", "", -1)
//		msg.Value = sarama.ByteEncoder(value)
//		paritition, offset, err := producer.SendMessage(msg)
//
//		if err != nil {
//			fmt.Println("Send Message Fail")
//		}
//
//		fmt.Printf("Partion = %d, offset = %d\n", paritition, offset)
//	}
//}