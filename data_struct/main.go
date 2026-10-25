package main

import (
	"ogrev/data_struct/linkedlist"
	"ogrev/data_struct/stack"
)

func main() {
	listOfNames := linkedlist.NewLinkedList()
	listOfNames.Add("Ver")
	listOfNames.Add("Mika")
	listOfNames.Add("Lim")
	listOfNames.FindByName("Ver")
	listOfNames.FindByName("Tara")
	listOfNames.PrintOnConsole()
	listOfNames.DeleteByName("Ver")
	listOfNames.PrintOnConsole()
	stacks := stack.NewStack()
	stacks.Add("Stack 1")
	stacks.Add("Stack 2")
	stacks.Add("3")
	stacks.Peek()
	stacks.Remove()
	stacks.Peek()
}
