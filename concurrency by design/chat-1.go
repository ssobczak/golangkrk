// Copyright 2013 Szymon Sobczak: http://about.me/ssobczak
// Licensed under the MIT license: http://opensource.org/licenses/MIT
// The above copyright notice shall be included in all copies or substantial portions of the Software.

package main

import (
	"fmt"
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

		go match(conn)
	}
}

var partner = make(chan io.ReadWriteCloser)

func match(user io.ReadWriteCloser) {
	fmt.Fprintln(user, "Waiting for a partner...")

	select {
	case p2 := <-partner:
		chat(user, p2)
	case partner <- user:
		// our user is taken by the other grorutine
	}
}

func chat(a, b io.ReadWriteCloser) {
	fmt.Fprintln(a, "Say hi!")
	fmt.Fprintln(b, "Say hi!")

	go io.Copy(a, b)
	io.Copy(b, a)
}
