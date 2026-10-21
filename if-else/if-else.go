package ifelse

import (
	"fmt"
	"time"
)

func IfElseNormal(n int) {
	if n == 7 || n == 6 {
		fmt.Println("SIX SEVEN!")
	} else {
		fmt.Println("aww shucks :(")
	}
}

func SwitchCaseDemo() {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Its The Weekend")
	default:
		fmt.Println("Its Weekdays")
	}
}
