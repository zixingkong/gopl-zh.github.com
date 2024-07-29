package main

import "fmt"

//练习 4.6： 编写一个函数，原地将一个UTF-8编码的[]byte类型的slice中相邻的空格（参考unicode.IsSpace）替换成一个空格返回

func removeExtraSpace(s []byte) []byte {
	i := 0
	for j := 0; j < len(s); j++ {
		if s[j] == ' ' && s[j] == s[j-1] {
			continue
		} else {
			s[i] = s[j]
			i++
		}
	}
	return s[:i+1]
}

func main() {
	s := []byte("a b c d e  f  g")
	s = removeExtraSpace(s)
	fmt.Println(string(s))
}
