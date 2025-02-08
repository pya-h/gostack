package main

import (
	"fmt"

	stack "github.com/pya-h/gostack/stack"
)

func main() {
	s1 := stack.NewStack()
	// push test
	s1.Push("whatever")
	s1.Push(20, 30, 4.5, -3+4i, 23.65)

	// popout test
	fmt.Println(s1.PopOut())

	// object push
	s2 := stack.NewStack()
	s2.Push(1, 2, 34)
	s1.Push(s2)

	// pop test
	if v, ok := s1.Pop(); ok {
		fmt.Println(v)
	}
}
