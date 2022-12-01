package main

import (
	"fmt"
)

type StackMethods interface {
	Push()
	Pop()
	PopOut()
	NewStack()
}

type Stack struct {
	values []interface{}
}

func (stack *Stack) Push(values ...interface{}) {
	for _, v := range values {
		stack.values = append(stack.values, v)
	}
}

func NewStack() Stack {
	return Stack{values: make([]interface{}, 0, 0)}
}

func (stack *Stack) Pop() (result interface{}, ok bool) {

	if last := len(stack.values) - 1; last >= 0 {
		result = stack.values[last]
		ok = true
		stack.values = stack.values[:last]
	} else {
		ok = false
	}
	return
}

// pop all items as an array
func (stack *Stack) PopOut() (all []interface{}) {

	value, ok := stack.Pop()

	for ; ok; value, ok = stack.Pop() {
		all = append(all, value)
	}
	return
}

func main() {
	stack := NewStack()
	// push test
	stack.Push("whatever")
	stack.Push(20, 30, 4.5, -3+4i, 23.65)

	// popout test
	fmt.Println(stack.PopOut())

	// object push
	s := NewStack()
	s.Push(1, 2, 34)
	stack.Push(s)

	// pop test
	if v, ok := stack.Pop(); ok {
		fmt.Println(v)
	}
}
