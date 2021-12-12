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
	address = "localhost:50050"
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
			fmt.Println("entra a no registro creado")
			new_city := &pb1.CiudadSolicitada{Ciudad: city, Planeta: planet, RelojVector: rv, UltimoServidorFulcrum: server}
			l.ciudades_solic_list.Ciudades = append(l.ciudades_solic_list.Ciudades, new_city)
		}
	}
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewStarWarsClient(conn)
	no_quit := true
	var leia_server *LeiaServer = NewLeiaServer()
	fmt.Println("Bienvenida Princesa Leia")
	for no_quit {
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
				r, err := c.CityLeiaBroker(ctx, &pb.NewCity{NombrePlaneta: planeta, NombreCiudad: ciudad, Action: action})
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
				fmt.Println("Lista Ciudades:", leia_server.ciudades_solic_list)

			} else {
				fmt.Println("Comando Invalido")
			}
		}

	}

}
