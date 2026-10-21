package main

import (
	"fmt"
	"ogrev/arr"
	"ogrev/funct"
	ifelse "ogrev/if-else"
	"ogrev/loop"
	"ogrev/slices"
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
	arr.ArrOfNum[1] = 999
	arr.LoopOverArray()
	fmt.Println(arr.GetArrLen())
	arr.CheckOutTwoDArray()
	slices.AddToSlice(99)
	slices.PrintSlice()
}
