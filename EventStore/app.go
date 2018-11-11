package main

import (
	"context"
	"net"
	"log"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/eggegg/zknews/natsutils"
	"github.com/eggegg/zknews/EventStore/pb"
	
	
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type EventStore struct{
	Id bson.ObjectId `bson:"_id,omitempty"`
	EventId string `bson:"event_id"`
	EventType string `bson:"event_type"`
	AggregateId string `bson:"aggregate_id"`
	AggregateType string `bson:"aggregate_type"`
	EventData string `bson:"event_data"`
	Channel string `bson:"channel"` 
	CreatedOn  time.Time  `bson:"createdon,omitempty"`		
}

// App is the struct with app configuration values
type App struct {
	Session     *mgo.Session
	Component   *natsutil.StreamingComponent
}

// Initialize create the DB connection and prepare all the routes
func (a *App) Initialize(session *mgo.Session, component *natsutil.StreamingComponent) {
	a.Session = session
	a.Component = component
}

func (a *App) CreateEvent(event *pb.Event) error {
	// insert into mongo
	session := a.Session.Copy()
	defer session.Close()

	collection := session.DB("zknews").C("eventstore")

	n := EventStore{
		Id: bson.NewObjectId(),
		EventId: event.EventId,
		EventType: event.EventType,
		AggregateId: event.AggregateId,
		AggregateType: event.AggregateType,
		EventData: event.EventData,
		Channel: event.Channel,
		CreatedOn: time.Now(),
	}

	err := collection.Insert(n)
	return err
}

func (a *App) GetEvents(filter *pb.EventFilter) []*pb.Event {
	var events []*pb.Event
	return events
}


type grpcHandler struct {
	app *App
}

func (s *grpcHandler) CreateEvent(ctx context.Context, in *pb.Event) (*pb.Response, error) {

	err := s.app.CreateEvent(in)
	if err != nil {
		return nil, err
	}

	// Publish event on NATS Streaming Server
	go publishEvent(s.app.Component, in, )
	return &pb.Response{IsSuccess: true}, nil
	
}

func (s *grpcHandler) GetEvents(ctx context.Context, in *pb.EventFilter) (*pb.EventResponse, error) {
	events := s.app.GetEvents(in)
	return &pb.EventResponse{Events: events}, nil
}

// Run grpc 
func (a *App) RunGRPServer(port string)  {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterEventStoreServer(s, &grpcHandler{app: a})
	reflection.Register(s)

	log.Println("GRPC Listening on port ", port)	
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}


// publishEvent publishes an event via NATS Streaming server
func publishEvent(component *natsutil.StreamingComponent, event *pb.Event) {
	sc := component.NATS()
	channel := event.Channel
	eventMsg := []byte(event.EventData)
	// Publish message on subject (channel)
	sc.Publish(channel, eventMsg)
	log.Println("Published message on channel: " + channel)
}
