package arr

import "fmt"

var ArrOfNum = [...]int{0, 3, 1, 2, 5, 6, 8, 9, 10, 66}
var twoDArr [3][5]int

func LoopOverArray() {
	for i := range ArrOfNum {
		fmt.Println(ArrOfNum[i])
	}
}
func CheckOutTwoDArray() {
	for i := range len(twoDArr) {
		fmt.Print("| ")
		for j := range len(twoDArr[i]) {
			fmt.Printf("[%d][%d] | ", i, j)
		}
		fmt.Println("")
	}
}
func GetArrLen() int {
	return len(ArrOfNum)
}
