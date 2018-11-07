package main

import (
	stan "github.com/nats-io/go-nats-streaming"
	"log"
	"time"
	"runtime"
	"github.com/eggegg/zknews/nats-utils"
)


const (
	clusterID = "test-cluster"
	clientID  = "news-service-pub"
	subscribeChannel   = "news-created"
	durableID = "news-service-durable"
)


func main()  {
		// Register new NATS component within the system.
		comp := natsutil.NewStreamingComponent(clientID)
	
		// Connect to NATS Streaming server
		err := comp.ConnectToNATSStreaming(
			clusterID,
			stan.NatsURL(stan.DefaultNatsURL),
		)
		if err != nil {
			log.Fatal(err)
		}
		// Get the NATS Streaming Connection
		sc := comp.NATS()

		for range time.NewTicker(500 * time.Millisecond).C {
			err := sc.Publish(subscribeChannel, []byte("hello world!") )
			if err != nil {
				log.Fatalf("Error: %s", err)
			}
			log.Println("published")
		}

		for i :=0; i <= 10; i++ {
			sc.Publish(subscribeChannel, []byte("[Sender] hello world!") )
			time.Sleep(100 * time.Millisecond)
			log.Printf("[Publish] %v \r\n" , i)
		}

		// Publish message on subject (channel)
		

		runtime.Goexit()
}