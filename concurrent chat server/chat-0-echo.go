// Copyright 2013 Szymon Sobczak: http://about.me/ssobczak
// Licensed under the MIT license: http://opensource.org/licenses/MIT
// The above copyright notice shall be included in all copies or substantial portions of the Software.

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
