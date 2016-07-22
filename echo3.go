package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], "\n"))
	fmt.Printf("%f [sec]\n", time.Since(start).Seconds())
}
