package main

import "fmt"

//练习 4.4： 编写一个rotate函数，通过一次循环完成旋转。

// 1. 循环,需要额外空间
func rangeRotate(slice []int, k int) []int {
	res := make([]int, len(slice))
	for index, val := range slice {
		res[(index+len(slice)-k)%len(slice)] = val
	}
	return res
}

// 2.通过反转完成,不需要额外空间
func rangeRotate2(slice []int, k int) {
	l := len(slice)
	k = l - k
	reverse4(slice)
	reverse4(slice[k:])
	reverse4(slice[:k])
}

func reverse4(s []int) {
	l := len(s)
	for i := range l / 2 {
		s[i], s[l-i-1] = s[l-i-1], s[i]
	}
}

func rotate(s []int, n int) []int {
	n %= len(s)
	return append(s[n:], s[:n]...)
}

func main() {
	a := []int{0, 1, 2, 3, 4, 5}
	a = rangeRotate(a, 2)
	fmt.Println(a) // "[2 3 4 5 0 1]"

	b := []int{0, 1, 2, 3, 4, 5}
	rangeRotate2(b, 2)
	fmt.Println(b) // "[2 3 4 5 0 1]"

}
