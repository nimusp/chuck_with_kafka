package producer

import (
	"../model"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

const (
	categoryURL = "https://api.chucknorris.io/jokes/categories"
	jokeURL     = "https://api.chucknorris.io/jokes/random?category="
)

type publisher struct {
	topicName string
	brokerURL string
}

func NewPublisher(topicName string, brokerURL string) *publisher {
	return &publisher{
		topicName: topicName,
		brokerURL: brokerURL,
	}
}

func (p *publisher) StartPublishingToTopic() {
	connection, err := kafka.DialLeader(context.Background(), "tcp", p.brokerURL, p.topicName, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()

	clientWithTimeout := http.Client{Timeout: 5 * time.Second}
	mutex := sync.Mutex{}

	processJoke := func() {
		category := loadCategory(&clientWithTimeout, categoryURL)
		joke := loadJoke(&clientWithTimeout, category)

		mutex.Lock()
		defer mutex.Unlock()
		_, err = connection.WriteMessages(
			kafka.Message{
				Key:   []byte(category),
				Value: []byte(joke.Value),
			},
		)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Wrote to Kafka " + category)
	}

	for {
		go processJoke()
		time.Sleep(500 * time.Millisecond)
	}
}

func loadCategory(client *http.Client, categoryURL string) string {
	resp, err := client.Get(categoryURL)
	logFatal(err)

	categoryList := make([]interface{}, 0)
	body, err := ioutil.ReadAll(resp.Body)
	logFatal(err)
	defer resp.Body.Close()
	json.Unmarshal(body, &categoryList)

	categoryNumber := rand.Intn(len(categoryList))
	item := categoryList[categoryNumber]
	category, ok := item.(string)
	if !ok {
		log.Fatal("Error on cast category")
	}

	return category
}

func loadJoke(client *http.Client, category string) *model.ChuckResponseModel {
	response, err := client.Get(jokeURL + category)
	logFatal(err)

	respModel := &model.ChuckResponseModel{}
	rawData, err := ioutil.ReadAll(response.Body)
	logFatal(err)
	defer response.Body.Close()

	err = json.Unmarshal(rawData, &respModel)
	logFatal(err)

	return respModel
}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
