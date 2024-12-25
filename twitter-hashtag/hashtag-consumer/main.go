package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"twitter-hashtag/consumer/consumer"
	database "twitter-hashtag/consumer/db"
)

func main() {

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	topic := "hashtag-topic"
	db, err := database.Connect(fmt.Sprintf("%s/db/schema.sql", dir))
	if err != nil {
		log.Fatal(err)
	}
	go consumer.Consumer(topic, 0)
	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)
	<-quitChannel
	defer db.Close()
}
