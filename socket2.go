package main

import (
	"net/http"
	"golang.org/x/net/websocket"
	"io"
)

func echoHandler(ws *websocket.Conn) {
	io.Copy(ws, ws)
}


func main() {
	http.Handle("/echo", websocket.Handler(echoHandler))
	http.Handle("/", http.FileServer(http.Dir("./")))
	if err := http.ListenAndServe(":9999", nil); err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}