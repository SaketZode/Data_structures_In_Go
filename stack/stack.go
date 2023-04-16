package stack

import "fmt"

type Stack struct {
	elements []int
}

func (stack *Stack) Push(element int) {
	stack.elements = append(stack.elements, element)
}

func (stack *Stack) Pop() (int, error) {
	if len(stack.elements) == 0 {
		return -1, fmt.Errorf("Stack is empty")
	}
	var element int = -1
	if len(stack.elements) > 0 {
		element = stack.elements[len(stack.elements)-1]
		if len(stack.elements) == 1 {
			stack.elements = []int{}
		} else {
			stack.elements = stack.elements[:len(stack.elements)-1]
		}
	}
	return element, nil
}

func (stack *Stack) Peek() (int, error) {
	if len(stack.elements) == 0 {
		return -1, fmt.Errorf("Stack is empty")
	}
	return stack.elements[len(stack.elements)-1], nil
}
