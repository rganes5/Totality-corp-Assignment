package api

import (
	"fmt"
	"net"
	"user-service/pkg/config"
	"user-service/pkg/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	Gs   *grpc.Server
	Lis  net.Listener
	Port string
}

func NewGrpcServe(cfg *config.Config, service pb.UserServiceServer) (*Server, error) {
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, service)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		return nil, err
	}
	return &Server{
		Gs:   grpcServer,
		Lis:  lis,
		Port: cfg.Port,
	}, nil

}

func (s *Server) Start() error {
	fmt.Println("User service on :", s.Port)
	return s.Gs.Serve(s.Lis)
}
