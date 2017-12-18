package queue

import (
	. "../../datastructure"
	. "../arrayutils"
)

type ArrayQueue struct {
	size int
	front int
	rear int
	listArray []Elem
}

func (a ArrayQueue) setup(size int) {
	a.size = size
	a.front = 0
	a.rear = 0
	a.listArray = make([]Elem, size)
}

func (a ArrayQueue) clear() {
	a.front = 0
	a.rear = 0
}

func (a ArrayQueue) enqueue(item Elem) error {
	if a.rear + 1 % a.size == a.front {
		return FullListError{}
	}
	a.rear = (a.rear + 1) % a.size
	a.listArray[a.rear] = item
	return nil
}

func (a ArrayQueue) dequeue() (Elem, error) {
	if a.isEmpty() {
		return nil, EmptyListError{}
	}
	a.front = (a.front + 1) % a.size
	return a.listArray[a.front], nil
}

func (a ArrayQueue) firstValue() (Elem, error) {
	if a.isEmpty() {
		return nil, EmptyListError{}
	}
	return a.listArray[(a.front + 1) % a.size], nil
}

func (a ArrayQueue) isEmpty() bool {
	return a.front == a.rear
}