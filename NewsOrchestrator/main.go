package main

import (
	"log"
)


func main() {
		
	// connectionString := os.Getenv("DATABASE_URL")

	// db, err := sqlx.Open("postgres", connectionString)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	log.Println("start service....")
	

	a := App{}
	a.Initialize()
	a.initializeRoutes()
	a.Run(":3000")
}
