package main

// Go 常量_03

import (
	"fmt"
	"unsafe"
)

// 常量可以用函数计算表达式的值。常量表达式中，函数必须是内置函数，否则编译不过
const (
	a = "abc"
	b = len(a)
	c = unsafe.Sizeof(a)
)

func main() {

	fmt.Println(a, b, c)
}
