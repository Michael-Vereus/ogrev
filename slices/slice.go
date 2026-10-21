package slices

import "fmt"

var disSlice []int = []int{9, 8, 7, 10}

func AddToSlice(i int) {
	disSlice = append(disSlice, i)
}
func PrintSlice() {
	for i := range disSlice {
		fmt.Printf("|[%d]|", disSlice[i])
	}
	fmt.Println()
}
