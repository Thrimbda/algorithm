package linkedstack

import (
	"github.com/Thrimbda/dune/datastructure/list/linkedlist"
)

type LinkedStack struct {
	list *linkedlist.LinkedList
}

func NewLinkedStack() *LinkedStack {
	return &LinkedStack{linkedlist.NewLinkedList()}
}

func ConverToLinkedStack(items ...interface{}) *LinkedStack {
	return &LinkedStack{linkedlist.ConvertToLinkedList(items)}
}

func (l *LinkedStack) Clear() {
	l.list.Clear()
}

func (l *LinkedStack) Push(item interface{}) {
	l.list.Insert(0, item)
}

func (l *LinkedStack) Pop() interface{} {
	return l.list.Remove(0)
}

func (l *LinkedStack) Peek() interface{} {
	return l.list.Get(0)
}

func (l *LinkedStack) IsEmpty() bool {
	return l.list.IsEmpty()
}
