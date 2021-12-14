# Laboratorio 3:
## Información grupo
Grupo: 14
Fecha de entrega: Martes 14 Noviembre 2021
## Integrantes:
    - Nicolás Araya 201773550-3
    - Matías Carvajal 201873559-0
    - Fabián Jiménez 201673503-8



# Asignación de máquinas
- Máquina virtual dist53: Broker
- Máquina virtual dist54: Fulcrum1 + Leia
-  Máquina virtual dist55: Fulcrum2 + Almirante
-  Máquina virtual dist56: Fulcrum3 + Ahsoka


# Instrucciones
## Ejecución de programas:
- Se ingresa a las 4 máquinas. Se ingresa a la máquina dist54, dist55, dist56 y dist5. En cada ventana se ejecutan los siguientes comandos siguiendo el orden:
- [alumno@dist54 ~]$ make run_fulcrum1
- [alumno@dist55 ~]$ make run_fulcrum2
- [alumno@dist56 ~]$ make run_fulcrum3
- [alumno@dist53 ~]$ make run_broker
- Se ingresa a la máquina dist54, dist55 y dist56 en 3 ventanas diferentes y se ejecuta:
- [alumno@dist54 ~]$ make run_leia
- [alumno@dist55 ~]$ make run_almirante
- [alumno@dist56 ~]$ make run_ahsoka
### Los siguientes comandos permiten ver las carpetas de registrosPlanetarios, que contienen los registros de los planetas, por ejemplo, Saturno.txt. Cada registro.txt  contiene lineas de la forma nombre_planeta nombre_ciudad cantidad_rebeldes.
- [alumno@dist54 ~]$ make registros
- [alumno@dist55 ~]$ make registros
- [alumno@dist56 ~]$ make registros
### Los siguientes comandos permiten ver la carpeta de logsRegistros, que permiten ver la carpeta de registros, por ejemplo, Saturno.txt.  Cada registro.txt  contiene lineas de las acciones realizadas en cada fulcrum considerando UpdateNumber, AddCity, UpdateName y DeleteCity.
- [alumno@dist54 ~]$ make logs
- [alumno@dist55 ~]$ make logs
- [alumno@dist56 ~]$ make logs

### Los siguientes comandos permiten limpiar los registros de las carpetas de registrosPlanetarios y logsRegistros.
- [alumno@dist54 ~]$ make clean
- [alumno@dist55 ~]$ make clean
- [alumno@dist56 ~]$ make clean


# Decisiones de diseño
1. Cuando se retorna cantidad de rebeldes igual a "-1" significa que nunca se creo el registro para esa ciudad o fue eliminado/cambio de nombre por algún informante.

2. Al hacer un merge, este toma al servidor que posee más cambios y replica la información de ese servidor al resto de servidores (para un
planeta en particular). 




