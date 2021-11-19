package main

import "fmt"

// iota 特殊常量

func main() {

	const (
		i = 1 << iota
		j = 3 << iota

		k
		l
	)

	fmt.Println(i, j, k, l)
}
