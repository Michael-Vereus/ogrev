package linkedlist

import "fmt"

type li struct {
	name string
	next *li
}
type LinkedList struct {
	head *li
	tail *li
}

func NewLinkedList() *LinkedList {
	return &LinkedList{head: nil, tail: nil}
}

func newLi(name string) *li {
	return &li{name: name, next: nil}
}

func (LinkedList *LinkedList) Add(name string) {
	startLi := LinkedList.head
	newList := newLi(name)
	// ini buat safeguard kalo abis pertama buat LinkedList kan masih nil,
	// waktu di add biar engga error makanya tambahin check startLi aka head nya,
	// kemudian set dengan pointer new list kalo nil head nya
	if startLi == nil {
		LinkedList.head = newList
		LinkedList.tail = newList
		return // langsung keluar kalo block if udh sls
	}
	// tinggal nambah elemen / li di akhir dengan cara ubah pointer,
	// di tail.next biar ngarah ke mem address newList tsb,
	endLi := LinkedList.tail
	endLi.next = newList
	LinkedList.tail = newList // kemudian jangan lupa update LinkedList.tail menjadi mem address newList
	// kenapa? ya karena jelas" linkedlist mencatat catatan statis kea index
	// otomatis harus di update tiap kali ada operasi
}
func (LinkedList *LinkedList) PrintOnConsole() {
	start := LinkedList.head
	for start != nil {
		fmt.Println(start.name)
		start = start.next
	}
}
