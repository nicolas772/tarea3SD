# Laboratorio 3:
##Información grupo
Grupo: 14
Fecha de entrega: Martes 14 Noviembre 2021
## Integrantes:
    - Nicolás Araya 201773550-3
    - Matías Carvajal 201873559-0
    - Fabián Jiménez 201673503-8



# Asignación de máquinas
## Máquina virtual dist53: Broker
## Máquina virtual dist54: Fulcrum1 + Leia
## Máquina virtual dist55: Fulcrum2 + Almirante
## Máquina virtual dist56: Fulcrum3 + Ahsoka


# Instrucciones
Se ingresa a las 4 máquinas. Se ingresa a la máquina dist53, dist54, dist55 y dist 56. En cada ventana se ejecuta lo siguiente:
[alumno@dist53 ~]$ make createDir
[alumno@dist54 ~]$ make createDir
[alumno@dist55 ~]$ make createDir
[alumno@dist56 ~]$ make createDir
[alumno@dist53 ~]$ make run_broker
[alumno@dist54 ~]$ make run_fulcrum1
[alumno@dist55 ~]$ make run_fulcrum2
[alumno@dist56 ~]$ make run_fulcrum3
Se ingresa a la máquina dist54, dist55 y dist56 en 3 ventanas diferentes y se ejecuta:
[alumno@dist54 ~]$ make run_leia
[alumno@dist55 ~]$ make run_almirante
[alumno@dist56 ~]$ make run_ahsoka

# Decisiones de diseño
1. Cuando se retorna cantidad de rebeldes igual a "-1" significa que nunca se creo el registro para esa ciudad o fue eliminado/cambio de nombre por algún informante.

2. Al hacer un merge, este toma al servidor que posee más cambios y replica la información de ese servidor al resto de servidores. 




