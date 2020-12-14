## Universidad de San Carlos de Guatemala
## Facultad de ingeniería
## Escuela de vacaciones diciembre 2020
___
#### Manual técnico
#### Practica 1
#### Grupo 9
### Acceso a la aplicación
Para poder utilizar la aplicación de monitorización de procesos, lo primero que debemos hacer es ingresar al siguiente enlace:

http://34.122.228.97:3000/

Este nos dirigira a la pagina de login
![N|Inicio](https://github.com/ing-rich/SO2_G9_Practica1/blob/main/imagenes/login.PNG)

En esta pagina debemos agregar las credenciales para tener acceso, en este caso podemos utilizar las siguientes:
- usuario: admin
- contraseña: admin
![N|Usuario](https://github.com/ing-rich/SO2_G9_Practica1/blob/main/imagenes/usuario.PNG)

### Panel de control
Despues de habernos loggeado se nos desplegara la primera pagina en la cual encontraremos información general de los procesos y especifica de cada uno de ellos

#### Datos generales de los procesos
Entre los datos generales de los procesos encontraremos 
> TOTAL PROCESOS
> PROCESOS EN EJECUCIÓN
> PROCESOS SUSPENDIDOS
> PROCESOS DETENIDOS
> PROCESOS ZOMBIE

#### Datos
En la parte de abajo podremos ver una tabla que contiene una lista de los procesos padre donde nos da mas datos de cada uno de ellos, los cuales son
> PID(Identificador del proceso)
> Nombre (nombre del proceso)
> Usuario (usuario que lo ejecuto)
> Memoria (porcentaje de memoria utilizada por el proceso)
> Estado (Estado en el que se encuentra el proceso actualmente)

![N|Tablero](https://github.com/ing-rich/SO2_G9_Practica1/blob/main/imagenes/panelControl.PNG)
#### Lista de procesos
Algunos de los procesos posee una lista de otros procesos hijos, los cuales tambien podremos consultar, para ello bastara presionar el boton que se encuentra a la par de cada uno y de esta forma se nos desplegara la lista desde la cual podremos dar kill al proceso, lo cual veremos posteriormente
![N|botonLista](https://github.com/ing-rich/SO2_G9_Practica1/blob/main/imagenes/botonLista.PNG)
![N|Hijos](https://github.com/ing-rich/SO2_G9_Practica1/blob/main/imagenes/HijosProceso.PNG)

#### Terminar procesos
De la misma forma que vimos anteriormente, a la par de cada proceso encontraremos un boton que nos permitira dar KILL al proceso con solo presionarlo, podemos apoyarnos de la barra de busqueda ubicada en la parte superior de la tabla para buscar un proceso en específico.
![N|stress](https://github.com/ing-rich/SO2_G9_Practica1/blob/main/imagenes/stress.PNG)
![N|botonEliminar](https://github.com/ing-rich/SO2_G9_Practica1/blob/main/imagenes/botonEliminar.PNG)
>se debe tomar en cuenta, que no todos los procesos pueden ser eliminados, ya que para ello se necesitan permisos de usuario

### Menu
Podremos utilizar las opciones ubicadas en el lado izquierdo de la pantalla, aca encontraremos las siguientes:
- SALIR - Esta cierra la sesión actual y debemos hacer login de nuevo si deseamos regresar al panel administrativo
- Tablero Mando - Nos lleva a la administración de procesos del CPU del cual ya hemos hablado anteriormente
- Gráficas Memoria RAM - Esta sera la opción que nos llevara a la pagina que muestra los datos sobre la memoria
![N|Menu](https://github.com/ing-rich/SO2_G9_Practica1/blob/main/imagenes/opciones.PNG)

### Gráficas Memoria RAM
Esta pagina nos muestra información sobre la memoria del servidor, en la parte superior encontraremos los siguientes datos
- TOTAL MEMORIA (Memoria total del servidor)
- MEMORIA CONSUMIDA (Memoria consumida )
- PORCENTAJE (porcentaje de memoria consumida)

En la parte inferior nos encontramos con una grafica en tiempo real del consumo de memoria del servidor.
![N|grafica](https://github.com/ing-rich/SO2_G9_Practica1/blob/main/imagenes/grafica.PNG)

