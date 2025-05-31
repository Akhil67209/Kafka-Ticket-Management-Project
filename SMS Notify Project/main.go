package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/segmentio/kafka-go"
)

type TicketMessage struct {
	TicketID  string `json:"ticket_id"`
	UserPhone string `json:"user_phone"`
	Status    string `json:"status"`
}

func sendSMS(phone, text string) {
	fmt.Printf("Sending SMS to %s: %s\n", phone, text)
}

func main() {

	ctx := context.Background()
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "ticket-topic",
		GroupID:  "sms_notification_group_debug_v2", // Use a new group ID to consume from the beginning
		MinBytes: 1,
		MaxBytes: 10e6,
	})

	defer reader.Close()

	fmt.Println("SMS Notification Service started...")

	for {
		m, err := reader.ReadMessage(ctx)
		if err != nil {
			fmt.Println("❌ Error reading message:", err)
			continue
		}

		fmt.Printf("Message at offset %d: %s\n", m.Offset, string(m.Value))

		var ticket TicketMessage
		if err := json.Unmarshal(m.Value, &ticket); err != nil {
			fmt.Println("Error decoding message:", err)
			continue
		}

		message := fmt.Sprintf("Your ticket %s is %s.", ticket.TicketID, ticket.Status)
		sendSMS(ticket.UserPhone, message)
	}
}
