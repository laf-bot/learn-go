package main

import "fmt"

/**
循环语句- break语句

break 用于循环语句中跳出循环，并开始循环之后的语句。

**/

func main() {

	/* 定义局部变量 */

	var a int = 10

	for a < 20 {
		fmt.Println(a)
		a++
		if a > 15 {
			break
		}
	}

	// 不使用标记

	fmt.Println("-----break-----")
	for i := 1; i <= 3; i++ {
		fmt.Println(i)

		for i2 := 11; i2 <= 13; i2++ {
			fmt.Println(i2)
			break
		}
	}

	fmt.Println("-----break label-----")

re:
	for i := 1; i <= 3; i++ {
		fmt.Println(i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Println(i2)
			break re
		}
	}
}
