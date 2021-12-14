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
	"time"

	pb "example.com/go-starwars-grpc/starwars2"
	"google.golang.org/grpc"
)

const (
	port_almirante           = ":50058"
	port_broker              = ":50061"
	port_ahsoka              = ":50055"
	port_merged		         = ":50085"
	path_registro_planetario = "./fulcrum1/registrosPlanetarios/"
	path_log_registro        = "./fulcrum1/logRegistros/"
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
	fmt.Println("server listening at ", lis.Addr())
	return s.Serve(lis)
}

func leerArchivo(path string) []string {
	var input, err = ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}
	lines := strings.Split(string(input), "\n")
	fmt.Println(lines[0])
	return lines
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
	cant_rebeldes := -1
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
	fmt.Println("")
	fmt.Println("Received from", in.GetSender(),":", in.GetNombrePlaneta())
	fmt.Println("Received from", in.GetSender(),":", in.GetNombreCiudad())
	fmt.Println("Received from", in.GetSender(),":", in.GetNuevoValor())
	fmt.Println("Received from", in.GetSender(),":", in.GetAction())
	accion := in.GetAction()
	planeta := in.GetNombrePlaneta()
	ciudad := in.GetNombreCiudad()
	valor := in.GetNuevoValor()
	cant_rebeldes := BuscarCantidadRebeldes(planeta, ciudad) //si retorna -1, quiere decir que no esta creada la ciudad, en cualquier otro caso, si esta creada
	//Manejo de archivo RegistroPlanetario
	ruta2 := path_registro_planetario + planeta + ".txt"
	pudo_realizar_comando := false
	if accion == "AddCity" {
		if cant_rebeldes == -1 {
			linea_archivo := planeta + " " + ciudad + " " + valor
			crearArchivo(ruta2)
			escribeArchivo(linea_archivo, ruta2)
			pudo_realizar_comando = true
		} else {
			fmt.Println("No se puede añadir ciudad, ya esta creada")
		}
	} else {
		if cant_rebeldes >= 0 {
			editarArchivo(ruta2, ciudad, valor, accion)
			pudo_realizar_comando = true
		} else {
			fmt.Println("No se puede editar o eliminar ciudad que no existe en el registro. El otro informante la elimino o jamas ha sido creada.")
		}
	}
	//manejo de vector_reloj y logRegistro
	vector_retorno := &pb.RespFulcrum1{}
	if pudo_realizar_comando {
		ruta1 := path_log_registro + planeta + ".txt"
		crearArchivo(ruta1) //si ya eviste, no lo crea
		escribeArchivo(accion+" "+planeta+" "+ciudad+" "+valor, ruta1)
		no_vector_creado := true
		for _, vect := range s.vectores_list.Vectores {
			if vect.Planeta == planeta {
				vect.RelojVector[0]++ //aumento en 1 la cantidad de cambios
				no_vector_creado = false
				vector_retorno = &pb.RespFulcrum1{RelojVector: vect.RelojVector, Planeta: planeta, SeRealizoMod: "si"}
			}
		}
		if no_vector_creado {
			created_vector := &pb.Vector{RelojVector: []int32{1, 0, 0}, Planeta: planeta}
			s.vectores_list.Vectores = append(s.vectores_list.Vectores, created_vector)
			vector_retorno = &pb.RespFulcrum1{RelojVector: created_vector.RelojVector, Planeta: planeta, SeRealizoMod: "si"}
		}
	} else { //quiere decir que el informante no pudo realizar cambios
		existe_planeta := false
		for _, vect := range s.vectores_list.Vectores {
			if vect.Planeta == planeta {
				vector_retorno = &pb.RespFulcrum1{RelojVector: vect.RelojVector, Planeta: planeta, SeRealizoMod: "no"}
				existe_planeta = true
			}
		}
		if !existe_planeta { //si no existe planeta, retorno un reloj vector nulo
			vector_aux := &pb.RespFulcrum1{RelojVector: []int32{0, 0, 0}, Planeta: planeta, SeRealizoMod: "no"}
			vector_retorno = vector_aux
		}
	}
	return vector_retorno, nil
}

