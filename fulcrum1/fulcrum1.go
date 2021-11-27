package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"

	pb "example.com/go-starwars-grpc/starwars2"
	"google.golang.org/grpc"
)

const (
	port                     = ":50052"
	path_registro_planetario = "./fulcrum1/registrosPlanetarios/"
	path_log_registro        = "./fulcrum1/logRegistros/"
)

type FulcrumServer struct {
	pb.UnimplementedStarWars1Server
}

func crearArchivo(path string) {
	//Verifica que el archivo existe
	var _, err = os.Stat(path)
	//Crea el archivo si no existe
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if existeError(err) {
			return
		}
		defer file.Close()
	}
	//fmt.Println("File Created Successfully", path)
}
func escribeArchivo(texto string, path string) {
	// Abre archivo usando permisos READ & WRITE
	var file, err = os.OpenFile(path, os.O_RDWR, 0666)
	if existeError(err) {
		return
	}
	defer file.Close()
	// Escribe algo de texto linea por linea
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	_, err = file.WriteString(string(dat) + texto + "\n")

	if existeError(err) {
		return
	}
	// Salva los cambios
	err = file.Sync()
	if existeError(err) {
		return
	}
	//fmt.Println("Archivo actualizado existosamente.")
}

func existeError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func editarArchivo(path string, ciudad string, valor_nuevo string, modo string) {
	input, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}
	lines := strings.Split(string(input), "\n")
	for i, line := range lines {
		if line != "" {
			split_line := strings.Split(line, " ")
			if split_line[1] == ciudad {
				if modo == "UpdateName" {
					split_line[1] = valor_nuevo
					linea := strings.Join(split_line, " ")
					lines[i] = linea
				} else if modo == "UpdateNumber" {
					split_line[2] = valor_nuevo
					linea := strings.Join(split_line, " ")
					lines[i] = linea
				} else if modo == "DeleteCity" {
					lines = append(lines[:i], lines[i+1:]...)
				}
			}
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(path, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}

func (s *FulcrumServer) CityMgmtFulcrum(ctx context.Context, in *pb.NewCity1) (*pb.RespFulcrum1, error) {
	log.Printf("Received: %v", in.GetNombrePlaneta())
	log.Printf("Received: %v", in.GetNombreCiudad())
	log.Printf("Received: %v", in.GetNuevoValor())
	log.Printf("Received: %v", in.GetAction())
	accion := in.GetAction()
	planeta := in.GetNombrePlaneta()
	ciudad := in.GetNombreCiudad()
	valor := in.GetNuevoValor()
	reloj_vector := [3]int{1, 2, 3}
	ruta1 := path_log_registro + planeta + ".txt"
	crearArchivo(ruta1)
	escribeArchivo(accion+" "+planeta+" "+ciudad+" "+valor, ruta1)
	ruta2 := path_registro_planetario + planeta + ".txt"
	if accion == "AddCity" {
		linea_archivo := planeta + " " + ciudad + " " + valor
		crearArchivo(ruta2)
		escribeArchivo(linea_archivo, ruta2)
	} else {
		editarArchivo(ruta2, ciudad, valor, accion)
	}

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
