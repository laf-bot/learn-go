package main

import "fmt"

/**

循环语句 - continue 语句

continue 不是跳出循环，而是跳过当前循环执行下一次循环语句

**/

func main() {

	var a int = 10

	for a < 20 {

		if a == 15 {

			a = a + 1
			continue
		}
		fmt.Println(a)
		a++
	}

	fmt.Println("----- continue -------")

	for i := 1; i <= 3; i++ {
		fmt.Println(i)

		for i2 := 11; i2 <= 13; i2++ {
			fmt.Println(i2)
			continue
		}
	}

}