func (s *FulcrumServer) CityBrokerFulcrum(ctx context.Context, in *pb.NewCity1) (*pb.RespFulcrum2, error) {
	fmt.Println("Received from Broker:", in.GetNombrePlaneta())
	fmt.Println("Received from Broker:", in.GetNombreCiudad())
	fmt.Println("Received from Broker:", in.GetAction())
	planeta := in.GetNombrePlaneta()
	ciudad := in.GetNombreCiudad()
	cant_rebeldes := BuscarCantidadRebeldes(planeta, ciudad)
	reloj_vector := BuscarRelojVector(planeta, s)
	return &pb.RespFulcrum2{CantRebeldes: int32(cant_rebeldes), RelojVector: reloj_vector, ServidorContactado: "Fulcrum1"}, nil
}

func (s *FulcrumServer) RelojesBrokerFulcrum(ctx context.Context, in *pb.Planeta) (*pb.RespFulcrum1, error) {
	fmt.Println("Received from Broker:", in.GetNombrePlaneta())
	planeta_consultado := in.GetNombrePlaneta()
	reloj_vector := BuscarRelojVector(planeta_consultado, s)
	return &pb.RespFulcrum1{RelojVector: reloj_vector, Planeta: planeta_consultado}, nil
}

func PreguntarFul2() *pb.RelojesYRegistros {
	direccion := "10.6.40.195:50064"

	conn, err := grpc.Dial(direccion, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewStarWars1Client(conn)
	ctx1, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()
	r, err := c.PreguntarRelojesYRegistros(ctx1, &pb.SolMerge{HacerMerge: true})
	if err != nil {
		log.Fatalln(err)
	}
	return r
}

func mandarFul2(f2 *pb.RelojesYRegistros) bool{
	direccion := "10.6.40.195:50064"

	conn, err := grpc.Dial(direccion, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewStarWars1Client(conn)
	ctx1, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()
	r, err := c.MandarFulcrums(ctx1, f2)
	if err != nil {
		log.Fatalln(err)
	}
	return r.GetHacerMerge()
}

func PreguntarFul3() *pb.RelojesYRegistros {
	direccion := "10.6.40.196:50065"

	conn, err := grpc.Dial(direccion, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewStarWars1Client(conn)
	ctx1, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()
	r, err := c.PreguntarRelojesYRegistros(ctx1, &pb.SolMerge{HacerMerge: true})
	if err != nil {
		log.Fatalln(err)
	}
	return r
}

func mandarFul3(f3 *pb.RelojesYRegistros) bool{
	direccion := "10.6.40.196:50065"

	conn, err := grpc.Dial(direccion, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewStarWars1Client(conn)
	ctx1, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()
	r, err := c.MandarFulcrums(ctx1, f3)
	if err != nil {
		log.Fatalln(err)
	}
	return r.GetHacerMerge()
}

func planetasEnServidores(f1 *pb.RelojesYRegistros, f2 *pb.RelojesYRegistros, f3 *pb.RelojesYRegistros) []string {
	lista_planetas := []string{}
	estaPlaneta := false
	//Fulcrum1
	for _, planeta := range f1.ListaVectores.Vectores {
		for _, planeta2 := range lista_planetas {
			if planeta2 == planeta.Planeta {
				estaPlaneta = true
			}
		}
		if !estaPlaneta {
			lista_planetas = append(lista_planetas, planeta.Planeta)
		}
	}
	estaPlaneta = false
	//Fulcrum2
	for _, planeta := range f2.ListaVectores.Vectores {
		for _, planeta2 := range lista_planetas {
			if planeta2 == planeta.Planeta {
				estaPlaneta = true
			}
		}
		if !estaPlaneta {
			lista_planetas = append(lista_planetas, planeta.Planeta)
		}
	}
	estaPlaneta = false
	//Fulcrum3
	for _, planeta := range f3.ListaVectores.Vectores {
		for _, planeta2 := range lista_planetas {
			if planeta2 == planeta.Planeta {
				estaPlaneta = true
			}
		}
		if !estaPlaneta {
			lista_planetas = append(lista_planetas, planeta.Planeta)
		}
	}
	return lista_planetas
}

func agregarPlanetasYConsistencia(arregloPlanetas []string, f1 *pb.RelojesYRegistros, f2 *pb.RelojesYRegistros, f3 *pb.RelojesYRegistros) {
	relojPlaneta := []int32{0,0,0}
	posicionF1 := 0
	posicionF2 := 0
	posicionF3 := 0
	iterador := 0

	for _, planeta := range arregloPlanetas {
		//Fulcrum1
		iterador = 0
		inPlaneta := false
		posMayoresCambios := int32(0)
		cantidadCambios := int32(0)

		for _, planeta2 := range f1.ListaVectores.Vectores {
			if planeta == planeta2.Planeta {
				relojPlaneta[0]	= planeta2.RelojVector[0]
				inPlaneta = true
				if cantidadCambios <= planeta2.RelojVector[0] {
					posMayoresCambios = 0
					cantidadCambios = planeta2.RelojVector[0]
				} 
				posicionF1 = iterador
			}
			iterador++
		}
		if !inPlaneta {//Planeta no se encuentra en el servidor, hay que agregarlo al final
			f1.ListaVectores.Vectores = append(f1.ListaVectores.Vectores, &pb.Vector{Planeta: planeta, RelojVector: relojPlaneta})
			f1.LogRegistros.Registr = append(f1.LogRegistros.Registr, &pb.RegistroUnitario{Array: []string{}})
			posicionF1 = len(f1.ListaVectores.Vectores) - 1
			
		}
		iterador = 0
		inPlaneta = false
		//Fulcrum2
		for _, planeta2 := range f2.ListaVectores.Vectores {
			if planeta == planeta2.Planeta {
				relojPlaneta[1]	= planeta2.RelojVector[1]
				inPlaneta = true
				if cantidadCambios <= planeta2.RelojVector[1] {
					posMayoresCambios = 1
					cantidadCambios = planeta2.RelojVector[1]
				} 
				posicionF2 = iterador
			}
			iterador++
		}
		if !inPlaneta {//Planeta no se encuentra en el servidor, hay que agregarlo al final
			f2.ListaVectores.Vectores = append(f2.ListaVectores.Vectores, &pb.Vector{Planeta: planeta, RelojVector: relojPlaneta})
			f2.LogRegistros.Registr = append(f2.LogRegistros.Registr, &pb.RegistroUnitario{Array: []string{}})
			posicionF2 = len(f2.ListaVectores.Vectores) - 1
		}
		iterador = 0
		inPlaneta = false
		//Fulcrum3
		for _, planeta2 := range f3.ListaVectores.Vectores {
			if planeta == planeta2.Planeta {
				relojPlaneta[2]	= planeta2.RelojVector[2]
				inPlaneta = true
				if cantidadCambios <= planeta2.RelojVector[2] {
					posMayoresCambios = 2
					cantidadCambios = planeta2.RelojVector[2]
				} 
				posicionF3 = iterador
			}
			iterador++
		}
		if !inPlaneta {//Planeta no se encuentra en el servidor, hay que agregarlo al final
			f3.ListaVectores.Vectores = append(f3.ListaVectores.Vectores, &pb.Vector{Planeta: planeta, RelojVector: relojPlaneta})
			f3.LogRegistros.Registr = append(f3.LogRegistros.Registr, &pb.RegistroUnitario{Array: []string{}})
			posicionF3 = len(f3.ListaVectores.Vectores) - 1	
		}

		//Se actualizan los relojes de los demas falcrums
		f1.ListaVectores.Vectores[posicionF1].RelojVector = relojPlaneta	
		f2.ListaVectores.Vectores[posicionF2].RelojVector = relojPlaneta	
		f3.ListaVectores.Vectores[posicionF3].RelojVector = relojPlaneta
		
		//Todos misma información
		if posMayoresCambios == 0 {//fulcrum1 posee más cambios
			f2.LogRegistros.Registr[posicionF2].Array = f1.LogRegistros.Registr[posicionF1].Array
			f3.LogRegistros.Registr[posicionF3].Array = f1.LogRegistros.Registr[posicionF1].Array
		}else if posMayoresCambios == 1 {//fulcrum2 posee más cambios
			f1.LogRegistros.Registr[posicionF1].Array = f2.LogRegistros.Registr[posicionF2].Array
			f3.LogRegistros.Registr[posicionF3].Array = f2.LogRegistros.Registr[posicionF2].Array
		}else{//fulcrum3 posee más cambios
			f1.LogRegistros.Registr[posicionF1].Array = f3.LogRegistros.Registr[posicionF3].Array
			f2.LogRegistros.Registr[posicionF2].Array = f3.LogRegistros.Registr[posicionF3].Array
		}
		
		relojPlaneta = []int32{0,0,0}
	}

}
func (s *FulcrumServer) ConsistenciaEventual(ctx context.Context, in *pb.SolMerge) (*pb.RespBroker3, error) {
	posicion := 0
	seHizo := false

	//Reglas: el servictor que tenga más cambios en un planeta tiene prioridad
	fulcrum1 := &pb.RelojesYRegistros{ListaVectores: &pb.VectoresList{}, LogRegistros: &pb.RegistroList{}}

	for _, vect := range s.vectores_list.Vectores {
		fulcrum1.ListaVectores.Vectores = append(fulcrum1.ListaVectores.Vectores, vect)//agregamos el vector a los vectores
		registro := leerArchivo(path_registro_planetario + vect.Planeta + ".txt")
		fulcrum1.LogRegistros.Registr = append(fulcrum1.LogRegistros.Registr, &pb.RegistroUnitario{Array: registro})//agregamos el registro al vector de registros
	}

	//Buscar Información de los otros servidores Fulcrum
	fulcrum2 := PreguntarFul2()
	fulcrum3 := PreguntarFul3()
	//Comenzar consistencia eventual
	//Misma cantidad de planetas en los servidores (se agregan servidores)
	arrayPlanetas := planetasEnServidores(fulcrum1, fulcrum2, fulcrum3)
	//misma informacion en los servidores
	agregarPlanetasYConsistencia(arrayPlanetas, fulcrum1, fulcrum2, fulcrum3)

	//Se actualiza fulcrum 1
	for _, vect := range fulcrum1.LogRegistros.Registr {
		ruta1 := path_registro_planetario + fulcrum1.ListaVectores.Vectores[posicion].Planeta + ".txt"
		os.Remove(ruta1)
		crearArchivo(ruta1) 
		ruta2 := path_log_registro + fulcrum1.ListaVectores.Vectores[posicion].Planeta + ".txt"
		os.Remove(ruta2)
		crearArchivo(ruta2) 
		for _, linea := range vect.Array {
			if (linea != ""){
				escribeArchivo(linea, ruta1)
			}
		}
		
		posicion++
	}

	mandarFul2(fulcrum2)
	mandarFul3(fulcrum3)
	if mandarFul2(fulcrum2) && mandarFul3(fulcrum3) {
		seHizo = true
	}
	return &pb.RespBroker3{SeHizoMerge: seHizo}, nil
}

func main() {
	var fulcrum_server *FulcrumServer = NewFulcrumServer()
	//solo para el Merge
	go func() {
		if err := fulcrum_server.Run(port_merged); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
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
