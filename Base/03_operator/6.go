package main

import "fmt"

/**
operator 运算符-其他运算符

& 返回变量存储地址
* 指针变量

**/

func main() {
	var a int = 4
	var b int32
	var c float32
	var ptr *int

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)

	ptr = &a
	fmt.Println(*ptr)

}
