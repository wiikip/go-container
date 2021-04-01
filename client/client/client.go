package client

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"time"
)

type Client struct {
	conn net.Conn
}


func (client *Client) Connect(address string, port string) error{
	conn, err := net.Dial("tcp", address + ":" + port)
	if err != nil {
		return err
	}
	client.conn = conn
	return nil
}

func (client *Client) SendMsg(ctx context.Context,msg string, payload interface{}, callback func(string)) error {
	if client.conn == nil {
		log.Println("Error: client is not connected")
	}
	conn := client.conn
	_ , err := fmt.Fprintf(conn,`{"msg": "%v", "payload": "%v"}` + "\n", msg, payload )
	if err != nil {
		return err
	}
	c := make(chan string)

	cancelCtx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancel()
	go func(){
		buffer, err := bufio.NewReader(conn).ReadBytes('\n')
		if err != nil {
			cancel()
			fmt.Println("Failed to read response")
			return
		}
		c <- string(buffer)
	}()

	select{
	case <- cancelCtx.Done():
		return cancelCtx.Err()
	case s:=  <- c:
		callback(s)
	}
	return nil
}
func (client *Client) Disconnect(){
	client.conn.Close()
}