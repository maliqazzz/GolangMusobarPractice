package main

import (
	"fmt"
)

func main() {
	var arr [5]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5
	// looping dengan array
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	// memanggil array
	fmt.Println(arr[2])

	//penulisan array dengan cara lain
	arrStr := [5]string{"Alex", "Ben", "Den", "Sarah", "Mury"}
	fmt.Println(arrStr)
	fmt.Println(arrStr[3])

	arrMulti := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arrMulti[1][2])
}
