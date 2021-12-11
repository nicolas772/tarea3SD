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

func PreguntarRelojFul1(planeta string) []int32 {
	direccion := "localhost:50061"

	conn, err := grpc.Dial(direccion, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb1.NewStarWars1Client(conn)
	ctx1, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.RelojesBrokerFulcrum(ctx1, &pb1.Planeta{NombrePlaneta: planeta})
	if err != nil {
		log.Fatalln(err)
	}
	reloj_vector_resp := r.GetRelojVector()
	return reloj_vector_resp
}
func PreguntarRelojFul2(planeta string) []int32 {
	direccion := "localhost:50062"

	conn, err := grpc.Dial(direccion, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb1.NewStarWars1Client(conn)
	ctx1, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.RelojesBrokerFulcrum(ctx1, &pb1.Planeta{NombrePlaneta: planeta})
	if err != nil {
		log.Fatalln(err)
	}
	reloj_vector_resp := r.GetRelojVector()
	return reloj_vector_resp
}
func PreguntarRelojFul3(planeta string) []int32 {
	direccion := "localhost:50063"

	conn, err := grpc.Dial(direccion, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb1.NewStarWars1Client(conn)
	ctx1, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.RelojesBrokerFulcrum(ctx1, &pb1.Planeta{NombrePlaneta: planeta})
	if err != nil {
		log.Fatalln(err)
	}
	reloj_vector_resp := r.GetRelojVector()
	return reloj_vector_resp
}

func revisarConsistencia(reloj_fulcrum1 []int32, reloj_fulcrum2 []int32, reloj_fulcrum3 []int32) bool {
	doMerge := false
	//comparo fulcrum 1 con fulcrum 2
	var aux [3]int // si el reloj fulcrum1 es mayor en todos los indices al de fulcrum 2, aux : [1,1,1]
	for i := range reloj_fulcrum1 {
		if reloj_fulcrum1[i] > reloj_fulcrum2[i] {
			aux[i] = 1 //1 si es mayor, 0 si es menor, 2 si es igual
		} else if reloj_fulcrum1[i] < reloj_fulcrum2[i] {
			aux[i] = 0
		} else if reloj_fulcrum1[i] == reloj_fulcrum2[i] {
			aux[i] = 2
		}
	}
	encontro_mayor_cruzado1 := false
	for _, val := range aux {
		if val != 2 {
			for _, val2 := range aux {
				if val2 != 2 {
					if val != val2 {
						encontro_mayor_cruzado1 = true
					}
				}
			}
			break
		}

	}
	//comparo fulcrum 1 con fulcrum 3
	var aux2 [3]int
	for i := range reloj_fulcrum1 {
		if reloj_fulcrum1[i] > reloj_fulcrum3[i] {
			aux2[i] = 1 //1 si es mayor, 0 si es menor, 2 si es igual
		} else if reloj_fulcrum1[i] < reloj_fulcrum3[i] {
			aux2[i] = 0
		} else if reloj_fulcrum1[i] == reloj_fulcrum3[i] {
			aux2[i] = 2
		}
	}
	encontro_mayor_cruzado2 := false
	for _, val := range aux2 {
		if val != 2 {
			for _, val2 := range aux2 {
				if val2 != 2 {
					if val != val2 {
						encontro_mayor_cruzado2 = true
					}
				}
			}
			break
		}

	}
	//comparo fulcrum 2 con fulcrum 3
	var aux3 [3]int
	for i := range reloj_fulcrum2 {
		if reloj_fulcrum2[i] > reloj_fulcrum3[i] {
			aux3[i] = 1 //1 si es mayor, 0 si es menor, 2 si es igual
		} else if reloj_fulcrum2[i] < reloj_fulcrum3[i] {
			aux3[i] = 0
		} else if reloj_fulcrum2[i] == reloj_fulcrum3[i] {
			aux3[i] = 2
		}
	}
	encontro_mayor_cruzado3 := false
	for _, val := range aux3 {
		if val != 2 {
			for _, val2 := range aux3 {
				if val2 != 2 {
					if val != val2 {
						encontro_mayor_cruzado3 = true
					}
				}
			}
			break
		}

	}

	if encontro_mayor_cruzado1 || encontro_mayor_cruzado2 || encontro_mayor_cruzado3 {
		doMerge = true
	}
	return doMerge
}

func getCandidato(sender string, dir_ultimo_servidor string, reloj_from_informante []int32, reloj_fulcrum1 []int32, reloj_fulcrum2 []int32, reloj_fulcrum3 []int32) string {
	candidato := ""
	rand.Seed(int64(time.Now().UnixNano()))
	direcciones_fulcrum_ahsoka := [3]string{"localhost:50055", "localhost:50056", "localhost:50057"}
	direcciones_fulcrum_almirante := [3]string{"localhost:50058", "localhost:50059", "localhost:50060"}
	index_a_comparar := 0
	var candidatos []string
	if sender == "almirante" { //si el sender es almirante
		for i, direccion := range direcciones_fulcrum_almirante {
			if direccion == dir_ultimo_servidor {
				index_a_comparar = i
			}
		}
		if reloj_fulcrum1[index_a_comparar] >= reloj_from_informante[index_a_comparar] {
			candidatos = append(candidatos, direcciones_fulcrum_almirante[0])
		}
		if reloj_fulcrum2[index_a_comparar] >= reloj_from_informante[index_a_comparar] {
			candidatos = append(candidatos, direcciones_fulcrum_almirante[1])
		}
		if reloj_fulcrum3[index_a_comparar] >= reloj_from_informante[index_a_comparar] {
			candidatos = append(candidatos, direcciones_fulcrum_almirante[2])
		}
	} else { //si el sender es ahsoka
		for i, direccion := range direcciones_fulcrum_ahsoka {
			if direccion == dir_ultimo_servidor {
				index_a_comparar = i
			}
		}
		if reloj_fulcrum1[index_a_comparar] >= reloj_from_informante[index_a_comparar] {
			candidatos = append(candidatos, direcciones_fulcrum_ahsoka[0])
		}
		if reloj_fulcrum2[index_a_comparar] >= reloj_from_informante[index_a_comparar] {
			candidatos = append(candidatos, direcciones_fulcrum_ahsoka[1])
		}
		if reloj_fulcrum3[index_a_comparar] >= reloj_from_informante[index_a_comparar] {
			candidatos = append(candidatos, direcciones_fulcrum_ahsoka[2])
		}
	}
	fmt.Println("Candidatos:", candidatos)
	if len(candidatos) != 0 {
		candidato = candidatos[rand.Intn(len(candidatos))]
	}
	return candidato
}

func (s *BrokerServer) CityMgmtBroker(ctx context.Context, in *pb.NewCity) (*pb.RespBroker1, error) {
	log.Printf("Received from %v: %v", in.GetSender(), in.GetNombrePlaneta())
	log.Printf("Received from %v: %v", in.GetSender(), in.GetNombreCiudad())
	log.Printf("Received from %v: %v", in.GetSender(), in.GetNuevoValor())
	log.Printf("Received from %v: %v", in.GetSender(), in.GetAction())
	log.Printf("Received from %v: %v", in.GetSender(), in.GetRelojVector())
	var direccion string
	reloj_from_informante := in.GetRelojVector()
	ultimo_servidor := in.GetUltimoServidor()
	reloj_fulcrum1 := PreguntarRelojFul1(in.GetNombrePlaneta())
	reloj_fulcrum2 := PreguntarRelojFul2(in.GetNombrePlaneta())
	reloj_fulcrum3 := PreguntarRelojFul3(in.GetNombrePlaneta())
	fmt.Println("reloj from informante: ", reloj_from_informante)
	fmt.Println("ultimo servidor from informante: ", ultimo_servidor)
	fmt.Println("reloj fulcrum 1: ", reloj_fulcrum1)
	fmt.Println("reloj fulcrum 2: ", reloj_fulcrum2)
	fmt.Println("reloj fulcrum 3: ", reloj_fulcrum3)
	doMerge := revisarConsistencia(reloj_fulcrum1, reloj_fulcrum2, reloj_fulcrum3)
	fmt.Println("do merge: ", doMerge)
	sender := in.GetSender()
	if !doMerge { //si no hay que hacer merge, hay que revisar candidatos
		//direccion = direcciones_fulcrum_almirante[rand.Intn(3)]
		direccion = getCandidato(sender, ultimo_servidor, reloj_from_informante, reloj_fulcrum1, reloj_fulcrum2, reloj_fulcrum3)
	} else {
		direccion = ""
		fmt.Println("hay que hacer merge")
		//codear una funcion para hacer merge y llamarla aqui!!!
	}

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
