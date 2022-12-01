package main

import (
	"fmt"
)

type BasicTypes interface {
	int | int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | bool | float32 | float64 | complex64 | complex128 | string
}

type StackMethods interface {
	Push()
	Pop()
	PopOut()
}

type Stack[Type BasicTypes] struct {
	values []Type
}

func (stack *Stack[BasicTypes]) Push(value BasicTypes) {
	stack.values = append(stack.values, value)
}

func NewStack[Type BasicTypes]() Stack[Type] {
	return Stack[Type]{values: make([]Type, 0, 0)}
}

func (stack *Stack[BasicTypes]) Pop() (result BasicTypes, ok bool) {

	if last := len(stack.values) - 1; last >= 0 {
		result = stack.values[last]
		ok = true
		stack.values = stack.values[:last]
	} else {
		ok = false
	}
	return
}

func (stack *Stack[BasicTypes]) PopOut() (all []BasicTypes) {

	value, ok := stack.Pop()

	for ; ok; value, ok = stack.Pop() {
		all = append(all, value)
	}
	return
}

func main() {
	stack := NewStack[complex128]()
	stack.Push(10)
	stack.Push(20)

	fmt.Println(stack.PopOut())
}
