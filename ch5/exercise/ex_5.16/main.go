package main

import (
	"fmt"
	"strings"
)

// 练习5.16：编写多参数版本的strings.Join。
func join(strs ...string) string {
	if len(strs) < 2 {
		return ""
	}
	sep := strs[len(strs)-1]
	last := strs[len(strs)-2]
	temp := strs[:len(strs)-2]
	var res strings.Builder
	for _, s := range temp {
		res.WriteString(sep)
		res.WriteString(s)
	}
	res.WriteString(last)
	return res.String()
}

func main() {
	fmt.Println(join("aaa", "bbb", "ccc", "|"))
}
