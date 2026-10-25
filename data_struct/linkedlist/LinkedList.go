package linkedlist

import (
	"fmt"
)

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
func (LinkedList *LinkedList) FindByName(name string) {
	startLi := LinkedList.head
	for startLi != nil {
		if startLi.name == name {
			fmt.Println("Found It", name)
			return
		}
		startLi = startLi.next // move to next gaada ini basically for i tanpa i++
	}
	fmt.Println("Not Found ", name)
}

func (LinkedList *LinkedList) DeleteByName(name string) {
	startLi := LinkedList.head
	var prevLi *li = nil
	for startLi != nil {
		if startLi.name == name {
			// logika e kalo org mau nge hapus data yg posisi e di head otomatis kan pointer prevLi masi kosong alias nil
			// jadi ya tinggal tambahin checking if prevli == nil
			if prevLi == nil {
				// basically geser pointer head di linked list ke pointer yg di next nya
				LinkedList.head = LinkedList.head.next
				fmt.Println("Deleted ", name)
				return
			}
			// kalo ini buat handle kasus dimana user hapus di tail alias di akhir
			// basically ubah pointer tail jadi ke sebelumnya dan set next nya = nil
			if startLi.next == nil {
				LinkedList.tail = prevLi
				LinkedList.tail.next = nil
				fmt.Println("Deleted ", name)
				return
			}
			//
			prevLi.next = startLi.next
			fmt.Println("Deleted ", name)
			return
		}
		prevLi = startLi
		startLi = startLi.next
	}
}

func (LinkedList *LinkedList) PrintOnConsole() {
	start := LinkedList.head
	for start != nil {
		fmt.Println(start.name)
		start = start.next
	}
}
