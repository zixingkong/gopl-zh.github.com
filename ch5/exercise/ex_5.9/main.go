package main

//练习 5.9： 编写函数expand，将s中的"foo"替换为f("foo")的返回值。

import (
	"fmt"
)

func main() {
	foo := "foo"
	fmt.Println(expand(foo, replace))
}

func expand(s string, f func(string) string) string {
	return f(s)
}

func replace(s string) string {
	return s + "-next"
}
