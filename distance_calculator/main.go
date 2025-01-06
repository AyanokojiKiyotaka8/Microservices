package main

import (
	"log"
)

func main() {
	var kafkaTopic = "obudata"
	calcService := NewCalculatorService()
	kafkaConsumer, err := NewKafkaConsumer(kafkaTopic, calcService)
	if err != nil {
		log.Fatal(err)
	}
	kafkaConsumer.Start()
}
