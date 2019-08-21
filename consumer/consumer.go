package consumer

import (
	"context"
	"log"
	"sync"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

func SubscribeToTopic(kafkaTopic string, brokerURL string, wg *sync.WaitGroup) {
	defer wg.Done()
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{brokerURL},
		Topic:     kafkaTopic,
		Partition: 0,
	})

	for {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		msg, err := reader.ReadMessage(ctx)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Received: theme — " + string(msg.Key) + ", joke — " + string(msg.Value))
	}
}
