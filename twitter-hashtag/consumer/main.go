package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
	"twitter-hashtag/consumer/controller"
	database "twitter-hashtag/consumer/db"

	"github.com/segmentio/kafka-go"
)

func main() {

	group := "hashtag-count-store-consumer"
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal("error getting current directory: ", err)
	}

	db, err := database.Connect(fmt.Sprintf("%s/db/schema.sql", dir))
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

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
	var store = make(map[string]int)
	batch := 0
	go func() {
		ticker := time.NewTicker(10 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				if batch > 0 {
					processBatch(db, store)
					batch = 0
				}
			default:
				msg, err := reader.ReadMessage(context.Background())
				if err != nil {
					log.Printf("Error reading message: %v", err)
					continue
				}

				var tweet controller.Tweet
				if err := json.Unmarshal(msg.Value, &tweet); err != nil {
					log.Printf("Error unmarshalling message: %v", err)
					continue
				}

				batch++
				store[tweet.HashTag]++

				if batch >= 1000 {
					processBatch(db, store)
					batch = 0
				}
			}
		}
	}()
	<-signalChannel
	log.Println("Received shutdown signal. Exiting...")
}

var mu sync.Mutex

func processBatch(db *sql.DB, store map[string]int) error {
	var wg sync.WaitGroup
	errChan := make(chan error, len(store))

	mu.Lock()

	for hashtag, count := range store {
		wg.Add(1)
		go func(hashtag string, count int) {
			defer wg.Done()
			if err := controller.IncrementHashtagCount(db, hashtag, count); err != nil {
				errChan <- fmt.Errorf("could not increment hashtag count for %s: %v", hashtag, err)
			}
		}(hashtag, count)
	}

	wg.Wait()
	close(errChan)

	for hashtag := range store {
		store[hashtag] = 0
	}
	
	mu.Unlock()

	for err := range errChan {
		if err != nil {
			log.Println(err)
		}
	}

	return nil
}
