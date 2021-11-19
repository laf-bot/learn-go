package main

import "fmt"

/**
循环语句-for循环

**/

func main() {

	sum := 0

	for i := 0; i <= 10; i++ {
		sum += i
	}

	fmt.Println(sum)

	sum = 1

	for sum <= 10 {
		sum += sum
	}

	fmt.Println(sum)

}
