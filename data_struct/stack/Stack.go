package stack

import "ogrev/data_struct/linkedlist"

type Stack struct {
	LinkedList *linkedlist.LinkedList
}

func NewStack() *Stack {
	return &Stack{LinkedList: linkedlist.NewLinkedList()}
}

func (Stack *Stack) Add(name string) {
	Stack.LinkedList.Add(name)
}
func (Stack *Stack) Remove() {
	Stack.LinkedList.MovwHeadToNext()
}

func (Stack *Stack) Peek() {
	Stack.LinkedList.PrintOnConsole()
}
