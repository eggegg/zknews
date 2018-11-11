package main

import (
	"context"
	"log"
	"net/http"

	"encoding/json"
	"strconv"

	// "gopkg.in/mgo.v2"
	
	// "github.com/urfave/negroni"
	"github.com/gorilla/mux"

	"github.com/eggegg/zknews/NewsOrchestrator/pb"

	"google.golang.org/grpc"
	// "google.golang.org/grpc/reflection"
)

// App is the struct with app configuration values
type App struct {
	Router *mux.Router
	SportRPCAddress string
	SearchRPCAddress string
}

// Initialize create the DB connection and prepare all the routes
func (a *App) Initialize(SportRPCAddress, SearchRPCAddress string) {
	a.Router = mux.NewRouter()
	a.SportRPCAddress = SportRPCAddress
	a.SearchRPCAddress = SearchRPCAddress
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/healthcheck", a.healthcheck).Methods("GET")

	a.Router.HandleFunc("/news", a.createNews).Methods("POST")
	a.Router.HandleFunc("/news/{newsType}/{id}", a.getNews).Methods("GET") 
	a.Router.HandleFunc("/all", a.allNews).Methods("GET")
	a.Router.HandleFunc("/search", a.searchNews).Methods("GET")
}

// Run initialize the server
func (a *App) Run(addr string) {
	log.Println("start listen http on port:", addr)
	log.Fatal(http.ListenAndServe(":"+addr, a.Router))
}

func (a *App) healthcheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func (a *App) createNews(w http.ResponseWriter, r *http.Request) {
	var news News
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&news); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	
	// call sport news rpc
	if (news.NewsType == "sports") {
		url := a.SportRPCAddress
		conn, err := grpc.Dial(url, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Unable to connect: %v", err)
		}
		defer conn.Close()
	
		client := pb.NewNewsServiceClient(conn)


		ret, err := client.PostNews(context.Background(), 
			&pb.PostNewsRequest{
			Title: news.Title,
			Content: news.Content,
			Author: news.Author,
			NewsType: news.NewsType,
			Tags: news.Tags,
		})
		if err != nil {
			respondWithJSON(w, http.StatusOK, err.Error())
			return
		}

		news.ID = ret.News.Id
		respondWithJSON(w, http.StatusOK, news)
	}
	
	// respondWithError(w, http.StatusServiceUnavailable, "error occurs")
	return
}

func (a *App) getNews(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		respondWithError(w, http.StatusBadRequest, "Can not get id")
		return
	}

	newsType, ok := vars["newsType"]
	if !ok {
		respondWithError(w, http.StatusBadRequest, "Can not get news type")
		return
	}

	// call sport news rpc
	if (newsType == "sports") {
		url := a.SportRPCAddress
		conn, err := grpc.Dial(url, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Unable to connect: %v", err)
		}
		defer conn.Close()
	
		client := pb.NewNewsServiceClient(conn)


		ret, err := client.GetNews(context.Background(), 
			&pb.GetNewsRequest{
				NewsType: newsType,
				Id: id,
		})
		if err != nil {
			respondWithJSON(w, http.StatusOK, err.Error())
			return
		}

		news := &News{
			ID: ret.News.Id,
			Title: ret.News.Title,
			Content: ret.News.Content,
			Author: ret.News.Content,
			Tags: ret.News.Tags,
		}
		respondWithJSON(w, http.StatusOK, news)
	}

	return

}	

func (a *App) allNews(w http.ResponseWriter, r *http.Request) {
	skip, _ := strconv.Atoi(r.FormValue("skip"))
	take, _ := strconv.Atoi(r.FormValue("take"))


	// call news rpc
	url := a.SportRPCAddress
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Unable to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewNewsServiceClient(conn)


	ret, err := client.GetAllNews(context.Background(), &pb.GetAllNewsRequest{
		NewsType : "sports",
		Skip: int32(skip),
		Take: int32(take),
	})
	if err != nil {
		respondWithJSON(w, http.StatusOK, err.Error())
		return
	}

	allNews :=  []News{}
	for _, news := range ret.Allnews {
		n := News{
			ID: news.Id,
			Title: news.Title,
			Content: news.Content,
			Author: news.Author,
			Tags: news.Tags,
		}
		allNews = append(allNews, n)
	}

	respondWithJSON(w, http.StatusOK, allNews)

	return
}	

func (a *App) searchNews(w http.ResponseWriter, r *http.Request) {

	query := r.FormValue("query")
	skip, _ := strconv.Atoi(r.FormValue("skip"))
	take, _ := strconv.Atoi(r.FormValue("take"))


	url := a.SearchRPCAddress
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Unable to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewNewsServiceClient(conn)


	ret, err := client.SearchNews(context.Background(), &pb.SearchNewsRequest{
		Query : query,
		Skip: int32(skip),
		Take: int32(take),
	})
	if err != nil {
		respondWithJSON(w, http.StatusOK, err.Error())
		return
	}

	allNews :=  []News{}
	for _, news := range ret.Allnews {
		n := News{
			ID: news.Id,
			Title: news.Title,
			Content: news.Content,
			Author: news.Author,
			Tags: news.Tags,
		}
		allNews = append(allNews, n)
	}

	respondWithJSON(w, http.StatusOK, allNews)

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