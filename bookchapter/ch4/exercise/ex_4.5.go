package main

import "fmt"

//练习 4.5： 写一个函数在原地完成消除[]string中相邻重复的字符串的操作。

func duplicate(s []string) []string {
	i := 0
	for j := 1; j < len(s); j++ {
		if s[j] == s[i] {
			continue
		} else {
			i++
			s[i] = s[j]
		}
	}
	return s[:i+1]
}
func duplicate2(s []string) []string {
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			copy(s[i:], s[i+1:])
			s = s[:len(s)-1]
		}
	}
	return s
}
func main() {
	s := []string{"a", "a", "b", "b", "c", "c"}
	s = duplicate2(s)
	fmt.Println(s)
}
