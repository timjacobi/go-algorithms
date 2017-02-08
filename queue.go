package main

import (
	"errors"
	"fmt"
)

type Queue []int

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(value int) {
	*q = append(*q, value)
}

func (q *Queue) Dequeue() (int, error) {
	if q.Size() == 0 {
		return 0, errors.New("No more elements")
	}
	value := (*q)[0]
	*q = (*q)[1:]
	return value, nil
}

func (q *Queue) Size() int {
	return len(*q)
}

type QueueBasedStack struct {
	temp    *Queue
	ordered *Queue
}

func (qbs *QueueBasedStack) Push(value int){
	for qbs.ordered.Size() > 0 {
		item, _ := qbs.ordered.Dequeue()
		qbs.temp.Enqueue(item)
	}

	qbs.ordered.Enqueue(value)

	for qbs.temp.Size() > 0 {
		item, _ := qbs.temp.Dequeue()
		qbs.ordered.Enqueue(item)
	}
}

func (qbs *QueueBasedStack) Pop() (int, error) {
	return qbs.ordered.Dequeue()
}

func NewQueueBasedStack() *QueueBasedStack {
	return &QueueBasedStack{
		temp: NewQueue(),
		ordered: NewQueue(),
	}
}

func main(){
	fmt.Println("Queue operations")
	q := NewQueue()
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())

	fmt.Println("QueueBasedStack operations")
	qbs := NewQueueBasedStack()
	qbs.Push(10)
	qbs.Push(20)
	qbs.Push(30)
	fmt.Println(qbs.Pop())
	qbs.Push(40)
	fmt.Println(qbs.Pop())
	fmt.Println(qbs.Pop())
	fmt.Println(qbs.Pop())
	fmt.Println(qbs.Pop())
}