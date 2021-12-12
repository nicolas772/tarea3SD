package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	pb "example.com/go-starwars-grpc/starwars"
	pb1 "example.com/go-starwars-grpc/starwars2"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50054"
)

func NewAhsokaServer() *AhsokaServer {
	return &AhsokaServer{
		registros_modificados_list: &pb1.RegistrosModificadosList{},
	}
}

type AhsokaServer struct {
	registros_modificados_list *pb1.RegistrosModificadosList
}

func ActualizarListaRegistro(planet string, server string, a *AhsokaServer, rv []int32) {
	no_registro_creado := true
	for _, vect := range a.registros_modificados_list.Registros {
		if vect.Planeta == planet {
			vect.RelojVector = rv
			vect.UltimoServidorFulcrum = server
			no_registro_creado = false
		}
	}
	if no_registro_creado {
		new_registro := &pb1.RegistroModificado{Planeta: planet, RelojVector: rv, UltimoServidorFulcrum: server}
		a.registros_modificados_list.Registros = append(a.registros_modificados_list.Registros, new_registro)
	}
}

func BuscarRelojVectorYServidor(planeta string, s *AhsokaServer) ([]int32, string) {
	vector_retorno := []int32{0, 0, 0}
	ultimo_servidor := ""
	for _, vect := range s.registros_modificados_list.Registros {
		if vect.Planeta == planeta {
			vector_retorno = vect.RelojVector
			ultimo_servidor = vect.UltimoServidorFulcrum
		}
	}
	return vector_retorno, ultimo_servidor
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewStarWarsClient(conn)
	no_quit := true
	var ahsoka_server *AhsokaServer = NewAhsokaServer()
	fmt.Println("Bienvenida Ahsoka Tano")
	for no_quit {
		fmt.Println("")
		fmt.Println("Por favor, ingrese el comando. Para salir, presione 'Q'")
		reader := bufio.NewReader(os.Stdin)
		entrada, _ := reader.ReadString('\n')         // Leer hasta el separador de salto de línea
		comando := strings.TrimRight(entrada, "\r\n") // Remover el salto de línea de la entrada del usuario

		if comando == "q" || comando == "Q" {
			fmt.Println("Hasta pronto Almirante Thrawn")
			no_quit = false
		} else {
			split := strings.Split(comando, " ")

			if split[0] == "AddCity" && (len(split) == 3 || len(split) == 4) { //comando "AddCity"
				planeta := split[1]
				ciudad := split[2]
				action := split[0]
				cant_soldados := "0"
				if len(split) == 4 {
					cant_soldados = split[3]
				}

				ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
				defer cancel()
				vector_reloj, ultimo_servidor := BuscarRelojVectorYServidor(planeta, ahsoka_server)
				r, err := c.CityMgmtBroker(ctx, &pb.NewCity{NombrePlaneta: planeta, NombreCiudad: ciudad, Action: action, NuevoValor: &cant_soldados, Sender: "ahsoka", RelojVector: vector_reloj, UltimoServidor: &ultimo_servidor})
				if err != nil {
					log.Fatalf("could not create city in broker: %v", err)
				}
				servidor_asignado := r.GetDireccionServidor()
				fmt.Println("Servidor Fulcrum asignado:", servidor_asignado)

				conn1, err1 := grpc.Dial(servidor_asignado, grpc.WithInsecure(), grpc.WithBlock())
				if err1 != nil {
					log.Fatalf("did not connect: %v", err)
				}
				//defer conn1.Close()
				c1 := pb1.NewStarWars1Client(conn1)
				ctx1, cancel := context.WithTimeout(context.Background(), time.Second)
				defer cancel()
				r1, errr := c1.CityMgmtFulcrum(ctx1, &pb1.NewCity1{NombrePlaneta: planeta, NombreCiudad: ciudad, Action: action, NuevoValor: &cant_soldados, Sender: "ahsoka"})
				if errr != nil {
					log.Fatalf("could not create city in fulcrum: %v", errr)
				}
				reloj_vector := r1.GetRelojVector()
				se_realizo_mod := r1.GetSeRealizoMod()
				fmt.Println("Reloj vector: ", reloj_vector)
				if se_realizo_mod == "si" {
					ActualizarListaRegistro(planeta, servidor_asignado, ahsoka_server, reloj_vector)
				} else {
					fmt.Println("No se ha podido realizar la modificacion")
				}
				fmt.Println("lista de registros:", ahsoka_server.registros_modificados_list)
				conn1.Close()

			} else if (split[0] == "UpdateName" || split[0] == "UpdateNumber") && len(split) == 4 { //Comando "UpdateName" y "UpdateNumber"
				planeta := split[1]
				ciudad := split[2]
				action := split[0]
				cant_soldados := split[3]
				ctx, cancel := context.WithTimeout(context.Background(), time.Second)
				defer cancel()
				vector_reloj, ultimo_servidor := BuscarRelojVectorYServidor(planeta, ahsoka_server)
				r, err := c.CityMgmtBroker(ctx, &pb.NewCity{NombrePlaneta: planeta, NombreCiudad: ciudad, Action: action, NuevoValor: &cant_soldados, Sender: "ahsoka", RelojVector: vector_reloj, UltimoServidor: &ultimo_servidor})
				if err != nil {
					log.Fatalf("could not create city: %v", err)
				}
				servidor_asignado := r.GetDireccionServidor()
				fmt.Println("Servidor Fulcrum asignado:", servidor_asignado)
				conn1, err1 := grpc.Dial(servidor_asignado, grpc.WithInsecure(), grpc.WithBlock())
				if err1 != nil {
					log.Fatalf("did not connect: %v", err)
				}
				//defer conn1.Close()
				c1 := pb1.NewStarWars1Client(conn1)
				ctx1, cancel := context.WithTimeout(context.Background(), time.Second)
				defer cancel()
				r1, errr := c1.CityMgmtFulcrum(ctx1, &pb1.NewCity1{NombrePlaneta: planeta, NombreCiudad: ciudad, Action: action, NuevoValor: &cant_soldados, Sender: "ahsoka"})
				if errr != nil {
					log.Fatalf("could not create city in fulcrum: %v", errr)
				}
				reloj_vector := r1.GetRelojVector()
				se_realizo_mod := r1.GetSeRealizoMod()
				fmt.Println("Reloj vector: ", reloj_vector)
				if se_realizo_mod == "si" {
					ActualizarListaRegistro(planeta, servidor_asignado, ahsoka_server, reloj_vector)
				} else {
					fmt.Println("No se ha podido realizar la modificacion")
				}
				fmt.Println("lista de registros:", ahsoka_server.registros_modificados_list)
				conn1.Close()

			} else if split[0] == "DeleteCity" && len(split) == 3 { //Comando "DeleteCity"
				planeta := split[1]
				ciudad := split[2]
				action := split[0]
				ctx, cancel := context.WithTimeout(context.Background(), time.Second)
				defer cancel()
				vector_reloj, ultimo_servidor := BuscarRelojVectorYServidor(planeta, ahsoka_server)
				r, err := c.CityMgmtBroker(ctx, &pb.NewCity{NombrePlaneta: planeta, NombreCiudad: ciudad, Action: action, Sender: "ahsoka", RelojVector: vector_reloj, UltimoServidor: &ultimo_servidor})
				if err != nil {
					log.Fatalf("could not create city: %v", err)
				}
				servidor_asignado := r.GetDireccionServidor()
				fmt.Println("Servidor Fulcrum asignado:", servidor_asignado)
				conn1, err1 := grpc.Dial(servidor_asignado, grpc.WithInsecure(), grpc.WithBlock())
				if err1 != nil {
					log.Fatalf("did not connect: %v", err)
				}
				//defer conn1.Close()
				c1 := pb1.NewStarWars1Client(conn1)
				ctx1, cancel := context.WithTimeout(context.Background(), time.Second)
				defer cancel()
				r1, errr := c1.CityMgmtFulcrum(ctx1, &pb1.NewCity1{NombrePlaneta: planeta, NombreCiudad: ciudad, Action: action, Sender: "ahsoka"})
				if errr != nil {
					log.Fatalf("could not create city in fulcrum: %v", errr)
				}
				reloj_vector := r1.GetRelojVector()
				se_realizo_mod := r1.GetSeRealizoMod()
				fmt.Println("Reloj vector: ", reloj_vector)
				if se_realizo_mod == "si" {
					ActualizarListaRegistro(planeta, servidor_asignado, ahsoka_server, reloj_vector)
				} else {
					fmt.Println("No se ha podido realizar la modificacion")
				}
				fmt.Println("lista de registros:", ahsoka_server.registros_modificados_list)
				conn1.Close()

			} else {
				fmt.Println("Comando Invalido")
			}

		}

	}

}
