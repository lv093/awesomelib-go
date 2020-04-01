package sarama_client

import (
	"github.com/Shopify/sarama"
	"fmt"
	"strconv"
)

/**
异步生产
	初始化NewConfig配置 sarama.NewConfig
	创建生产者sarama.NewAsyncProducer
	创建消息sarama.ProducerMessage
	发送消息producer.Input() <- msg
	接收消息回调：select：{suc := <-producer.Successes()}{fail := <-producer.Errors()}
 */
func AsyncProducer() {
	fmt.Printf("producer_test\n")
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Version = sarama.V0_10_1_1

	fmt.Println("begin producer construct")
	producer, err := sarama.NewAsyncProducer([]string{"localhost:9091","localhost:9092","localhost:9093"}, config)
	fmt.Println("producer,%t", producer)
	if err != nil {
		fmt.Printf("producer_test create producer error :%s\n", err.Error())
		return
	}

	defer producer.AsyncClose()

	// send message
	msg := &sarama.ProducerMessage{
		Topic: "replicated_topic",
		Key:   sarama.StringEncoder("go_test"),
	}

	value := "this is message"
	i := 1
	for {
		//fmt.Scanln(&value)
		value = "hello:" + strconv.Itoa(i)
		i++
		msg.Value = sarama.ByteEncoder(value)
		fmt.Printf("input [%s]\n", value)

		// send to chain
		producer.Input() <- msg

		select {
		case suc := <-producer.Successes():
			fmt.Printf("offset: %d,  timestamp: %s", suc.Offset, suc.Timestamp.String())
		case fail := <-producer.Errors():
			fmt.Printf("err: %s\n", fail.Err.Error())
		}
	}
}


/*
同步生产
	初始化NewConfig配置 sarama.NewConfig
	创建生产者sarama.NewSyncProducer
	创建消息sarama.ProducerMessage
	发送消息client.SendMessage，同步获取返回信息
*/
func SyncProducer() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	msg := &sarama.ProducerMessage{}
	msg.Topic = "replicated_topic"
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

