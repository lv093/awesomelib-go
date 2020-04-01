package sarama_client

import (
	"fmt"
	"github.com/Shopify/sarama"
	"sync"
)

var (
	wg sync.WaitGroup
)

func Consumer() {
	consumer, err := sarama.NewConsumer([]string{"localhost:9091","localhost:9092","localhost:9093"}, nil)

	if err != nil {
		panic(err)
	}

	partitionList, err := consumer.Partitions("replicated_topic")
	println(partitionList)

	if err != nil {
		panic(err)
	}

	for partition := range partitionList {
		println(partition)
		pc, err := consumer.ConsumePartition("replicated_topic", int32(partition), 3)
		if err != nil {
			panic(err)
		}

		defer pc.AsyncClose()

		wg.Add(1)

		go func(sarama.PartitionConsumer) {
			defer wg.Done()
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d, Offset:%d, Key:%s, Value:%s\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
			}
		}(pc)
		wg.Wait()
		consumer.Close()
	}
}