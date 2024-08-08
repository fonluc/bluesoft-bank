package main

import (
	"fmt"
	"log"

	"github.com/fonluc/bluesoft-bank/internal/messaging"
)

func main() {
	writer, err := messaging.ConnectToKafka()
	if err != nil {
		log.Fatal("Error connecting to Kafka:", err)
	}
	fmt.Println("Kafka connection successful")
	writer.Close()
}
