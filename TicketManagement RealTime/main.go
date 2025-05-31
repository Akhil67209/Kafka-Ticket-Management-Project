package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/segmentio/kafka-go"
	//"time"
)

type TicketMessage struct {
	TicketID  string `json:"ticket_id"`
	UserPhone string `json:"user_phone"`
	Status    string `json:"status"`
}

func main() {

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "ticket-topic",
	})

	defer writer.Close()

	message := TicketMessage{
		TicketID:  "T12345",
		UserPhone: "+919790702167",
		Status:    "confirmed",
	}

	msgBytes, err := json.Marshal(message)
	if err != nil {
		panic(err)
	}

	err = writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(message.TicketID),
			Value: msgBytes,
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("Ticket confirmation sent successfully.")
}
