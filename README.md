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
1. Cuando se le retorna a leia una cantidad de rebeldes igual a "-1", significa que nunca se creó el registro para esa ciudad o fué eliminado/cambió de nombre por algún informante.

2. En cuanto al merge, hay dos formas en las cuales se producen:

- Evento cada dos minutos, en la que el broker envia un mensaje al fulcrum maestro (en nuestro caso, fulcrum 1), y este gestiona todo el merge.
- Cuando un informante o leia envia una instrucción, el broker se encarga de comparar los relojes de todos los servidores para el planeta en particular. Si el broker se da cuenta que hay inconsistencias (o cambios concurrentes), este envia un mensaje disparando el "merge de emergencia". Luego de realizar el merge de emergencia, se resuelve automaticamente la petición original del informante o de leia.

El criterio utilizado para resolver problemas entre planetas fue el siguiente:

- Si en dos servidores hay contenidos distintos para un mismo planeta, el servidor fulcrum maestro resuelve quedarse con el archivo del servidor fulcrum que más cambios realizó al registro planetario.






