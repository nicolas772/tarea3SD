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

func NewLeiaServer() *LeiaServer {
	return &LeiaServer{
		ciudades_solic_list: &pb1.CiudadesSolicitadasList{},
	}
}

type LeiaServer struct {
	ciudades_solic_list *pb1.CiudadesSolicitadasList
}

func ActualizarListaCiudades(city string, planet string, server string, l *LeiaServer, rv []int32) {
	no_registro_creado := true
	if (rv[0] != 0) || (rv[1] != 0) || (rv[2] != 0) {
		for _, vect := range l.ciudades_solic_list.Ciudades {
			if vect.Ciudad == city {
				vect.RelojVector = rv
				vect.UltimoServidorFulcrum = server
				vect.Planeta = planet
				no_registro_creado = false
			}
		}
		if no_registro_creado {
			new_city := &pb1.CiudadSolicitada{Ciudad: city, Planeta: planet, RelojVector: rv, UltimoServidorFulcrum: server}
			l.ciudades_solic_list.Ciudades = append(l.ciudades_solic_list.Ciudades, new_city)
		}
	}
}

func BuscarRelojVectorYServidor(planeta string, s *LeiaServer) ([]int32, string) {
	vector_retorno := []int32{0, 0, 0}
	ultimo_servidor := ""
	for _, vect := range s.ciudades_solic_list.Ciudades {
		if vect.Planeta == planeta {
			vector_retorno = vect.RelojVector
			ultimo_servidor = vect.UltimoServidorFulcrum
		}
	}
	return vector_retorno, ultimo_servidor
}

func main() {
	fmt.Println("---AQUI- Inicio----")
	address := "10.6.40.193:50050"
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	fmt.Println("---AQUI--1-")
	c := pb.NewStarWarsClient(conn)
	fmt.Println("---AQUI--2--")
	no_quit := true
	var leia_server *LeiaServer = NewLeiaServer()
	fmt.Println("---AQUI--3-")
	fmt.Println("Bienvenida Princesa Leia")
	for no_quit {
		fmt.Println("")
		fmt.Println("Por favor, ingrese el comando. Para salir, presione 'Q'")
		reader := bufio.NewReader(os.Stdin)
		entrada, _ := reader.ReadString('\n')         // Leer hasta el separador de salto de línea
		comando := strings.TrimRight(entrada, "\r\n") // Remover el salto de línea de la entrada del usuario
		if comando == "q" || comando == "Q" {
			fmt.Println("Hasta pronto princesa Leia")
			no_quit = false
		} else {
			split := strings.Split(comando, " ")

			if split[0] == "GetNumberRebelds" && len(split) == 3 { //comando "GetNumberRebelds"
				planeta := split[1]
				ciudad := split[2]
				action := split[0]
				ctx, cancel := context.WithTimeout(context.Background(), time.Second)
				defer cancel()
				ultimo_vector, ultimo_servidor := BuscarRelojVectorYServidor(planeta, leia_server)
				r, err := c.CityLeiaBroker(ctx, &pb.NewCity{NombrePlaneta: planeta, NombreCiudad: ciudad, Action: action, Sender: "leia", RelojVector: ultimo_vector, UltimoServidor: &ultimo_servidor})
				if err != nil {
					log.Fatalf("could not create city in broker: %v", err)
				}
				cantidad_rebeldes := r.GetCantRebeldes()
				reloj := r.GetRelojVector()
				servidor_consultado := r.GetServidorContactado()
				fmt.Println("Cantidad Rebeldes:", cantidad_rebeldes)
				fmt.Println("reloj vector:", reloj)
				fmt.Println("servidor contactado:", servidor_consultado)
				ActualizarListaCiudades(ciudad, planeta, servidor_consultado, leia_server, reloj)
			} else {
				fmt.Println("Comando Invalido")
			}
		}

	}

}
