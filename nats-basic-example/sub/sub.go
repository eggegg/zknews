package main

import (
	stan "github.com/nats-io/go-nats-streaming"

	"time"
	"log"
	"github.com/eggegg/zknews/nats-utils"

	"runtime"
	
)


const (
	clusterID = "test-cluster"
	clientID  = "news-service-sub"
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
		// Subscribe with manual ack mode, and set AckWait to 60 seconds
		aw, _ := time.ParseDuration("60s")
		// Subscribe the channel
		sc.Subscribe(subscribeChannel, func(msg *stan.Msg) {
			msg.Ack() // Manual ACK
			
			log.Printf("[Received]: %s", string(msg.Data))
			

		}, stan.DurableName(durableID),
			stan.MaxInflight(25),
			stan.SetManualAckMode(),
			stan.AckWait(aw),
		)
		
		runtime.Goexit()
}