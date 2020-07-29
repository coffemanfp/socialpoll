package main

import (
	"log"

	"github.com/bitly/go-nsq"
)

func publishVotes(votes <-chan string) <-chan struct{} {
	stopChan := make(chan struct{}, 1)

	// NSQ conn
	pub, err := nsq.NewProducer("localhost:4150", nsq.NewConfig())
	if err != nil {
		log.Fatalln("Publisher: ", err)
	}
	go func() {
		defer func() {
			stopChan <- struct{}{}
		}()
		for vote := range votes {
			pub.Publish("votes", []byte(vote)) // publish vote
		}

		log.Println("Publisher: Stopping")

		pub.Stop()

		log.Println("Publisher: Stopped")
	}()

	return stopChan
}
