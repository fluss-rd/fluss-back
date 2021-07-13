package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/robfig/cron/v3"
)

type Module struct {
	phoneNumber string
	quality     string
}

type MessageRequestBody struct {
	PhoneNumber  string        `json:"phoneNumber"`
	Measurements []Measurement `json:"measurements"`
}

type Measurement struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

var modules = []Module{
	{
		quality:     "normal",
		phoneNumber: "+18091231122",
	},
	{
		quality:     "bad",
		phoneNumber: "+18091231123",
	},
	{
		quality:     "good",
		phoneNumber: "+18091231124",
	},
}

func main() {
	c := cron.New()

	_, err := c.AddFunc("* * * * *", sendFakeData)

	if err != nil {
		log.Fatal(err)
	}

	c.Start()
	defer c.Stop()

	forever := make(chan bool)

	<-forever
}

func getPh(quality string) float32 {
	switch quality {
	case "good":
		return rand.Float32() * (8.7 - 7.4)
	case "normal":
		return rand.Float32() * (9 - 6)
	case "bad":
		return rand.Float32() * (4 - 0)
	}

	return 0
}

func getDissolvedOxygen(quality string) float32 {
	switch quality {
	case "good":
		return rand.Float32() * (20 - 9) // 20 is the max the sensor can read
	case "normal":
		return rand.Float32() * (8 - 5)
	case "bad":
		return rand.Float32() * (5 - 0)
	}

	return 0
}

func getTDS(quality string) float32 {
	switch quality {
	case "good":
		return rand.Float32() * (600 - 0)
	case "normal":
		return rand.Float32() * (900 - 600)
	case "bad":
		return rand.Float32() * (1000 - 900)
	}

	return 0
}

func getTemperature() float32 {
	hour := time.Now().Hour()
	if hour > 20 && hour < 6 { // evening, "cold water"
		return rand.Float32() * (20 - 18)
	}

	if hour >= 6 && hour < 12 { // day / morning , "average water"
		return rand.Float32() * (25 - 23)
	}

	if hour >= 12 && hour < 18 { // day / afternoon , "hot water"
		return rand.Float32() * (30 - 26)
	}

	return 28
}

func getTurbidity(quality string) float32 {
	switch quality {
	case "good":
		return rand.Float32() * (33.33 - 0)
	case "normal":
		return rand.Float32() * (67 - 34)
	case "bad":
		return rand.Float32() * (100 - 67)
	}

	return 0
}

func sendFakeData() {
	log.Println("Starting send messages task...")

	client := http.Client{}

	for _, module := range modules {
		params := []Measurement{}

		turbidity := getTurbidity(module.quality)
		params = append(params, Measurement{
			Name:  "ty",
			Value: float64(turbidity),
		})

		temperature := getTemperature()
		params = append(params, Measurement{
			Name:  "tmp",
			Value: float64(temperature),
		})

		tds := getTDS(module.quality)
		params = append(params, Measurement{
			Name:  "tds",
			Value: float64(tds),
		})

		dissolvedOxygen := getDissolvedOxygen(module.quality)
		params = append(params, Measurement{
			Name:  "do",
			Value: float64(dissolvedOxygen),
		})

		ph := getPh(module.quality)
		params = append(params, Measurement{
			Name:  "ph",
			Value: float64(ph),
		})

		message := MessageRequestBody{
			PhoneNumber:  module.phoneNumber,
			Measurements: params,
		}

		log.Println("Unmarshalling body...")
		body, err := json.Marshal(message)
		if err != nil {
			log.Println("failed to send message: ", err.Error())
			continue
		}
		// w.Close()

		fmt.Println(string(body))
		log.Println("Sending request")

		request, err := http.NewRequest(http.MethodPost, os.Getenv("SEND_MESSAGES_URL"), bytes.NewBuffer(body))
		if err != nil {
			log.Println("failed to create request: ", err.Error())
			continue
		}

		response, err := client.Do(request)
		if err != nil {
			log.Println("failed to send message: ", err.Error())
			continue
		}

		if response.StatusCode != http.StatusAccepted {
			log.Println("invalid status code in response: ", response.StatusCode)
			continue
		}

		log.Println("Message sent succesfully for module: ", module.phoneNumber)
	}

}
