package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"twitter-hashtag/consumer/controller"
	database "twitter-hashtag/consumer/db"

	"github.com/segmentio/kafka-go"
)

func main() {

	group := os.Args[1]
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal("error getting current directory: ", err)
	}

	db, err := database.Connect(fmt.Sprintf("%s/db/schema.sql", dir))

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)

	broker := "localhost:29092"
	topic := "hashtag-topic"

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker},
		Topic:   topic,
		GroupID: group,
	})

	defer reader.Close()

	go func() {
		for {
			// Fetch a message
			msg, err := reader.ReadMessage(context.Background())
			if err != nil {
				log.Printf("Error reading message: %v", err)
				continue
			}
			var tweet controller.Tweet
			err = json.Unmarshal(msg.Value, &tweet)
			if err != nil {
				log.Printf("Error unmarshalling message: %v", err)
				continue
			}

			tweet.IncreamentHashtagCount(db)
		}
	}()
	<-signalChannel
	log.Println("Received shutdown signal. Exiting...")
}
