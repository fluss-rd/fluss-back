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

type Range struct {
	Min float32
	Max float32
}

var (
	rangeGoodMeasurementTypePH = Range{
		Min: 7.4,
		Max: 7.5,
	}
	rangeNormalMeasurementTypePH = Range{
		Min: 6.5,
		Max: 8.5,
	}
	rangeBadMeasurementTypePH = Range{
		Min: 9,
		Max: 14,
	}

	rangeGoodMeasurementTypeDO = Range{
		Min: 9,
		Max: 20, // 20 is the max the sensor can read
	}
	rangeNormalMeasurementTypeDO = Range{
		Min: 5,
		Max: 8,
	}
	rangeBadMeasurementTypeDO = Range{
		Min: 0,
		Max: 4,
	}

	rangeColdMeasurementTypeTMP = Range{
		Min: 18,
		Max: 20,
	}
	rangeNormalMeasurementTypeTMP = Range{
		Min: 23,
		Max: 25,
	}
	rangeWarnMeasurementTypeTMP = Range{
		Min: 26,
		Max: 30,
	}

	rangeGoodMeasurementTypeTDY = Range{
		Min: 0,
		Max: 5,
	}
	rangeNormalMeasurementTypeTDY = Range{
		Min: 10,
		Max: 20,
	}
	rangeBadMeasurementTypeTDY = Range{
		Min: 20,
		Max: 1000,
	}

	rangeGoodMeasurementTypeTDS = Range{
		Min: 0,
		Max: 300,
	}
	rangeNormalMeasurementTypeTDS = Range{
		Min: 300,
		Max: 900,
	}
	rangeBadMeasurementTypeTDS = Range{
		Min: 900,
		Max: 1000,
	}
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
		phoneNumber: "+18091234321",
	},
	{
		quality:     "bad",
		phoneNumber: "+18291234321",
	},
	{
		quality:     "good",
		phoneNumber: "+18491234321",
	},
	{
		quality:     "bad",
		phoneNumber: "+18290987890",
	},
	{
		quality:     "good",
		phoneNumber: "+18090987890",
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

func randFloat(randRange Range) float32 {
	return randRange.Min + rand.Float32() * (randRange.Max - randRange.Min)
}

func getPh(quality string) float32 {
	switch quality {
		case "good":
			return randFloat(rangeGoodMeasurementTypePH)
		case "normal":
			return randFloat(rangeNormalMeasurementTypePH)
		case "bad":
			return randFloat(rangeBadMeasurementTypePH)
	}

	return 0
}

func getDissolvedOxygen(quality string) float32 {
	switch quality {
		case "good":
			return randFloat(rangeGoodMeasurementTypeDO)
		case "normal":
			return randFloat(rangeNormalMeasurementTypeDO)
		case "bad":
			return randFloat(rangeBadMeasurementTypeDO)
	}

	return 0
}

func getTDS(quality string) float32 {
	switch quality {
		case "good":
			return randFloat(rangeGoodMeasurementTypeTDS)
		case "normal":
			return randFloat(rangeNormalMeasurementTypeTDS)
		case "bad":
			return randFloat(rangeBadMeasurementTypeTDS)
	}

	return 0
}

func getTemperature() float32 {
	hour := time.Now().Hour()
	if hour > 20 && hour < 6 { // evening, "cold water"
		return randFloat(rangeColdMeasurementTypeTMP)
	}

	if hour >= 6 && hour < 12 { // day / morning , "average water"
		return randFloat(rangeNormalMeasurementTypeTMP)
	}

	if hour >= 12 && hour < 18 { // day / afternoon , "hot water"
		return randFloat(rangeWarnMeasurementTypeTMP)
	}

	return 28
}

func getTurbidity(quality string) float32 {
	switch quality {
		case "good":
			return randFloat(rangeGoodMeasurementTypeTDY)
		case "normal":
			return randFloat(rangeNormalMeasurementTypeTDY)
		case "bad":
			return randFloat(rangeBadMeasurementTypeTDY)
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
