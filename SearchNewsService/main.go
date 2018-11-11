package main

import (
	// "time"
	"log"
	"flag"
	// "os"

	
	// "github.com/olivere/elastic"

	stan "github.com/nats-io/go-nats-streaming"		
	"github.com/eggegg/zknews/natsutils"
	

)

const (
	clusterID = "test-cluster"
	clientID  = "search-news"
)


func main() {
	
	var elasticsearchAddress string
	var nataddress string
	var httpPort string
	var grpcPort string

	flag.StringVar(&httpPort, "http_port", "3001", "http port") 
	flag.StringVar(&grpcPort, "grpc_port", "50061", "grpc port") 
	
	
	flag.StringVar(&elasticsearchAddress, "elasticsearch_address", "http://localhost:9200", "Elasticsearch Address") 
	//os.Getenv("QUERYBD_HOST")
	flag.StringVar(&nataddress, "nats_address", "nats://localhost:4222", "nats_address")
	
	flag.Parse()
	

	// 1. Create service
	// Connect to Elasticsearch
	var r Repository
	r, err := NewElasticsearchRepository(elasticsearchAddress)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	s := NewSearchService(r)

	// 2. Connect to Nats
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
	a.Initialize(s, comp)
	a.StartListenNATS() // start listen to nats
	a.initializeRoutes()
	go a.RunGRPServer(grpcPort)
	a.Run(httpPort)
}
