package funct

import "fmt"

func ReturnFunct() func() {
	return func() {
		fmt.Println("this is called returning a function")
	}
}
