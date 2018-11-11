package main


import (
	"context"
	"encoding/json"
	// "errors"
	"log"
	// "net"
	"net/http"
	// "strconv"
	"time"

	"github.com/gorilla/mux"

	// mgo "gopkg.in/mgo.v2"
	// _ "gopkg.in/mgo.v2/bson"

	stan "github.com/nats-io/go-nats-streaming"	
	"github.com/eggegg/zknews/natsutils"
	
	"github.com/eggegg/zknews/SearchNewsService/pb"
	
	// "google.golang.org/grpc"
	// "google.golang.org/grpc/reflection"
)


const (
	NewsCreatedEvent = "news-created"	
	AggregateType = "news"
	DurableID  = "news-search-durable"
	
)

// App is the struct with app configuration values
type App struct {
	service Service
	comp  *natsutil.StreamingComponent
	Router *mux.Router
}

// Initialize create the DB connection and prepare all the routes
func (a *App) Initialize(s Service, component *natsutil.StreamingComponent) {
	a.service = s
	a.comp = component
	a.Router = mux.NewRouter()
}

func  (a *App) StartListenNATS() {
	// Get the NATS Streaming Connection
	sc := a.comp.NATS()
	// Subscribe with manual ack mode, and set AckWait to 60 seconds
	aw, _ := time.ParseDuration("60s")

	// Subscribe the channel
	sc.Subscribe(NewsCreatedEvent, func(msg *stan.Msg) {
		msg.Ack() // Manual ACK
		pbNews := pb.News{}
		// Unmarshal JSON that represents the data
		err := json.Unmarshal(msg.Data, &pbNews)
		if err != nil {
			log.Print(err)
			return
		}

		news := &News{
			ID: pbNews.Id,
			Title : pbNews.Title,
			Content : pbNews.Content,
			Author :pbNews.Author,
			NewsType : pbNews.NewsType,
			Tags : pbNews.Tags,
			CreatedOn : time.Now(),
		}

		// Handle the message
		err = a.service.InsertNews(context.Background(), news)
		if err != nil {
			log.Fatalf("error listen to nats: %v", err.Error())
		}

		log.Printf("News insert success: %v\n", news)
		

	}, stan.DurableName(DurableID),
		stan.MaxInflight(25),
		stan.SetManualAckMode(),
		stan.AckWait(aw),
	)

}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/healthcheck", a.healthcheck).Methods("GET")
	// a.Router.HandleFunc("/news", a.createNews).Methods("POST")
	// a.Router.HandleFunc("/news/{id}", a.getNews).Methods("GET") 
	// a.Router.HandleFunc("/all", a.allNews).Methods("GET")
	
}

// Run grpc 
func (a *App) RunGRPServer(port string)  {
	// var r Repository
	// r, err := New(a.service)
	// if err != nil{
	// 	log.Println(err)
	// }
	// s := NewNewsService(r)
	log.Println("GRPC Listening on port:", port)	
	log.Fatal(ListenGRPC(a.service, port))
}

// Run initialize the server
func (a *App) Run(addr string) {
	log.Println("HTTP Listening on port ", addr)		
	log.Fatal(http.ListenAndServe(":"+addr, a.Router))
}

func (a *App) healthcheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
	return
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
