package main

//练习 4.1： 编写一个函数，计算两个SHA256哈希码中不同bit的数目。（参考2.6.2节的PopCount函数。)

import (
	"crypto/sha256"
	"fmt"
)

func PopCount(x byte) int {
	count := 0
	for x != 0 {
		x = x & (x - 1)
		count++
	}
	return count
}

func DiffBits(h1, h2 [32]byte) int {
	count := 0
	for i := range h1 {
		count += PopCount(h1[i] ^ h2[i])
	}
	return count
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	fmt.Println(DiffBits(c1, c2))
}
