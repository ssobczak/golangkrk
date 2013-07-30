// Copyright 2013 Szymon Sobczak: http://about.me/ssobczak
// Licensed under the MIT license: http://opensource.org/licenses/MIT
// The above copyright notice shall be included in all copies or substantial portions of the Software.

package main

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.Handle("/ws", websocket.Handler(ws_handler))
	http.HandleFunc("/", handler)

	if err := http.ListenAndServe("localhost:1234", nil); err != nil {
		log.Fatal(err)
	}
}

func ws_handler(conn *websocket.Conn) {
	for {
		var msg string
		fmt.Fscan(conn, &msg)
		fmt.Println("Recieved:", msg)
		fmt.Fprint(conn, "You said ", msg)
	}
}

func handler(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprint(resp, `
		var sock = new WebSocket("ws://localhost:1234/ws");
		sock.onmessage = function(msg) { console.log("Recieved:", msg.data); }
		sock.send("Hello!\n");
	`)
}
