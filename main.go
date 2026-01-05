package main

import (
	"fmt"
	"net"
)
func handleConnection(conn net.Conn){
	defer conn.Close()
	fmt.Println("Client connected:", conn.RemoteAddr())

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
	conn.Write([]byte("Server Respone: " + message))
	}
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