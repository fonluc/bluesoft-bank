package messaging

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

// ConnectToKafka estabelece uma conexão com o Kafka.
func ConnectToKafka() (*kafka.Writer, error) {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "transactions",
	})

	// Testa se é possível escrever uma mensagem
	err := writer.WriteMessages(context.TODO(), kafka.Message{
		Value: []byte("Test message"),
	})
	if err != nil {
		return nil, err
	}
	log.Println("Connected to Kafka successfully")
	return writer, nil
}
