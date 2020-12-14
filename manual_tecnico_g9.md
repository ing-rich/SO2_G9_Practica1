## Universidad de San Carlos de Guatemala
## Facultad de ingeniería
## Escuela de vacaciones diciembre 2020
___
#### Manual técnico
#### Practica 1
#### Grupo 9

### Librerías a instalar
Para poder ejecutar la aplicación es necesario instalar algunas librerías previamente para las cuales se darán los comandos en Linux.
- gcc
- Make
- golang
- git

La distribución con la que fue desarrollado el sistema es Ubuntu 20.04 LTS

#### Instalar GCC
gcc es el compilador para ANSI C, C++ y que a su vez instalara make el cual nos permite compilar los archivos del kernel

Desde una terminal en Linux actualizaremos primero la lista de paquetes
```sh
$ sudo apt update
```
Ahora instalaremos el paquete build-essential
```sh
$ sudo apt install build-essential
```
El comando anterior instalara gcc, g++ y make por último instalaremos el manual con la documentación para GNU/Linux de desarrollo

```sh
$ sudo apt-get install manpages-dev
```
De esta forma ya tendremos instalado el compilador para los archivos del módulo y el kernel, lo siguiente será instalar go, para poder hacer uso del api y sus servicios.

#### Instalar GO
Lo primero será descargar el tarball de go, para este caso lo haremos desde consola con wget, tener en cuenta que para este proyecto se utilizó la versión go1.15.6
```sh
wget https://dl.google.com/go/go1.15.6.linux-amd64.tar.gz
```
Una vez descargado el archivo, lo siguiente será descomprimirlo desde su ubicación
```sh
sudo tar -xvf go1.15.6.linux-amd64.tar.gz -C /usr/local/
```
Lo siguiente consistirá en establecer la ubicación de go a través de la variable de entorno
```sh
export PATH=$PATH:/usr/local/go/bin
```
Después de realizar lo anterior solo quedara guardar y cargar la nueva variable en la terminal actual, para ello usamos el siguiente comando
```sh
source ~/.profile
```

Después de haber instalado todo lo anterior ya podremos compilar los archivos del proyecto y ejecutarlo para su uso
### Compilar Módulos de Kernel
Es necesario compilar los módulos de kernel, ya que estos crearan los archivos con los datos del CPU y la memoria, para esto debemos ubicarnos en la carpeta que contenga el archivo makefile y la clase .c
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
El proyecto se compone de 3 carpetas principales debido a su desglose de desarrollo, estas son
- cpu_grupo9
- m_grupo9
- monitor_memoria_g9

#### cpu_grupo9
Contiene los archivos para los módulos kernel que obtienen información de los procesos del CPU
- Makefile - Crea y limpia los módulos del kernel de la clase .c 
- cpu_grupo9.c - Construye y finaliza las estructuras que obtienen los datos del módulo de kernel, utiliza las librerías propias del kernel de Linux.
Los datos son guardados en formato Json, para poder ser utilizado posteriormente por el api de go
#### m_grupo9
Contiene los archivos que construyen los módulos de kernel para obtener la información de la RAM
- Makefile - Crea y limpia los módulos del kernel de la clase .c
- m_grupo9.c - Construye y finaliza la estructura que obtiene los datos del módulo de kernel para la RAM.
 Los datos son almacenados en formato Json para poder ser utilizados posteriormente por el api de go
#### monitor_memoria_g9
En esta carpeta se encuentra las plantillas de las paginas HTML que conforma la interfaz de la aplicación, la parte principal en ella es el archivo
- main.go

El cuál es la clase principal del api desarrollada en golang, esta se encarga de leer los datos generados por los módulos de kernel, y renderizar las vistas en HTML para mostrar esta información

Para poder utilizar esta clase es necesario previamente instalar la librería de gorilla/websocket la cual fue utilizada para crear los servicios del api a través de socket's
Para instalarla usaremos el siguiente comando
```sh
$ go get github.com/gorilla/websocket
```

En esta clase se encuentran básicamente tres estructuras las cuales se utilizan para contener los datos que se obtienen del CPU, la RAM y sus procesos.
Estas estructuras básicas son
- UtilizacionR - Contiene los datos de la RAM
- ListProceso - Contiene los datos generales de los procesos
- Proceso - Almacena la información de cada proceso individual
