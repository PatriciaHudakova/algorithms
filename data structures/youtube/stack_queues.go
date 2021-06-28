/*
A basic implementation of a stack and a queue in go
*/
package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) push(i int) {
	s.items = append(s.items, i)
}

func (s *Stack) pop() {
	s.items = s.items[:len(s.items)-1]
}

type Queue struct {
	items []int
}

func (q *Queue) enqueue(i int) {
	q.items = append(q.items, i)
}

func (q *Queue) dequeue() {
	q.items = q.items[1:]
}

func main() {
	// stack
	myStack := Stack{items: []int{1, 2, 3}}
	myStack.push(4)
	myStack.push(5)
	myStack.pop()
	fmt.Println(myStack)

	// queue
	myQueue := Queue{items: []int{1, 2, 3}}
	myQueue.enqueue(4)
	myQueue.enqueue(5)
	myQueue.dequeue()
	fmt.Println(myQueue)
}
