package main

import (
	"context"
	"log"
	"net"

	pb "example.com/go-starwars-grpc/starwars"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type BrokerServer struct {
	pb.UnimplementedStarWarsServer
}

func (s *BrokerServer) AddCity(ctx context.Context, in *pb.NewCity) (*pb.RespBroker1, error) {
	log.Printf("Received: %v", in.GetNombrePlaneta())
	log.Printf("Received: %v", in.GetNombreCiudad())
	log.Printf("Received: %v", in.GetNuevoValor())
	log.Printf("Received: %v", in.GetAction())
	var direccion string = ":localhost/9999"
	return &pb.RespBroker1{DireccionServidor: direccion}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterStarWarsServer(s, &BrokerServer{})
	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
