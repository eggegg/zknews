package main

import (
	"time"
	// "log"
	"flag"
	// "os"

	mgo "gopkg.in/mgo.v2"

)


func main() {
	
	var connectionString string
	flag.StringVar(&connectionString, "querydb_host", "localhost:27017", "Mongo Address") 
	//os.Getenv("QUERYBD_HOST")

	var eventURL string
	flag.StringVar(&eventURL, "event_url","localhost:50051", "event store grpc url")
	
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

	a := App{}
	a.Initialize(session)
	a.initializeRoutes()
	go a.RunGRPServer(eventURL)
	a.Run(":3000")
}
