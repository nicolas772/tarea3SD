package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	pb "example.com/go-starwars-grpc/starwars2"
	"google.golang.org/grpc"
)

const (
	port_almirante           = ":50060"
	port_broker              = ":50063"
	port_ahsoka              = ":50057"
	path_registro_planetario = "./fulcrum3/registrosPlanetarios/"
	path_log_registro        = "./fulcrum3/logRegistros/"
)

func NewFulcrumServer() *FulcrumServer {
	return &FulcrumServer{
		vectores_list: &pb.VectoresList{},
	}
}

type FulcrumServer struct {
	pb.UnimplementedStarWars1Server
	vectores_list *pb.VectoresList
}

func (server *FulcrumServer) Run(port string) error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterStarWars1Server(s, server)
	log.Printf("server listening at %v", lis.Addr())
	return s.Serve(lis)
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

func BuscarCantidadRebeldes(planeta string, ciudad string) int {
	cant_rebeldes := 0
	path := path_registro_planetario + planeta + ".txt"
	input, err := ioutil.ReadFile(path)
	if err != nil {
		//log.Fatalln(err)
		return cant_rebeldes
	}
	lines := strings.Split(string(input), "\n")
	for _, line := range lines {
		if line != "" {
			split_line := strings.Split(line, " ")
			if split_line[1] == ciudad {
				cant_rebeldes, err = strconv.Atoi(split_line[2])
				if err != nil {
					log.Fatalln(err)
				}
			}
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(path, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
	return cant_rebeldes
}

func BuscarRelojVector(planeta string, s *FulcrumServer) []int32 {
	vector_retorno := []int32{0, 0, 0}
	for _, vect := range s.vectores_list.Vectores {
		if vect.Planeta == planeta {
			vector_retorno = vect.RelojVector
		}
	}
	return vector_retorno
}

func (s *FulcrumServer) CityMgmtFulcrum(ctx context.Context, in *pb.NewCity1) (*pb.RespFulcrum1, error) {
	log.Printf("Received from %v: %v", in.GetSender(), in.GetNombrePlaneta())
	log.Printf("Received from %v: %v", in.GetSender(), in.GetNombreCiudad())
	log.Printf("Received from %v: %v", in.GetSender(), in.GetNuevoValor())
	log.Printf("Received from %v: %v", in.GetSender(), in.GetAction())
	accion := in.GetAction()
	planeta := in.GetNombrePlaneta()
	ciudad := in.GetNombreCiudad()
	valor := in.GetNuevoValor()

	//manejo de vector_reloj
	no_vector_creado := true
	vector_retorno := &pb.RespFulcrum1{}
	for _, vect := range s.vectores_list.Vectores {
		if vect.Planeta == planeta {
			vect.RelojVector[0]++ //aumento en 1 la cantidad de cambios
			no_vector_creado = false
			vector_retorno = vect
		}
	}
	if no_vector_creado {
		created_vector := &pb.RespFulcrum1{RelojVector: []int32{1, 0, 0}, Planeta: planeta}
		s.vectores_list.Vectores = append(s.vectores_list.Vectores, created_vector)
		vector_retorno = created_vector
	}

	//Manejo de archivos logRegistro y RegistroPlanetario
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
	fmt.Println("vectores list:", s.vectores_list)
	return vector_retorno, nil
}

func (s *FulcrumServer) CityBrokerFulcrum(ctx context.Context, in *pb.NewCity1) (*pb.RespFulcrum2, error) {
	log.Printf("Received from Broker: %v", in.GetNombrePlaneta())
	log.Printf("Received from Broker: %v", in.GetNombreCiudad())
	log.Printf("Received from Broker: %v", in.GetAction())
	planeta := in.GetNombrePlaneta()
	ciudad := in.GetNombreCiudad()
	cant_rebeldes := BuscarCantidadRebeldes(planeta, ciudad)
	reloj_vector := BuscarRelojVector(planeta, s)
	return &pb.RespFulcrum2{CantRebeldes: int32(cant_rebeldes), RelojVector: reloj_vector, ServidorContactado: "Fulcrum1"}, nil
}

func main() {
	var fulcrum_server *FulcrumServer = NewFulcrumServer()
	//conexion con Almirante
	go func() {
		if err := fulcrum_server.Run(port_almirante); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	//Conexion con broker
	go func() {
		if err := fulcrum_server.Run(port_broker); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	//conexion con ahsoka
	go func() {
		if err := fulcrum_server.Run(port_ahsoka); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	fmt.Println("Para salir del servidor, presione Q y luego Enter")
	var sa string
	fmt.Scan(&sa)
	fmt.Println("Servidor Fulcrum 1 finalizado")

}
