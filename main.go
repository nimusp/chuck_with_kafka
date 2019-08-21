package main

import (
	"./model"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
)

const (
	categoryURL = "https://api.chucknorris.io/jokes/categories"
	jokeURL     = "https://api.chucknorris.io/jokes/random?category="
)

func main() {
	resp, err := http.Get(categoryURL)
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

	log.Println("Load joke from category " + category + "...")
	jokeModel := loadJoke(category)
	log.Println(jokeModel.Value)
}

func loadJoke(category string) *model.ChuckResponseModel {
	response, err := http.Get(jokeURL + category)
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
