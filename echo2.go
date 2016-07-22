package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = "\n"
	}
	fmt.Println(s)
	fmt.Printf("%f [sec]\n", (time.Since(start)).Seconds())
}
