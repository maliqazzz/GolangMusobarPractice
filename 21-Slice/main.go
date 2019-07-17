package main

import (
	"fmt"
)

func main() {
	mySlice := []int{10, 20, 30, 40, 50}

	for i, v := range mySlice {
		fmt.Println(i, v)
	}

	mySliceStr := []string{"Alex", "Ben", "Den", "Wury"}
	mySliceStr = append(mySliceStr, "Zen") //menambahkan value
	//semula adaa 4 value sekarang jadi 5
	for _, v := range mySliceStr {
		fmt.Println(v)
	}
}
