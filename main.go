package main

import (
	"fmt"
	"ogrev/loop"
)

func main() {
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
	f := "hello"
	fmt.Println(f)
	loop.LoopClassic()
	loop.LoopUsingRange(7)
	loop.LoopWhile(false)
}
