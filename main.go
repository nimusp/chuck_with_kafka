package main

import (
	"./consumer"
	"./producer"
	"sync"
)

const (
	topic          = "jokes"
	kafkaBrokerURL = "127.0.0.1:9092"
)

func main() {
	producer := producer.NewPublisher(topic, kafkaBrokerURL)
	consumer := consumer.NewConsumer(topic, kafkaBrokerURL)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go producer.StartPublishingToTopic()
	go consumer.SubscribeToTopic(&wg)
	wg.Wait()
}
