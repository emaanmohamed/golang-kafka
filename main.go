package main

import (
	"fmt"
	appKafka "github.com/emaanmohamed/golang-kafka.git/kafka"
	"time"
)

func main() {
	fmt.Println("Okay")
	appKafka.StartKafka()
	fmt.Println("Kafka started successfully!")
	time.Sleep(10 * time.Second)

}
