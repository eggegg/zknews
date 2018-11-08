//go:generate protoc ./news.proto --go_out=plugins=grpc:./pb
package main

import (
	"context"
	"fmt"
	"net"

	"github.com/eggegg/zknews/SportNewsService/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	service Service
	
}

func ListenGRPC(s Service, port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil
	}
	serv := grpc.NewServer()
	pb.RegisterNewsServiceServer(serv, &grpcServer{s})
	reflection.Register(serv)
	return serv.Serve(lis)
}

func (s *grpcServer) PostNews(ctx context.Context, r *pb.PostNewsRequest) (*pb.PostNewsResponse, error) {
	a, err := s.service.PostNews(ctx, r.Title, r.Content, r.Author, r.NewsType, r.Tags)
	if err != nil {
		return nil, err
	}
	return &pb.PostNewsResponse{News: &pb.News{
		Id: a.ID.Hex(),
		Title: a.Title,
		Content: a.Content,
		Author: a.Author,
		NewsType: a.NewsType,
		Tags: a.Tags,
	}}, nil
}

func (s *grpcServer) GetNews(ctx context.Context, r *pb.GetNewsRequest) (*pb.GetNewsResponse, error) {
	a, err := s.service.GetNews(ctx, r.Id)
	if err != nil {
		return nil, err
	}
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