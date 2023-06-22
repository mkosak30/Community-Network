package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"time"
)

var readBufferSize = 1024

// Server holds the server's information
type Server struct {
	IP      string
	Port    string
	Addr    string // IP & Port
	Network string
}

// Client holds the client's information
type Client struct {
	IP         string
	Port       string
	Network    string
	MsgType    string
	Period     time.Duration
	Connection net.Conn
}

// CreateServer will create a new server instance
func CreateServer(ip string, port string) *Server {
	return &Server{
		IP:      ip,
		Port:    port,
		Addr:    ip + ":" + port,
		Network: "tcp"}
}

// CreateClient will create a new client instance
func CreateClient(ip string, port string, msgType string, period time.Duration) *Client {
	return &Client{
		IP:      ip,
		Port:    port,
		Network: "tcp",
		MsgType: msgType,
		Period:  period}
}

// ReadFromClient will read from a specified client
func ReadFromClient(client *Client) {
	buffer := make([]byte, readBufferSize)

	for {
		n, err := client.Connection.Read(buffer)

		if err != nil {
			if err.Error() != "EOF" {
				fmt.Println(err)
				os.Exit(1)
			}
			break
		}

		if n > 0 {
			message := buffer[0:n]
			fmt.Println("Received: ", string(message))
		}

		time.Sleep(client.Period * time.Millisecond)
	}
}

// WriteToClient will write a specified message to a target client
func WriteToClient(client *Client, message string) {
	_, err := client.Connection.Write([]byte(message))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Sent: ", message)
}

// Listen will listen for incoming connections
func Listen(server *Server) {
	// Start listening on specified address
	listener, err := net.Listen(server.Network, server.Addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer listener.Close()

	go func() {
		for {
			clientConnection, err := listener.Accept()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			go ReadFromClient(CreateClient(clientConnection.RemoteAddr().String(), "", "", 0))
			go writeToClient(CreateClient(clientConnection.RemoteAddr().String(), "", "", 0))
		}
	}()
}

// Connect will connect to a remote server
func Connect(client *Client) {
	// Connect to remote server
	client.Connection, err := net.Dial(client.Network, client.IP+":"+client.Port)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Connected to ", client.IP)
}

// Serve will start the web server
func Serve(server *Server) {
	http.ListenAndServe(server.Addr, nil)
}

func main() {
	// Create the server and start listening for connections
	server := CreateServer("127.0.0.1", "9999")
	Listen(server)

	// Create the client and connect
	client := CreateClient("127.0.0.1", "9999", "tcp", 100)
	Connect(client)

	// Start the web server
	Serve(server)
}