package main

import "fmt"

// 练习 4.7： 修改reverse函数用于原地反转UTF-8编码的[]byte。是否可以不用分配额外的内存？

func reverse3(s []byte, n int) {
	l := len(s)
	for i := range l / 2 {
		s[i], s[l-i-1] = s[l-i-1], s[i]
	}
}

func main() {
	s := []byte("1 2 3 4 5 6")
	reverse3(s, 2)
	fmt.Println(string(s))
}
