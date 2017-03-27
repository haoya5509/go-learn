package main

import (
	"bytes"
	"fmt"
)

const OP_BIT = 32 << (^uint(0) >> 63)

type IntSet struct {
	words []uint
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/OP_BIT, uint(x%OP_BIT)
	return word < len(s.words) && (s.words[word] & (1 << bit) != 0)
}

func (s *IntSet) Add(x int) {
	word, bit := x/OP_BIT, uint(x%OP_BIT)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= (1 << bit)
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		} else {
			s.words[i] = 0
		}
	}
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &^= tword
		} else {
			break
		}
	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < OP_BIT; j++ {
			if word & (1 << uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", OP_BIT*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) Len() int {
	cnt := 0
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < OP_BIT; j++ {
			if word & (1 << uint(j)) != 0 {
				cnt++
			}
		}
	}
	return cnt
}

func (s *IntSet) Remove(x int) {
	word, bit := x/OP_BIT, uint(x%OP_BIT)
	if (word > len(s.words)) {
		return
	}
	s.words[word] &^= (1 << bit)
}

func (s *IntSet) Clear() {
	s.words = []uint{}
}
