// 练习5.19： 使用panic和recover编写一个不包含return语句但能返回一个非零值的函数。
package main

import "fmt"

func nonZeroReturn() (result int) {
	defer func() {
		if r := recover(); r != nil {
			result = 42
		}
	}()

	panic("error")
}

func main() {
	fmt.Println(nonZeroReturn())
}
