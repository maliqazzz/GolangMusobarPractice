// package main

// import (
// 	"fmt"
// )

// func main() {
// 	defer fmt.Println("Dieksekusi ke-tiga")

// 	defer fmt.Println("Dieksekusi ke-dua")

// 	fmt.Println("Dieksekusi pertama")
// }

package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 6; i++ {
		defer process(i)
	}
}
func process(id int) {
	fmt.Printf("Process %d\n", id)
}
