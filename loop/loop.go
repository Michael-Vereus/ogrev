package loop

import (
	"fmt"
)

func LoopClassic() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

func LoopUsingRange(n int) {
	for i := range n {
		fmt.Println("In Range : ", i)
	}
}

func LoopWhile(bool bool) {
	for {
		if !bool {
			fmt.Println("true")
			break
		} else {
			fmt.Println("false")
			bool = false
		}
	}
}
