package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	pb "example.com/go-starwars-grpc/starwars"
	"google.golang.org/grpc"
)

const (
	port_almirante = ":50051"
	port_leia      = ":50050"
)

type BrokerServer struct {
	pb.UnimplementedStarWarsServer
}

func (s *BrokerServer) CityMgmtBroker(ctx context.Context, in *pb.NewCity) (*pb.RespBroker1, error) {
	log.Printf("Received: %v", in.GetNombrePlaneta())
	log.Printf("Received: %v", in.GetNombreCiudad())
	log.Printf("Received: %v", in.GetNuevoValor())
	log.Printf("Received: %v", in.GetAction())
	rand.Seed(int64(time.Now().UnixNano()))
	//direcciones_fulcrum := [3]string{"localhost:50052", "localhost:50053", "localhost:50054"}
	//direccion := direcciones_fulcrum[rand.Intn(3)]
	direccion := "localhost:50052"
	return &pb.RespBroker1{DireccionServidor: direccion}, nil
}
func (s *BrokerServer) CityLeiaBroker(ctx context.Context, in *pb.NewCity) (*pb.RespBroker2, error) {
	log.Printf("Received from Leia: %v", in.GetNombrePlaneta())
	log.Printf("Received from Leia: %v", in.GetNombreCiudad())
	log.Printf("Received from Leia: %v", in.GetAction())
	cant_rebeldes := 99
	reloj_vector := []int32{1, 0, 0}
	servidor_contactado := "localhost:50052"
	return &pb.RespBroker2{CantRebeldes: int32(cant_rebeldes), RelojVector: reloj_vector, ServidorContactado: servidor_contactado}, nil
}

func main() {
	//Conexion con almirante
	lis, err := net.Listen("tcp", port_almirante)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterStarWarsServer(s, &BrokerServer{})
	log.Printf("server listening at %v", lis.Addr())
	//conexion con Leia
	lis1, err1 := net.Listen("tcp", port_leia)
	if err1 != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s1 := grpc.NewServer()
	pb.RegisterStarWarsServer(s1, &BrokerServer{})
	log.Printf("server listening at %v", lis1.Addr())

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}

	}()

	go func() {
		if err := s.Serve(lis1); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	fmt.Println("Para salir del servidor, presione Q y luego Enter")
	var sa string
	fmt.Scan(&sa)
	fmt.Println("Servidor Broker finalizado")
}
