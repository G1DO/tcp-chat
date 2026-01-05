package main

import (
	"fmt"
	"net"
	"sync"
)
var (
	clients = make([]net.Conn,0)
	mutex = sync.Mutex{}
)
func addClient(conn net.Conn) {
    mutex.Lock()
    clients = append(clients, conn)
    mutex.Unlock()
}
func removeClient(conn net.Conn) {
    mutex.Lock()
	// it gives 2 things 1 is index i from 0 to clints number and conn itslf its the connection itsself maybe connA connB 
     for i, c := range clients{
		if c == conn {
			clients = append(clients[:i], clients[i+1:]...)
			break
		}
	}
    mutex.Unlock()
}
func broadcast(message string, sender net.Conn){
	mutex.Lock()
	for _,client := range clients{
		if client != sender{
			client.Write([]byte(message))
		}
	}
	mutex.Unlock()
}

func handleConnection(conn net.Conn){
	defer conn.Close()
	defer removeClient(conn)
	fmt.Println("Client connected:", conn.RemoteAddr())
addClient(conn)
	//read
buffer := make([]byte,1024)
	for {
	
	n, err := conn.Read(buffer)  
	if err != nil {
    fmt.Println("Client disconnected:", conn.RemoteAddr())
    break
}
	message := string(buffer[:n])

	fmt.Printf("Received: %s", message)
    broadcast(message, conn)	}
}

func main(){
	//listen on port 9999
	listener , err := net.Listen("tcp",":9999")
	if err != nil{
		fmt.Println("Error:",err)
		return
	}
	defer listener.Close()
	fmt.Println("Server listening on :9999")
//accept
for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Error:", err)
            continue
        }
        go handleConnection(conn)  // goroutine!
    }
}