package main

import (
	"log"
	"net/http"

	"gopkg.in/igm/sockjs-go.v2/sockjs"
)

//const BUFSIZE = 1024

func main() {
	sockHandler := sockjs.NewHandler("/echo", sockjs.DefaultOptions, echoHandler)
	log.Fatal(http.ListenAndServe("localhost:8081", sockHandler))
}

func echoHandler(session sockjs.Session) {
	for {
		if msg, err := session.Recv(); err == nil {
			session.Send(session.ID() + msg)
			continue
		}
		break
	}
}

func
