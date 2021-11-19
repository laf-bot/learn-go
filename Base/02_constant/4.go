package main

import "fmt"

// Go 常量_4 (iota 特殊常量)

const (
	a = iota
	b = iota
	c = iota
)

func main() {

	fmt.Println(a, b, c)
}
