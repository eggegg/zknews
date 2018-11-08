package main

import (
	"context"
	"encoding/json"
	// "errors"
	"log"
	// "net"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	mgo "gopkg.in/mgo.v2"
	_ "gopkg.in/mgo.v2/bson"

	// "google.golang.org/grpc"
	// "google.golang.org/grpc/reflection"
)

// App is the struct with app configuration values
type App struct {
	Session     *mgo.Session
	Router *mux.Router
}

// Initialize create the DB connection and prepare all the routes
func (a *App) Initialize(session *mgo.Session) {
	a.Session = session
	a.Router = mux.NewRouter()
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/healthcheck", a.healthcheck).Methods("GET")
	a.Router.HandleFunc("/news", a.createNews).Methods("POST")
	a.Router.HandleFunc("/news/{id}", a.getNews).Methods("GET") 
	a.Router.HandleFunc("/all", a.allNews).Methods("GET")
	
}

// Run grpc 
func (a *App) RunGRPServer()  {
	var r Repository
	r, err := NewMongoRepository(a.Session)
	if err != nil{
		log.Println(err)
	}
	s := NewNewsService(r)
	log.Println("GRPC Listening on port 8080...")	
	log.Fatal(ListenGRPC(s, 8080))
}

// Run initialize the server
func (a *App) Run(addr string) {
	log.Println("HTTP Listening on port ", addr)		
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) healthcheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
	return
}


func (a *App) createNews(w http.ResponseWriter, r *http.Request) {
	var news News
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&news); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()


	url := "localhost:8080"
	client , err := NewClient(url)
	res, err := client.PostNews(context.Background(), news.Title, news.Content, news.Author, news.NewsType, news.Tags)
	if err != nil {
		respondWithJSON(w, http.StatusOK, err.Error())
		return
	}
	
	respondWithJSON(w, http.StatusOK, res)
	return
}

func (a *App) getNews(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		respondWithError(w, http.StatusBadRequest, "Can not get id")
		return
	}

	url := "localhost:8080"
	client , err := NewClient(url)

	res, err := client.GetNews(context.Background(), id)
	if err != nil{
		respondWithError(w, http.StatusBadRequest, err.Error())
	}

	respondWithJSON(w, http.StatusOK, res)
}

func (a *App) allNews(w http.ResponseWriter, r *http.Request) {
	skip, _ := strconv.Atoi(r.FormValue("skip"))
	take, _ := strconv.Atoi(r.FormValue("take"))

	url := "localhost:8080"
	client , err := NewClient(url)

	res, err := client.GetAllNews(context.Background(), skip, take)
	if err != nil {
		respondWithError(w, http.StatusOK, err.Error())
	}

	respondWithJSON(w, http.StatusOK, res)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
