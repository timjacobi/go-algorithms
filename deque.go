package main

import "errors"

type Deque struct {
	items [10]int
	head int
	tail int
}

func NewDeque() *Deque {
	return &Deque{
		head: 0,
		tail: 9,
	}
}

func (d *Deque) PushFront(value int) error {
	d.head++

	if d.head == len(d.items) {
		d.head = 0
	}

	if d.head == d.tail {
		return errors.New("Deque is full!")
	} else {
		d.items[d.head] = value
	}

	return nil
}

func (d *Deque) PopFront() int {
	item := d.items[d.head]

	d.head--

	if d.head == -1 {
		d.head = len(d.items) - 1
	}

	return item
}

func (d *Deque) PushBack(value int) error {
	d.tail--

	if d.tail == -1 {
		d.tail = len(d.items) - 1
	}

	if d.tail == d.head {
		return errors.New("Deque is full!")
	} else {
		d.items[d.tail] = value
	}

	return nil
}

func (d *Deque) PopBack() int {
	item := d.items[d.tail]

	d.tail++

	if d.tail == len(d.items){
		d.tail = 0
	}

	return item
}