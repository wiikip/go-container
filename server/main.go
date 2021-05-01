package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"

	kubeclient "github.com/wiikip/go-container-server/kube"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

const (
	DOCKER_BUILD = "DOCKER_BUILD"
	GET_PODS     = "GET_PODS"
)

type Request struct {
	Msg     string `json:"msg"`
	Payload string `json:"payload"`
}

type Response struct {
	Payload string
}

type ResponsePods struct {
	Payload []PodsData `json:"payload"`
}

type PodsData struct {
	Name string `json:"name"`
	Uri  string `json:"uri"`
}

type HandleFunc = func(Request, net.Conn)
type Listener interface {
	handleMessage(match string, callback ...func(message string))
}
type MyServer struct {
	socket     net.Listener
	clients    []*net.Conn
	Handlers   map[string]HandleFunc
	KubeClient *kubeclient.KubeClient
}

func (myServer *MyServer) init(port string) error {
	socket, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}
	myServer.Handlers = make(map[string]HandleFunc)
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	var kubeClient kubeclient.KubeClient
	kubeClient.ClientSet = *clientset

	myServer.KubeClient = &kubeClient
	myServer.socket = socket
	return nil
}

func (myServer *MyServer) waitForConnections() {
	defer myServer.socket.Close()
	for {
		fmt.Println("Waiting for clients, connected client:", myServer.clients)
		conn, err := myServer.socket.Accept()
		if err != nil {
			fmt.Println("Error connecting:", err.Error())
			return
		}
		myServer.clients = append(myServer.clients, &conn)

		fmt.Println("Successfully loaded kube client")
		fmt.Println("Client connected.")
		fmt.Println("Client " + conn.RemoteAddr().String() + " connected.")
		go myServer.handleConn(conn)
	}

}
func main() {
	var server MyServer

	port := os.Getenv("PORT")

	err := server.init(port)
	if err != nil {
		log.Panic()
	}
	fmt.Printf("Listening on :%v", port)

	server.addHandler(DOCKER_BUILD, func(req Request, conn net.Conn) {
		fmt.Println("Handler 1 triggered")
		pod, err := server.KubeClient.CreatePod(kubeclient.GetPod(req.Payload))
		if err != nil {
			log.Println("ERROR:", err)
		}
		log.Println("Pod Created:", pod)

	})
	server.addHandler(GET_PODS, func(req Request, conn net.Conn) {
		var response ResponsePods
		fmt.Println("Handler 2 triggered")
		pods, err := server.KubeClient.GetPods()
		if err != nil {
			log.Println("ERROR: ", err)
			return
		}
		for _, pod := range pods.Items {
			response.Payload = append(response.Payload, PodsData{pod.Name, pod.Spec.Containers[0].Image})
		}

		parsedRes, err := json.Marshal(response)
		fmt.Println("Sending ", parsedRes)
		if err != nil {
			log.Println("ERROR: ", err)
			return
		}
		fmt.Fprintf(conn, "%s \n", parsedRes)

	})
	server.waitForConnections()
}

func (server *MyServer) addHandler(msg string, handler func(Request, net.Conn)) {
	server.Handlers[msg] = handler
}
func (server *MyServer) handleConn(conn net.Conn) {
	buffer, err := bufio.NewReader(conn).ReadBytes('\n')
	if err != nil {
		log.Println("Client", conn.RemoteAddr(), "disconnected")
		for i, c := range server.clients {
			if *c == conn {
				server.clients = append(server.clients[:i], server.clients[i+1:]...)
			}
		}
		return
	}

	var req Request
	log.Println("Received: ", string(buffer))

	err = json.Unmarshal(buffer, &req)
	if err != nil {
		log.Println("Received message format is not supported:", err)
	}
	go server.Handlers[req.Msg](req, conn)

	server.handleConn(conn)
}
