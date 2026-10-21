package main

import (
	"fmt"
	"ogrev/funct"
	ifelse "ogrev/if-else"
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
	ifelse.IfElseNormal(6)
	ifelse.SwitchCaseDemo()
	hey := funct.ReturnFunct()
	hey()
}
