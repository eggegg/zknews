package main

import (
	"log"
	"net/http"

	// "gopkg.in/mgo.v2"
	
	// "github.com/urfave/negroni"
	"github.com/gorilla/mux"
)

// App is the struct with app configuration values
type App struct {
	Router *mux.Router
}

// Initialize create the DB connection and prepare all the routes
func (a *App) Initialize() {
	a.Router = mux.NewRouter()
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/healthcheck", a.healthcheck).Methods("GET")
}

// Run initialize the server
func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) healthcheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}