//go:generate protoc ./news.proto --go_out=plugins=grpc:./pb
package main

import (
	"context"
	"fmt"
	"net"
	"encoding/json"
	"log"

	"github.com/satori/go.uuid"
	"github.com/pkg/errors"

	"github.com/eggegg/zknews/SportNewsService/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	NewsCreatedEvent = "news-created"	
	AggregateType = "news"
)

type grpcServer struct {
	service Service
	eventURL string
}

func ListenGRPC(s Service, port int, eventURL string) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("connect to grpc:", eventURL)

	serv := grpc.NewServer()
	pb.RegisterNewsServiceServer(serv, &grpcServer{s, eventURL})
	reflection.Register(serv)
	return serv.Serve(lis)
}

func (s *grpcServer) PostNews(ctx context.Context, r *pb.PostNewsRequest) (*pb.PostNewsResponse, error) {
	a, err := s.service.PostNews(ctx, r.Title, r.Content, r.Author, r.NewsType, r.Tags)
	if err != nil {
		return nil, err
	}

	go s.CreatePostNewsGrpcEvent(a)

	return &pb.PostNewsResponse{News: &pb.News{
		Id: a.ID.Hex(),
		Title: a.Title,
		Content: a.Content,
		Author: a.Author,
		NewsType: a.NewsType,
		Tags: a.Tags,
	}}, nil
}

// create event 
func (s *grpcServer) CreatePostNewsGrpcEvent(news *News) error {
	conn, err := grpc.Dial(s.eventURL, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Unable to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewEventStoreClient(conn)
	eventJson, _ := json.Marshal(news)

	newEventId, _ := uuid.NewV4()
	event := &pb.Event{
		EventId:    newEventId.String(),
		EventType:     NewsCreatedEvent,
		AggregateId:   news.NewsType,
		AggregateType: AggregateType,
		EventData:     string(eventJson),
		Channel:       NewsCreatedEvent,
	}
	resp, err := client.CreateEvent(context.Background(), event)
	if err != nil {
		return errors.Wrap(err, "Error from RPC server")		
	}
	if resp.IsSuccess {
		return nil
	} else {
		return errors.Wrap(err, "Error from RPC server")
	}
	
}

func (s *grpcServer) GetNews(ctx context.Context, r *pb.GetNewsRequest) (*pb.GetNewsResponse, error) {
	a, err := s.service.GetNews(ctx, r.Id)
	if err != nil {
		return nil, err
	}

	// send to eventstore


	return &pb.GetNewsResponse{
		News: &pb.News{
			Id: a.ID.Hex(),
			Title: a.Title,
			Content: a.Content,
			Author: a.Author,
			NewsType: a.NewsType,
			Tags: a.Tags,
		},
	},nil
}

func (s *grpcServer) GetAllNews(ctx context.Context, r *pb.GetAllNewsRequest) (*pb.GetAllNewsResponse, error) {
	res, err := s.service.GetAllNews(ctx, int(r.Skip), int(r.Take) )
	if err != nil {
		return nil, err
	}
	newsList := []*pb.News{}
	for _, a := range res {
		newsList = append(
			newsList,
			&pb.News{
				Id: a.ID.Hex(),
				Title: a.Title,
				Content: a.Content,
				Author: a.Author,
				NewsType: a.NewsType,
				Tags: a.Tags,
			},
		)
	}
	return &pb.GetAllNewsResponse{Allnews: newsList}, nil
}

func (s *grpcServer) SearchNews(ctx context.Context, r *pb.SearchNewsRequest) (*pb.SearchNewsResponse, error) {
	return &pb.SearchNewsResponse{}, nil
}