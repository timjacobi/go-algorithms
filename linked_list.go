package main

import (
	"errors"
	"fmt"
)

type ListNode struct {
	Prev *ListNode
	Next *ListNode
	Key int
}

type DoublyLinkedList struct {
	Nil *ListNode
}

func NewLinkedList() *DoublyLinkedList {
	linkedList := &DoublyLinkedList{}
	nil := &ListNode{}
	nil.Next = nil
	nil.Prev = nil
	linkedList.Nil = nil


	return linkedList
}

func (ll *DoublyLinkedList) Insert(key int){
	listNode := &ListNode{ Key: key }

	listNode.Next = ll.Nil.Next
	ll.Nil.Next.Prev = listNode
	ll.Nil.Next = listNode
	listNode.Prev = ll.Nil
}

func (ll *DoublyLinkedList) Search(key int) (*ListNode, error) {
	current := ll.Nil.Next

	for current != ll.Nil {
		if current.Key == key {
			return current, nil
		}
		current = current.Next
	}

	return nil, errors.New("Key not found!")
}

func (ll *DoublyLinkedList) Delete(key int) error {
	node, err := ll.Search(key)

	if err != nil {
		return err
	}

	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev

	return nil
}

type mapFn func(*ListNode)

func (ll *DoublyLinkedList) Map(callback mapFn){
	current := ll.Nil.Next

	for current != ll.Nil {
		callback(current)
		current = current.Next
	}
}

func Print(listNode *ListNode){
	fmt.Println(listNode.Key)
}

func main() {
	linkedList := NewLinkedList()
	linkedList.Insert(1)
	linkedList.Insert(3)
	linkedList.Insert(2)

	fmt.Println("Printing list")
	linkedList.Map(Print)

	fmt.Println("Printing node with key 2")
	result, _ := linkedList.Search(2)
	fmt.Println(result)

	fmt.Println("Deleting node with key 3 and printing list")
	linkedList.Delete(3)
	linkedList.Map(Print)


}