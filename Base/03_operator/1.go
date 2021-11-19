package main

import "fmt"

/**
* Operator 运算符-算数运算符

+  相加
-  相减
*  相乘
/  相除
%  求余
++ 自增
-- 自减

**/

func main() {

	var a int = 21
	var b int = 10
	var c int

	c = a + b
	fmt.Printf("第一行 - c 的值为 %d\n", c)
	c = a - b
	fmt.Printf("第二行 - c 的值为 %d\n", c)
	c = a * b
	fmt.Printf("第三行 - c 的值为 %d\n", c)
	c = a / b
	fmt.Printf("第四行 - c 的值为 %d\n", c)
	c = a % b
	fmt.Printf("第五行 - c 的值为 %d\n", c)
	a++
	fmt.Printf("第六行 - a 的值为 %d\n", a)
	a = 21
	a--
	fmt.Printf("第七行 - a 的值为 %d\n", a)
}
