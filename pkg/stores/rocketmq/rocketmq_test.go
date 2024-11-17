package rocketmq

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-clients/golang/v5"
	"strconv"
	"testing"
	"time"
)

const (
	Topic     = "testTopic"
	Endpoint  = "192.168.6.218:9876"
	AccessKey = "xxxxxx"
	SecretKey = "xxxxxx"
)

func TestPublisher(t *testing.T) {
	// In most case, you don't need to create many producers, singleton pattern is more recommended.
	producer, err := golang.NewProducer(
		&golang.Config{
			Endpoint: Endpoint,
		},
		golang.WithTopics(Topic),
	)
	if err != nil {
		panic(err)
	}
	// start producer
	err = producer.Start()
	if err != nil {
		panic(err)
	}

	// graceful stop producer
	defer producer.GracefulStop()

	for i := 0; i < 10; i++ {
		// new a message
		msg := &golang.Message{
			Topic: Topic,
			Body:  []byte("this is a message : " + strconv.Itoa(i)),
		}
		// set keys and tag
		msg.SetKeys("a", "b")
		msg.SetTag("ab")
		// send message in sync
		resp, err := producer.Send(context.TODO(), msg)
		if err != nil {
			fmt.Println(err)
		}
		for i := 0; i < len(resp); i++ {
			fmt.Printf("%#v\n", resp[i])
		}
		// wait a moment
		time.Sleep(time.Second * 1)
	}
}
