package main

import (
	"bufio"
	"fmt"
	"os"
)

//练习 4.9： 编写一个程序wordfreq程序，报告输入文本中每个单词出现的频率。
//在第一次调用Scan前先调用input.Split(bufio.ScanWords)函数，这样可以按单词而不是按行输入。

func main() {
	hash := make(map[string]int)

	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		hash[input.Text()]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range hash {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Println("%#v", hash)
}
