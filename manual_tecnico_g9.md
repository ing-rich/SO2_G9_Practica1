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
### Proyecto
El proyecto se compone de 3 carpetas principales debido a su desgloce de desarrollo, estas son
- cpu_grupo9
- m_grupo9
- monitor_memoria_g9

