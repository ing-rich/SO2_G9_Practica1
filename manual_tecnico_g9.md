## Universidad de San Carlos de Guatemala
## Facultad de ingeniería
## Escuela de vacaciones diciembre 2020
___
#### Manual técnico
#### Practica 1
#### Grupo 9

### Librerías a instalar
Para poder ejecutar la aplicación es necesario instalar algunas librerías previamente para las cuales se daran los comandos en linux.
- gcc
- Make
- golang
- git

La distribución con la que fue desarrollado el sistemas es Ubuntu 20.04 LTS

#### Instalar GCC
gcc es el compilador para ansi C, C++ y que a su vez instalara make el cual nos permite compilar los archivos del kernel

Desde una terminal en linux actualizaremos primero la lista de paquetes
```sh
$ sudo apt update
```
Ahora instalaremos el paquete build-essential
```sh
$ sudo apt install build-essential
```
El comando anterior instalara gcc,g++ y make, por ultimo instalaremos el manual con la documentación para Gnu/linux de desarrollo

```sh
$ sudo apt-get install manpages-dev
```
De esta forma ya tendremos instalado el compilador para los archivos del modulo y el kernel, lo siguiente sera instalar go, para poder hacer uso de la api y sus servicios.

#### Instalar GO
Lo primero sera descargar el tarball de go, para este caso lo haremos desde consola con wget, tener en cuenta que para este proyecto se utilizo la version go1.15.6
```sh
wget https://dl.google.com/go/go1.15.6.linux-amd64.tar.gz
```
Una vez descargado el archivo, lo siguiente sera descomprimirlo desde su ubicación
```sh
sudo tar -xvf go1.15.6.linux-amd64.tar.gz -C /usr/local/
```
Lo siguiente consistira en establecer la ubicación de go a través de la variable de entorno
```sh
export PATH=$PATH:/usr/local/go/bin
```
Despues de realizar lo anterior solo quedara guardar y cargar la  nueva variable en la terminal actual, para ello usamos el siguiente comando
```sh
source ~/.profile
```

Despues de haber instalado todo lo anterior ya podremos compilar los archivos del proyecto y ejecutarlo para su uso
### Compilar Modulos de Kernel
Es necesario compilar los modulos de kernel, ya que estos crearan los archivos con los datos del cpu y la memoria, para esto debemos ubicarnos en la carpeta que contenga el archivo makefile y la clase .c
```sh
$ cd /m_grupo9
```
Una vez dentro de la carpeta se aconseja primero utilizar el comando sudo su y luego make
```sh
$ sudo su
$ make
```
Una vez se ha compilado el kernel, debemos trasladar el archivo .ko generado a la carpeta /proc, para esto solo bastara usar el siguiente comando
```sh
$ insmod m_grupo9.ko
```

### Compilar y ejecutar api en go
Para compilar y correr el servidor de go, solo basta con ubicarnos en la carpeta que contiene la clase main.go y ejecutar el comando go run como se muestra a continuación:
```sh
$ cd /monitor_memoria_g9
$ go run main.go
```

### Carpetas y archivos del Proyecto
El proyecto se compone de 3 carpetas principales debido a su desgloce de desarrollo, estas son
- cpu_grupo9
- m_grupo9
- monitor_memoria_g9

#### cpu_grupo9
Contiene los archivos para los modulos kernel que obtienen información de los procesos del cpu
- Makefile - Crea y limpia los modulos del kernel de la clase .c 
- cpu_grupo9.c - Construye y finaliza las estructuras que obtienen los datos del modulo de kernel, utiliza las librerías propias del kernel de linux.
Los datos son guardados en formato Json, para poder ser utilizado posteriormente por la api de go
#### m_grupo9
Contiene los archivos que construyen los modulos de kernel para obtener la información de la RAM
- Makefile - Crea y limpia los modulos del kernel de la clase .c
- m_grupo9.c - Construye y finaliza la estructura que obtiene los datos del modulo de kernel para la RAM.
 Los datos son almacenados en formato Json para poder ser utilizados posteriormente por la api de go
#### monitor_memoria_g9
En esta carpeta se encuentra las plantillas de las paginas HTML que conforma la interfaz de la aplicación, la parte principal en ella es el archivo
- main.go

El cual es la clase principal de la api desarrollada en golang, esta se encarga de leer los datos generados por los modulos de kernel, y renderizar las vistas en html para mostrar esta información

Para poder utilizar esta clase es necesario previamente instalar la librería de gorilla/websocket la cual fue utilizada para crear los servicios de la api a través de socket's
Para instalarla usaremos el siguiente comando
```sh
$ go get github.com/gorilla/websocket
```

En esta clase se encuentran basicamente tres estructuras las cuales se utilizan para contener los datos que se obtienen del cpu, la ram y sus procesos.
Estas estructuras basicas son
- UtilizacionR - Contiene los datos de la RAM
- ListProceso - Contiene los datos generales de los procesos
- Proceso - Almacena la información de cada proceso individual