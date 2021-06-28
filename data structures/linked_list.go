package main

import "fmt"

type node struct {
	data string
	next *node // pointer to the next node
}

type linkedList struct {
	length int
	head   *node // first list element
	tail   *node // last list element
}

func (l linkedList) Length() int {
	return l.length
}

func (l linkedList) Read() {
	for l.head != nil {
		fmt.Printf("%v ->", l.head.data)
		l.head = l.head.next
	}
	fmt.Println()
}

func (l *linkedList) Add(n *node) {
	if l.head == nil {
		l.head = n
		l.tail = n
		l.length++
	} else {
		l.tail.next = n
		l.tail = n
		l.length++
	}
}

func main() {
	list := linkedList{}
	node1 := node{
		data: "data1",
		next: nil,
	}
	node2 := node{
		data: "data2",
		next: nil,
	}

	list.Add(&node1)
	list.Add(&node2)

	fmt.Println(list.length)
	list.Read()
}
