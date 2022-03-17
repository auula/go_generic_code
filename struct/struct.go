// Open Source: MIT License
// Author: Leon Ding <ding@ibyte.me>
// Date: 2022/3/17 - 12:31 下午 - UTC/GMT+08:00

package main

import "fmt"

type Element interface {
	int64 | float64 | string
}

type Stack[V Element] struct {
	size  int
	value []V
}

func (s *Stack[V]) Push(v V) {
	s.value = append(s.value, v)
	s.size++
}

func (s *Stack[V]) Pop() V {
	e := s.value[s.size-1]
	if s.size != 0 {
		s.value = s.value[:s.size-1]
		s.size--
	}
	return e
}

func main() {

	// INT STACK
	strS := Stack[int64]{}
	strS.Push(1)
	strS.Push(2)
	strS.Push(3)
	fmt.Println(strS.Pop())
	fmt.Println(strS.Pop())
	fmt.Println(strS.Pop())

	// FLOAT STACK
	floatS := Stack[float64]{}
	floatS.Push(1.1)
	floatS.Push(2.2)
	floatS.Push(3.3)
	fmt.Println(floatS.Pop())
	fmt.Println(floatS.Pop())
	fmt.Println(floatS.Pop())

}
