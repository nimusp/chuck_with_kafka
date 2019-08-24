package main

import (
	"./consumer"
	"./producer"
	"sync"
	"time"
)

const (
	topic          = "jokes"
	kafkaBrokerURL = "kafka:9092"
)

func main() {
	// ожидание запуска сервиса Kafka
	time.Sleep(15 * time.Second)

	producer := producer.NewPublisher(topic, kafkaBrokerURL)
	consumer := consumer.NewConsumer(topic, kafkaBrokerURL)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go producer.StartPublishingToTopic()
	go consumer.SubscribeToTopic(&wg)
	wg.Wait()
}
