package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

//estructura memoria RAM
type UtilizacionR struct {
	Total      int `json:"Mem_Total"`
	Libre      int `json:"Mem_Libre"`
	Buffer     int `json:"Buffer"`
	Compartida int `json:"Compartida"`
	Consumida  int `json:"Consumida"`
}

//estructura de lista de procesos
type ListProceso struct {
	Lista                []Proceso `json:"Lista"`
	MemoriaTotal         int       `json:"MemoriaTotal"`
	ProcesosTotal        int       `json:"ProcesosTotal"`
	ProcesosEjecucion    int       `json:"ProcesosEjecucion"`
	ProcesosSuspendido   int       `json:"ProcesosSustendido"`
	ProcesosDetenidos    int       `json:"ProcesosDetenidos"`
	ProcesosZombie       int       `json:"ProcesosZombie"`
	ProcesosDesconocidos int       `json:"ProcesosDesconocidos"`
}

//Proceso de la estructura
type Proceso struct {
	PID     string    `json:"PID"`
	Nombre  string    `json:"Nombre"`
	Memoria string    `json:"Memoria"`
	Usuario string    `json:"Usuario"`
	Estado  string    `json:"Estado"`
	Hijos   []Proceso `json:"Hijos"`
}

//variables globales
var clients = make(map[*websocket.Conn]string)
var listUsers = make(map[string]string)
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func main() {
	// routes
	http.HandleFunc("/", serveFiles)
	http.HandleFunc("/ws", serveWs)
	http.HandleFunc("/grafica", viewGrafica)
	http.HandleFunc("/dashboard", viewDashboard)

	go enviarDatos()
	//start the server
	fmt.Println("servidor iniciado en el puerto 3000")
	http.ListenAndServe(":3000", nil)
}

func serveFiles(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	p := "." + r.URL.Path
	if p == "./" {
		p = "./static/login.html"
	}
	http.ServeFile(w, r, p)
}

func viewGrafica(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	p := "." + r.URL.Path
	p = "./static/GraficaM.html"
	http.ServeFile(w, r, p)
}

func viewDashboard(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	p := "." + r.URL.Path
	p = "./static/Dashboard.html"
	http.ServeFile(w, r, p)
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	log.Println("1------------")
	fmt.Println(r.Host)
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("2----------------")
		log.Println(err)
	}
	defer ws.Close()
	reader(ws)
}

func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error socket: %v", err)
			delete(clients, conn)
			break
		}
		fmt.Println(string(p))
		clients[conn] = string(p)
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

func enviarDatos() {
	for {
		for client := range clients {
			var value string = clients[client]
			log.Println(value)
			if value == "PRINCIPAL" {
				//---INDEX---
				listaProcess := getCPU()
				if listaProcess != nil {
					errW := client.WriteJSON(listaProcess)
					if errW != nil {
						log.Printf("Error: %v", errW)
						client.Close()
						delete(clients, client)
					}
				}
			} else if value == "RAM" {
				//----RAM----
				ram := getRAM()
				if ram != nil {
					errW := client.WriteJSON(ram)
					if errW != nil {
						log.Printf("Error: %v", errW)
						client.Close()
						delete(clients, client)
					}
				}
			} else {
				clients[client] = "PRINCIPAL"
				if i, err := strconv.Atoi(value); err == nil {
					proc, err := os.FindProcess(i)
					if err != nil {
						log.Println(err)
					}
					proc.Kill()
					log.Println("se elimino el proceso")
				}
				continue
			}
		}
		fmt.Println(len(clients))
		log.Printf("*******")
		time.Sleep(2000 * time.Millisecond)
	}
}

func getCPU() *ListProceso {
	data, err := ioutil.ReadFile("/proc/cpu_grupo9")
	if err != nil {
		fmt.Println("File reading error", err)
		return nil
	}
	strData := string(data)
	listaProcess := ListProceso{}
	json.Unmarshal([]byte(strData), &listaProcess)
	listaProcess.setNombresUsuario()
	return &listaProcess
}

func getRAM() *UtilizacionR {
	data, err := ioutil.ReadFile("/proc/m_grupo9")
	if err != nil {
		fmt.Println("File reading error", err)
		return nil
	}
	strData := string(data)
	fmt.Println(strData)
	infoMem := UtilizacionR{}
	json.Unmarshal([]byte(strData), &infoMem)
	fmt.Println(infoMem)
	return &infoMem
}

func (dato *Proceso) setNombreUsuario() {
	if listUsers[dato.Usuario] == "" {
		cmd, err := exec.Command("bash", "-c", "getent passwd "+dato.Usuario+" | cut -d: -f1").Output()
		if err != nil {
			log.Fatal(err)
		}
		dato.Usuario = strings.Trim(string(cmd[:]), " \n")
	} else {
		dato.Usuario = listUsers[dato.Usuario]
	}
	for i := range dato.Hijos {
		dato.Hijos[i].setNombreUsuario()
	}
}

func (obj *ListProceso) setNombresUsuario() {
	for i := range obj.Lista {
		obj.Lista[i].setNombreUsuario()
	}
}
