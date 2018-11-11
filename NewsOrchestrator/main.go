package main

import (
	"flag"
	// "log"
)


func main() {

	var httpPort string
	var SportRPCAddress string
	var SearchPRCAddress string

	flag.StringVar(&httpPort, "http_port", "3100", "http port") 
	
	flag.StringVar(&SportRPCAddress, "Sport rpc address", "localhost:8080", "http port") 
	flag.StringVar(&SearchPRCAddress, "search rpc address", "localhost:50061", "http port") 
	
	
	flag.Parse()


	

	a := App{}
	a.Initialize(SportRPCAddress, SearchPRCAddress)
	a.initializeRoutes()
	a.Run(httpPort)
}
