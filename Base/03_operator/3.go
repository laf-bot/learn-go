package main

import "fmt"

/**
* Operator 运算符-逻辑运算符

&&  逻辑AND运算符
||  逻辑OR运算符
！  逻辑NOT运算符

**/

func main() {

	var a bool = true
	var b bool = false

	if a && b {
		fmt.Println("01 条件为TRUE")
	} else {
		fmt.Println("02 条件为FALSE")
	}

	if a || b {
		fmt.Println("03 条件为TRUE")
	}

	a = false
	b = true

	if a && b {
		fmt.Println("04 条件为TRUE")
	} else {
		fmt.Println("05 条件为FALSE")
	}

	if !(a && b) {
		fmt.Println("06 条件为TRUE")
	}
}
