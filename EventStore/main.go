// gRPC API for Event Store
package main

import (
	// "context"
	"log"
	// "net"
	"flag"

	"github.com/nats-io/go-nats-streaming"
	mgo "gopkg.in/mgo.v2"
	"time"
	// "google.golang.org/grpc"

	"github.com/eggegg/zknews/natsutils"
)

const (
	port      = ":50051"
	clusterID = "test-cluster"
	clientID  = "event-store-api"
)

func main() {
	var connectionString string
	var nataddress string
	flag.StringVar(&connectionString, "querydb_host", "localhost:27017", "Mongo Address") 
	//os.Getenv("QUERYBD_HOST")
	flag.StringVar(&nataddress, "nats_address", "nats://localhost:4222", "nats_address")
	
	flag.Parse()


	host := []string{
		connectionString,
	}
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs: host,
		Direct: true,
		Timeout: 1 * time.Second,
	})
	session.SetMode(mgo.Monotonic, true)
	if err != nil{
		panic(err)
	}
	defer session.Close()

	// Register new component within the NATS system.
	comp := natsutil.NewStreamingComponent(clientID)
	
	// Connect to NATS
	err = comp.ConnectToNATSStreaming(
		clusterID,
		stan.NatsURL(nataddress),
	)
	if err != nil {
		log.Fatal(err)
	}


	a := App{}
	a.Initialize(session, comp)
	a.RunGRPServer(port)

	
	// // Creates a new gRPC server
	// s := grpc.NewServer()
	// pb.RegisterEventStoreServer(s, &server { StreamingComponent: comp})
	// s.Serve(lis)
}
