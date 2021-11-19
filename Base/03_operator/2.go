package main

import "fmt"

/**
* Operator 运算符-关系运算符

== 检测两值是否相等
!= 检测两值是否不相等
>  检测左边值是否大于右边值
<  检测左边值是否小于右边值
>= 检测左边值是否大于等于右边值
<= 检测左边值是否小于等于右边值

**/

func main() {

	var a int = 21
	var b int = 20

	if a == b {
		fmt.Println("a等于b")
	} else {
		fmt.Println("a不等于b")
	}

	if a < b {
		fmt.Println("a小于b")
	} else {
		fmt.Println("a不小于b")
	}

	if a > b {
		fmt.Println("a大于b")
	} else {
		fmt.Println("a不大于b")
	}

	a = 5
	b = 20

	if a <= b {
		fmt.Println("a小于等于b")
	}

	if b >= a {
		fmt.Println("b大于等于a")
	}
}
