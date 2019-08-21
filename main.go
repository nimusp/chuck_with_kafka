package main

import (
	"./consumer"
	"./producer"
	"sync"
)

const (
	categoryURL    = "https://api.chucknorris.io/jokes/categories"
	jokeURL        = "https://api.chucknorris.io/jokes/random?category="
	topic          = "jokes"
	kafkaBrokerURL = "localhost:9092"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go producer.StartPublishingToTopic(topic, kafkaBrokerURL)
	go consumer.SubscribeToTopic(topic, kafkaBrokerURL, &wg)
	wg.Wait()
}
