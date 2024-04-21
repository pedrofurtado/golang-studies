package main

import (
	"fmt"
	"time"
  "crypto/rand"
  "encoding/hex"
)

func Sleep(i int) {
	time.Sleep(time.Duration(i) * time.Second)
}

func GenerateRandomData() string {
	n := 4
  bytes := make([]byte, n)
  if _, err := rand.Read(bytes); err != nil {
    return "err"
  }
  return hex.EncodeToString(bytes)
}

func InitProducers(ch chan string) {
	numberOfProducers := 3

	for i := range numberOfProducers {
		fmt.Printf("[Main] Init Producer %v\n", i)
		go InitProducer(ch, i)
	}
}

func InitProducer(ch chan string, i int) {
	for {
		fmt.Printf("[Producer %v] Polling data\n", i)
		Sleep(i + 1)
		randomData := GenerateRandomData()
		fmt.Printf("[Producer %v] [Data %v] Data polled | Inserting in the channel\n", i, randomData)
		ch <- randomData
	}
}

func InitConsumers(ch chan string) {
	numberOfConsumers := 5

	for i := range numberOfConsumers {
		fmt.Printf("[Main] Init Consumer %v\n", i)
		go InitConsumer(ch, i)
	}
}

func InitConsumer(ch chan string, i int) {
	for data := range ch {
		fmt.Printf("[Consumer %v] [Data %v] Processing\n", i, data)
		Sleep(i + 1)
		fmt.Printf("[Consumer %v] [Data %v] Processed\n", i, data)
	}
}

func main() {
	fmt.Println("[Main] Start")

	fmt.Println("[Main] Init Channel")
	ch := make(chan string)

	fmt.Println("[Main] Init Producers")
	InitProducers(ch)

	fmt.Println("[Main] Init Consumers")
	InitConsumers(ch)

	fmt.Println("[Main] Init loop daemon")
	select{}
}
