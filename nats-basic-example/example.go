package main


import (
	"log"
	"runtime"
	"encoding/json"

	"github.com/nats-io/go-nats"
)

func main() {
	nc, err := nats.Connect("nats://127.0.0.1:4222")
	if err != nil {
		log.Fatalf("Error: %s", err)
	}


	nc.Subscribe("order.created", func(m *nats.Msg) {
		log.Printf("[Orer] %s", string(m.Data))
	})


	payload := struct {
		OrderID string
		Status      string
	}{
		OrderID: "1234-5678-90",
		Status:      "Placed",
	}
	msg, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	nc.Publish("order.created", msg)
	log.Println("Message published")
	runtime.Goexit()

}
