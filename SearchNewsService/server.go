package main

import (
	"context"
	"fmt"
	"net"
	// "log"
	// "encoding/json"

	// "github.com/satori/go.uuid"
	// "github.com/pkg/errors"

	"github.com/eggegg/zknews/SearchNewsService/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)


type grpcServer struct {
	service Service
}


func ListenGRPC(s Service, port string) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		return nil
	}
	serv := grpc.NewServer()
	pb.RegisterNewsServiceServer(serv, &grpcServer{s})
	reflection.Register(serv)
	return serv.Serve(lis)
}

func (s *grpcServer) SearchNews(ctx context.Context, r *pb.SearchNewsRequest) (*pb.SearchNewsResponse, error) {
	res, err := s.service.SearchNews(ctx, r.Query, int(r.Skip), int(r.Take) )
	if err != nil {
		return nil, err
	}

	newsList := []*pb.News{}
	for _, a := range res {
		newsList = append(
			newsList,
			&pb.News{
				Id: a.ID,
				Title: a.Title,
				Content: a.Content,
				Author: a.Author,
				NewsType: a.NewsType,
				Tags: a.Tags,
			},
		)
	}
	return &pb.SearchNewsResponse{Allnews: newsList}, nil
}

func (s *grpcServer) GetNews(ctx context.Context, r *pb.GetNewsRequest) (*pb.GetNewsResponse, error) {
	return &pb.GetNewsResponse{},nil
}	
	
func (s *grpcServer) GetAllNews(ctx context.Context, r *pb.GetAllNewsRequest) (*pb.GetAllNewsResponse, error) {
	return &pb.GetAllNewsResponse{},nil
}

func (s *grpcServer) PostNews(ctx context.Context, r *pb.PostNewsRequest) (*pb.PostNewsResponse, error) {
	return &pb.PostNewsResponse{},nil
}