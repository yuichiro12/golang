package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	content, _ := ioutil.ReadFile("index2.html")
	fmt.Fprintf(w, string(content))
}
