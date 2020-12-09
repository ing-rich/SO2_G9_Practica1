package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

//estructura memoria RAM
type usoRAM struct {
	Total      int `json:"Mem_Total"`
	Consumida  int `json:"Mem_Consumida"`
	Porcentaje int `json:"Mem_porcentaje"`
}

//estructura de lista de procesos
type listaProcesos struct {
	ProcesosTotal       int       `json:"ProcesosTotal"`
	ProcesosEjecucion   int       `json:"ProcesosEjecucion"`
	ProcesosSuspendidos int       `json:"ProcesosSuspendidos"`
	ProcesosDetenidos   int       `json:"ProcesosDetenidos"`
	ProcesosZombies     int       `json:"ProcesosZombie"`
	Procesos            []Proceso `json:"Lista"`
}

//estructura de los Procesos
type Proceso struct {
	PID     string    `json:"PID"`
	Nombre  string    `json:"Nombre"`
	Usuario string    `json:"Usuario"`
	Estado  string    `json:"Estado"`
	Memoria string    `json:"Memoria"`
	Hijos   []Proceso `json:"Hijos"`
}

//variables globales
var clients = make(map[*websocket.Conn]string)
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func main() {
	// routes
	http.HandleFunc("/", serveFiles)
	http.HandleFunc("/ws", serveWs)
	//start the server
	http.ListenAndServe(":3000", nil)
}

func serveFiles(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	p := "." + r.URL.Path
	if p == "./" {
		p = "./static/index.html"
	}
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
			log.Printf("error: %v", err)
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
