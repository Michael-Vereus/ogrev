package main

import "fmt"

type orang struct {
	name string
	age  int
}

func main() {
	mike := orang{}
	mike.name = "Michael"
	mike.age = 20
	fmt.Println(mike)
	ptr := &mike
	*&ptr.name = "Mike"
	fmt.Println(mike)

	x := 10

	ptrx := &x // ngambil mem addr buat jdi pointr

	fmt.Println("Mem Address ", ptrx) // ptr = print mem addr raw
	*ptrx = 100                       // modify nilai asli via pointr
	fmt.Println("Num ", x)
}
