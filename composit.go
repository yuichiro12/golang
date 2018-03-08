package main

import (
	"fmt"
)

type Currency int

func main() {
	var a [3]int
	var q [3]int = [3]int{1, 2, 3}
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])
	fmt.Println(q[0])
	for index, value := range a {
		fmt.Printf("%d %d\n", index, value)
	}

	const (
		USD Currency = iota
		EUR
		GBP
		RMB
	)

	symbol := [...]string{USD: "$", EUR: "€", GBP: "£", RMB: "¥"}
	for i, v := range symbol {
		fmt.Printf("%d %s\n", i, v)
	}

	var runes []rune
	for _, r := range "Hello, 世界" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)

	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d  cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z

}
