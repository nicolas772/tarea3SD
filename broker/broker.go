package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	pb "example.com/go-starwars-grpc/starwars"
	pb1 "example.com/go-starwars-grpc/starwars2"
	"google.golang.org/grpc"
)

const (
	port_almirante = ":50051"
	port_leia      = ":50050"
	port_ahsoka    = ":50054"
)

type BrokerServer struct {
	pb.UnimplementedStarWarsServer
}

func (s *BrokerServer) CityMgmtBroker(ctx context.Context, in *pb.NewCity) (*pb.RespBroker1, error) {
	log.Printf("Received from %v: %v", in.GetSender(), in.GetNombrePlaneta())
	log.Printf("Received from %v: %v", in.GetSender(), in.GetNombreCiudad())
	log.Printf("Received from %v: %v", in.GetSender(), in.GetNuevoValor())
	log.Printf("Received from %v: %v", in.GetSender(), in.GetAction())
	rand.Seed(int64(time.Now().UnixNano()))
	direcciones_fulcrum_ahsoka := [3]string{"localhost:50055", "localhost:50056", "localhost:50057"}
	direcciones_fulcrum_almirante := [3]string{"localhost:50058", "localhost:50059", "localhost:50060"}
	var direccion string
	if in.GetSender() == "almirante" {
		direccion = direcciones_fulcrum_almirante[rand.Intn(3)]
	} else {
		direccion = direcciones_fulcrum_ahsoka[rand.Intn(3)]
	}
	//este debe ser aleatorio
	return &pb.RespBroker1{DireccionServidor: direccion}, nil
}
func (s *BrokerServer) CityLeiaBroker(ctx context.Context, in *pb.NewCity) (*pb.RespBroker2, error) {
	log.Printf("Received from Leia: %v", in.GetNombrePlaneta())
	log.Printf("Received from Leia: %v", in.GetNombreCiudad())
	log.Printf("Received from Leia: %v", in.GetAction())
	planeta := in.GetNombrePlaneta()
	ciudad := in.GetNombreCiudad()
	action := in.GetAction()
	direcciones_fulcrum_broker := [3]string{"localhost:50061", "localhost:50062", "localhost:50063"}
	servidor_contactado := direcciones_fulcrum_broker[rand.Intn(3)]
	fmt.Println("servidor contactado para leia:", servidor_contactado)
	conn, err := grpc.Dial(servidor_contactado, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb1.NewStarWars1Client(conn)
	ctx1, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.CityBrokerFulcrum(ctx1, &pb1.NewCity1{NombrePlaneta: planeta, NombreCiudad: ciudad, Action: action})
	if err != nil {
		log.Fatalln(err)
	}
	cant_rebeldes := r.GetCantRebeldes()
	reloj_vector := r.GetRelojVector()
	return &pb.RespBroker2{CantRebeldes: int32(cant_rebeldes), RelojVector: reloj_vector, ServidorContactado: servidor_contactado}, nil
}

func main() {
	//Conexion con almirante
	go func() {
		lis, err := net.Listen("tcp", port_almirante)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()
		pb.RegisterStarWarsServer(s, &BrokerServer{})
		log.Printf("server listening at %v", lis.Addr())
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}

	}()
	//Conexion con Leia
	go func() {
		lis1, err1 := net.Listen("tcp", port_leia)
		if err1 != nil {
			log.Fatalf("failed to listen: %v", err1)
		}

		s1 := grpc.NewServer()
		pb.RegisterStarWarsServer(s1, &BrokerServer{})
		log.Printf("server listening at %v", lis1.Addr())
		if err := s1.Serve(lis1); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	//conexion con ahsoka
	go func() {
		lis2, err2 := net.Listen("tcp", port_ahsoka)
		if err2 != nil {
			log.Fatalf("failed to listen: %v", err2)
		}

		s2 := grpc.NewServer()
		pb.RegisterStarWarsServer(s2, &BrokerServer{})
		log.Printf("server listening at %v", lis2.Addr())
		if err := s2.Serve(lis2); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	fmt.Println("Para salir del servidor, presione Q y luego Enter")
	var sa string
	fmt.Scan(&sa)
	fmt.Println("Servidor Broker finalizado")
}
