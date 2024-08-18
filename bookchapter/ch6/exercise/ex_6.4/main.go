//练习6.4: 实现一个Elems方法，返回集合中的所有元素，用于做一些range之类的遍历操作。

package main

import "fmt"

type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}
func (s *IntSet) Elems() []int {
	e := make([]int, 0)
	for i, word := range s.words {
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				e = append(e, i*64+j)
			}
		}
	}
	return e
}
func main() {
	var x IntSet
	x.Add(1)
	x.Add(2)
	x.Add(3)
	y := x.Elems()
	fmt.Println(y)
}
