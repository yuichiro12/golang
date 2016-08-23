package main

import "fmt"

var p = f()

func main() {
	v := 1
	incr(&v)
	fmt.Println(incr(&v))
	fmt.Println(incr(&v))
	fmt.Println(incr(&v))
}

func incr(p *int) int {
	*p++
	return *p
}

func f() *int {
	v := 1
	return &v
}
