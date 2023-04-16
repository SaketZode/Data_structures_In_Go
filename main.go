package main

import (
	// "datastructures/queue"
	"datastructures/stack"
	"fmt"
)

func main() {
	/* queue := queue.Queue{}
	fmt.Println(queue.Capacity())
	queue.Enqueue(1)
	fmt.Println(queue.Capacity())
	queue.Enqueue(2)
	fmt.Println(queue.Capacity())
	queue.Enqueue(3)
	fmt.Println(queue.Capacity())
	queue.Enqueue(4)
	fmt.Println(queue.Capacity())
	queue.Enqueue(5)
	fmt.Println(queue.Capacity())
	queue.Display()
	// fmt.Println(queue.Capacity())
	queue.Dequeue()
	queue.Display()
	queue.Dequeue()
	queue.Display()
	queue.Dequeue()
	queue.Display()
	queue.Dequeue()
	queue.Display()
	queue.Dequeue()
	// queue.Dequeue()
	queue.Display()
	fmt.Println(queue.Capacity()) */

	stack := stack.Stack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.Push(6)
	top, err := stack.Peek()
	if err != nil {
		panic(err)
	}
	fmt.Println(top)
	stack.Pop()
	stack.Pop()
	stack.Pop()
	stack.Pop()
	stack.Pop()
	stack.Pop()
	top, err = stack.Peek()
	if err != nil {
		panic(err)
	}
	fmt.Println(top)
}
