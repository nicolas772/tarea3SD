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
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewStarWarsClient(conn)

	flag := true
	fmt.Println("Bienvenido Almirante Thrawn")
	for flag {
		fmt.Println("Por favor, ingrese el comando. Para salir, presione 'Q'")
		reader := bufio.NewReader(os.Stdin)
		entrada, _ := reader.ReadString('\n')         // Leer hasta el separador de salto de línea
		comando := strings.TrimRight(entrada, "\r\n") // Remover el salto de línea de la entrada del usuario

		if comando == "q" || comando == "Q" {
			fmt.Println("Hasta pronto Almirante Thrawn")
			flag = false
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

				ctx, cancel := context.WithTimeout(context.Background(), time.Second)
				defer cancel()
				r, err := c.AddCity(ctx, &pb.NewCity{NombrePlaneta: planeta, NombreCiudad: ciudad, Action: action, NuevoValor: &cant_soldados})
				if err != nil {
					log.Fatalf("could not create city: %v", err)
				}
				fmt.Println("Servidor Fulcrum asignado: ", r.GetDireccionServidor())
			} else if (split[0] == "UpdateName" || split[0] == "UpdateNumber") && len(split) == 4 { //Comando "UpdateName" y "UpdateNumber"
				planeta := split[1]
				ciudad := split[2]
				action := split[0]
				nuevo_valor := split[3]
				ctx, cancel := context.WithTimeout(context.Background(), time.Second)
				defer cancel()
				r, err := c.AddCity(ctx, &pb.NewCity{NombrePlaneta: planeta, NombreCiudad: ciudad, Action: action, NuevoValor: &nuevo_valor})
				if err != nil {
					log.Fatalf("could not create city: %v", err)
				}
				fmt.Println("Servidor Fulcrum asignado: ", r.GetDireccionServidor())

			} else if split[0] == "DeleteCity" && len(split) == 3 { //Comando "DeleteCity"
				planeta := split[1]
				ciudad := split[2]
				action := split[0]
				ctx, cancel := context.WithTimeout(context.Background(), time.Second)
				defer cancel()
				r, err := c.AddCity(ctx, &pb.NewCity{NombrePlaneta: planeta, NombreCiudad: ciudad, Action: action})
				if err != nil {
					log.Fatalf("could not create city: %v", err)
				}
				fmt.Println("Servidor Fulcrum asignado: ", r.GetDireccionServidor())

			} else {
				fmt.Println("Comando Invalido")
			}

		}

	}

}
