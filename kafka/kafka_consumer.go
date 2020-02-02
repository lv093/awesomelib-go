package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"sync"
)

var (
	wg sync.WaitGroup
)

func main() {
	consumer, err := sarama.NewConsumer([]string{"localhost:9090","localhost:9091","localhost:9092"}, nil)

	if err != nil {
		panic(err)
	}

	partitionList, err := consumer.Partitions("test3")

	if err != nil {
		panic(err)
	}

	for partition := range partitionList {
		pc, err := consumer.ConsumePartition("test3", int32(partition), 3)
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