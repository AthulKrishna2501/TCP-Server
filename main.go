package main

import (
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp4", ":8081")
	if err != nil {
		log.Fatal(err)
	}

	for {
		log.Print("Waiting for client to connect")
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		log.Print("Client Connected")

		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	buf := make([]byte, 4096)
	_, err := c.Read(buf)

	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(3 * time.Second)

	c.Write([]byte("HTTP/1.1 200 OK\r\n\r\nHello World!\r\n"))
	c.Close()

}
