package main

import (
	"io"
	"log"
	"net"
)

func main() {
	listenner, err := net.Listen("tcp", "localhost:1234")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listenner.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go io.Copy(conn, conn)
	}
}
