package main

import (
	"context"
	"log"
	"net"

	pb "example.com/go-starwars-grpc/starwars2"
	"google.golang.org/grpc"
)

const (
	port = ":50052"
)

type FulcrumServer struct {
	pb.UnimplementedStarWars1Server
}

func (s *FulcrumServer) CityMgmtFulcrum(ctx context.Context, in *pb.NewCity1) (*pb.RespFulcrum1, error) {
	log.Printf("Received: %v", in.GetNombrePlaneta())
	log.Printf("Received: %v", in.GetNombreCiudad())
	log.Printf("Received: %v", in.GetNuevoValor())
	log.Printf("Received: %v", in.GetAction())
	reloj_vector := [3]int{1, 2, 3}
	return &pb.RespFulcrum1{RelojVector1: int32(reloj_vector[0]), RelojVector2: int32(reloj_vector[1]), RelojVector3: int32(reloj_vector[2])}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterStarWars1Server(s, &FulcrumServer{})
	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
