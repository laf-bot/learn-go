package main

/**
* 多变量声明
 */

var x, y int

// 这种因式分解关键字的写法一般用于全部变量
var (
	a int
	b bool
)

var c, d int = 1, 2

var e, f = 123, "hello"

// 这种不带声明格式的只能再函数体中出现
// g,h := 123,"hello"

func main() {
	g, h := 123, "hello"
	println(x, y, a, b, c, d, e, f, g, h)
}
