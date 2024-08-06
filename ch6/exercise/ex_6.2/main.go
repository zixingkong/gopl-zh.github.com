//练习 6.2： 定义一个变参方法(*IntSet).AddAll(...int)，这个方法可以为一组IntSet值求和，比如s.AddAll(1,2,3)。

package main

import (
	"bytes"
	"fmt"
)

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

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) Len() int {
	count := 0
	for _, word := range s.words {
		for mask := uint(0); mask < 64; mask++ {
			if word&(1<<mask) != 0 {
				count++
			}
		}
	}
	return count
}

func (s *IntSet) Remove(x int) { // remove x from the set
	word, bit := x/64, uint(x%64)
	if word > len(s.words) {
		return
	}

	s.words[word] &^= 1 << bit
}

func (s *IntSet) Clear() { // remove all elements from the set
	s.words = nil
}

func (s *IntSet) Copy() *IntSet { // return a copy of the set
	n := &IntSet{}
	n.words = make([]uint64, len(s.words))
	copy(n.words, s.words)
	return n
}

func (s *IntSet) AddAll(ints ...int) {
	for _, n := range ints {
		s.Add(n)
	}
}

func main() {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.AddAll(1, 2, 3)
	fmt.Println(&x)
}
