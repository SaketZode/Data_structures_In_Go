package main

import (
	"datastructures/queue"
	"fmt"
)

func main() {
	queue := queue.Queue{}
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
	fmt.Println(queue.Capacity())
}
