package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
)
type Request struct {
	Msg string `json:"msg"`
	Payload string `json:"payload"`
}

type Listener interface {
	handleMessage(match string, callback ...func(message string))
} 
type MyServer struct {
	socket net.Listener
	clients []*net.Conn
}

func (myServer *MyServer) init(port string) error{
	socket, err := net.Listen("tcp", ":" + port)
	if err != nil{
		return err
	}
	myServer.socket = socket
	return nil
}
func main(){

	err := godotenv.Load()
	handleError(err)
	var server MyServer
	
	port := os.Getenv("PORT")
	socket, err := net.Listen("tcp", ":"+ port)
	handleError(err)
	fmt.Printf("Listening on :%v",port)
	server.socket = socket

	defer socket.Close()
	for{
		fmt.Println("Waiting for clients, connected client:", server.clients)
        conn, err := socket.Accept()
        if err != nil {
            fmt.Println("Error connecting:", err.Error())
            return
        }
		server.clients = append(server.clients, &conn)

        fmt.Println("Client connected.")
        fmt.Println("Client " + conn.RemoteAddr().String() + " connected.")
		go server.handleConn(conn)
	}
}

func handleError(err error){
	if err != nil {
		log.Print("Error", err)
	}
}

func (server *MyServer) handleConn(conn net.Conn) {
	buffer, err := bufio.NewReader(conn).ReadBytes('\n')



	if err != nil {
		log.Println("Client", conn.RemoteAddr(), "disconnected")
		for i,c := range server.clients {
			if *c == conn { 
				server.clients = append(server.clients[:i], server.clients[i+1:]...)
			}
		}
		return
	}

	var req Request
	err = json.Unmarshal(buffer, &req)
	handleError(err)
	
	fmt.Fprintf(conn, "Received msg:%v, payload:%v\n", req.Msg, req.Payload)
	server.handleConn(conn)
}