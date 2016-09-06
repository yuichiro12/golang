package main

import (
	"log"
	"net/http"

	"gopkg.in/igm/sockjs-go.v2/sockjs"
)

func main() {
	sockHandler := sockjs.NewHandler("/echo", sockjs.DefaultOptions, echoHandler)
	sendHandler := sockjs.NewHandler("/send", sockjs.DefaultOptions, sendHandler)
	log.Fatal(http.ListenAndServe("localhost:8081", sockHandler))
	log.Fatal(http.ListenAndServe("localhost:8082", sendHandler))
}

func sendHandler(session sockjs.Session) {
	session.Send("hi")
}

func echoHandler(session sockjs.Session) {
	for {
		if msg, err := session.Recv(); err == nil {
			session.Send(msg)
			continue
		}
		break
	}
}
