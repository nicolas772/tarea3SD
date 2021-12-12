clean:
	rm -r fulcrum1/logRegistros
	rm -r fulcrum1/registrosPlanetarios
	rm -r fulcrum2/logRegistros
	rm -r fulcrum2/registrosPlanetarios
	rm -r fulcrum3/logRegistros
	rm -r fulcrum3/registrosPlanetarios
	mkdir fulcrum1/logRegistros
	mkdir fulcrum1/registrosPlanetarios
	mkdir fulcrum2/logRegistros
	mkdir fulcrum2/registrosPlanetarios
	mkdir fulcrum3/logRegistros
	mkdir fulcrum3/registrosPlanetarios

createDir:
	mkdir fulcrum1/logRegistros
	mkdir fulcrum1/registrosPlanetarios
	mkdir fulcrum2/logRegistros
	mkdir fulcrum2/registrosPlanetarios
	mkdir fulcrum3/logRegistros
	mkdir fulcrum3/registrosPlanetarios

run_broker:
	go run broker/broker.go

run_fulcrum1:
	go run fulcrum1/fulcrum1.go

run_fulcrum2:
	go run fulcrum2/fulcrum2.go

run_fulcrum3:
	go run fulcrum3/fulcrum3.go

run_almirante:
	go run almirante/almirante.go

run_ahsoka:
	go run ahsoka/ahsoka.go

run_leia:
	go run leia/leia.go

