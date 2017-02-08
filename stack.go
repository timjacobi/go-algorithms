package main

import (
	"fmt"
	"errors"
)

type Stack []int

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(value int){
	*s = append(*s, value)
}

func (s *Stack) Pop() (int, error) {
	if s.Size() == 0 {
		return 0, errors.New("No more elements")
	}
	lastIdx := len(*s) - 1
	value := (*s)[lastIdx]
	*s = (*s)[:lastIdx]
	return value, nil
}

func (s *Stack) Size() int {
	return len(*s)
}

type StackBasedQueue struct {
	temp    *Stack
	ordered *Stack
}

func NewStackBaseQueue() *StackBasedQueue {
	return &StackBasedQueue{
		temp: NewStack(),
		ordered: NewStack(),
	}
}

func (sbq *StackBasedQueue) Enqueue(value int) {
	for sbq.ordered.Size() > 0 {
		value, _ := sbq.ordered.Pop()
		sbq.temp.Push(value)
	}

	sbq.ordered.Push(value)

	for sbq.temp.Size() > 0 {
		value, _ := sbq.temp.Pop()
		sbq.ordered.Push(value)
	}
}

func (sbq *StackBasedQueue) Dequeue() (int, error) {
	return sbq.ordered.Pop()
}

func main(){
	fmt.Println("Stack operations")
	stack := NewStack()
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	fmt.Println(stack.Size())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Size())
	stack.Push(40)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())

	fmt.Println("StackBasedQueue operations")

	stackBasedQueue := NewStackBaseQueue()

	stackBasedQueue.Enqueue(10)
	stackBasedQueue.Enqueue(20)
	stackBasedQueue.Enqueue(30)

	fmt.Println(stackBasedQueue.Dequeue())
	fmt.Println(stackBasedQueue.Dequeue())
	fmt.Println(stackBasedQueue.Dequeue())
	fmt.Println(stackBasedQueue.Dequeue())
}