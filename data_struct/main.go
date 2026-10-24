package main

import (
	"fmt"
	"ogrev/data_struct/linkedlist"
)

type element struct {
	name string
	next *element
}

func newElement(newName string) *element {
	el := element{name: newName, next: nil}
	return &el
}
func showAllNames(headEl *element) {
	for {
		fmt.Println(headEl.name)
		if headEl.next == nil {
			break
		}
		headEl = headEl.next
	}
}

func addToList(elPoint *element, newElpoint *element) {
	for elPoint.next != nil {
		elPoint = elPoint.next
	}
	elPoint.next = newElpoint
}

func searchByName(head *element, name string) {
	for head != nil {
		if head.name == name {
			fmt.Printf("Found Name : {%s} \n", name)
			return
		}
		head = head.next
	}
	fmt.Printf("Not Found Name : {%s} \n", name)
}

func deleteByName(head *element, name string) *element {
	originalHead := head
	var oldPoint *element = nil
	for head != nil {
		if head.name == name {
			if oldPoint != nil {
				oldPoint.next = head.next
				if oldPoint.next == nil {
					fmt.Println("Deleted at the end")
					break
				}
				fmt.Println("Deleted at the middle")
				break
			}
			head = head.next
			fmt.Println("Deleted at the start")
			return head // where tf is free memory func ???
		}
		oldPoint = head
		head = head.next
	}
	return originalHead
}

func main() {
	// fmt.Println("ran check")
	// elpoin := newElement("ver")
	// fmt.Println(elpoin.name)
	// addToList(elpoin, newElement("mike"))
	// addToList(elpoin, newElement("lim"))
	// addToList(elpoin, newElement("tara"))
	// showAllNames(elpoin)
	// searchByName(elpoin, "ver")
	// searchByName(elpoin, "vereus")
	// elpoin = deleteByName(elpoin, "ver")
	// elpoin = deleteByName(elpoin, "tara")
	// showAllNames(elpoin)
	listOfNames := linkedlist.NewLinkedList()
	listOfNames.Add("Ver")
	listOfNames.Add("Mika")
	listOfNames.Add("Lim")
	listOfNames.FindByName("Ver")
	listOfNames.FindByName("Tara")
	listOfNames.PrintOnConsole()
}
