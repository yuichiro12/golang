package main

import (
	"fmt"
	"log"
	"net/http"

	"gopkg.in/igm/sockjs-go.v2/sockjs"
)

func main() {
	//http.HandleFunc("/", handler)
	//log.Fatal(http.ListenAndServe("localhost:1051", nil))
	sockHandler := sockjs.NewHandler("/echo", sockjs.DefaultOptions, echoHandler)
	log.Fatal(http.ListenAndServe("localhost:8081", sockHandler))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<div>chat</div><form><button type='type'>socket</button></form>")
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
