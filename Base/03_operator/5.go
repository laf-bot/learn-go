package main

import "fmt"

/**
* Operator 运算符-赋值运算符

=   简单的赋值运算符，将一个表达式赋值给一个左值
+=  相加后再赋值
-=  相见后再赋值
*=  相乘后再赋值
/=  相除后再赋值
%=  求余后再赋值
<<= 左移后赋值
>>= 右移后赋值
&= 按位与后赋值
^= 按位异或后赋值
|= 按位或后赋值
**/

func main() {

	var a int = 21
	var c int

	c = a
	fmt.Println("c = a = ", c)

	c += a
	fmt.Println("c += a = ", c)

	c -= a
	fmt.Println("c -= a = ", c)

	c *= a
	fmt.Println("c *= a = ", c)

	c /= a
	fmt.Println("c /= a = ", c)

	c = 200

	c <<= 2
	fmt.Println("c <<= 2 = ", c)

	c >>= 2
	fmt.Println("c >>= 2 = ", c)

	c &= 2
	fmt.Println("c &= 2 = ", c)

	c ^= 2
	fmt.Println("c ^= 2 = ", 0)

	c |= 2
	fmt.Println("c |= 2 = ", c)

}
