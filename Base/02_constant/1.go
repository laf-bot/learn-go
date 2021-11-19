package main

import "fmt"

/**
* Go 常量
 */

func main() {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int

	const a, b, c = 1, false, "str"

	area = LENGTH * WIDTH

	fmt.Printf("面积为: %d", area)

	fmt.Println()

	fmt.Println(a, b, c)
}
