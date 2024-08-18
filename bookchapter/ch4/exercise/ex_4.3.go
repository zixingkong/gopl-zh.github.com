package main

import "fmt"

// 练习 4.3： 重写reverse函数，使用数组指针代替slice。
// reverse reverses a slice of ints in place.
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverse2(s *[6]int) {
	length := len(s)
	for index := range length / 2 {
		s[index], s[len(s)-1-index] = s[len(s)-1-index], s[index]
	}

}

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Println(a) // "[5 4 3 2 1 0]"
	b := [6]int{0, 1, 2, 3, 4, 5}
	reverse2(&b)
	fmt.Println(b) // "[5 4 3 2 1 0]"
}
