package consumer

import (
	"context"
	"log"
	"sync"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

type consumer struct {
	topicName string
	brokerURL string
}

func NewConsumer(topicName string, brokerURL string) *consumer {
	return &consumer{
		topicName: topicName,
		brokerURL: brokerURL,
	}
}

func (c *consumer) SubscribeToTopic(wg *sync.WaitGroup) {
	defer wg.Done()
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{c.brokerURL},
		Topic:     c.topicName,
		Partition: 0,
	})

	for {
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()
		msg, err := reader.ReadMessage(ctx)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Received: theme — " + string(msg.Key) + ", joke — " + string(msg.Value))
	}
}
