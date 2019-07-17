// package main

// import (
// 	"errors"
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	iStr := "b"
// 	//value iStr coba ganti menjadi huruf alfabet

// 	i, err := strconv.Atoi(iStr)

// 	if err != nil {
// 		fmt.Println("Terjadi error : ", err.Error())
// 	} else {
// 		fmt.Println(i)
// 	}

// 	r, err := Div(10, 0)
// 	//coba ubah pembagi menjadi 0

// 	if err != nil {
// 		fmt.Println("Terjadi error : ", err.Error())
// 	} else {
// 		fmt.Println(r)
// 	}
// }

// func Div(x, y int) (int, error) {
// 	if y == 0 {
// 		return 0, errors.New("Tidak bisa membagi dengan 0")
// 	}
// 	result := x / y
// 	return result, nil
// }

package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	iStr := "b"
	//value iStr coba ganti menjadi huruf alfabet

	i, err := strconv.Atoi(iStr)

	if err != nil {
		fmt.Println("Terjadi error : ", err.Error())
	} else {
		fmt.Println(i)
	}

	r, err := Div(10, 0)
	//coba ubah pembagi menjadi 0

	if err != nil {
		fmt.Println("Terjadi error : ", err.Error())
	} else {
		fmt.Println(r)
	}

	if x, err := Div(25, 0); err != nil {
		// cara inline handling error(sama seperti cara yang atas namun lebih simple)
		fmt.Println("Terjadi error : ", err.Error())
	} else {
		fmt.Println(x)
	}
}

func Div(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("Tidak bisa membagi dengan 0")
	}
	result := x / y
	return result, nil
}
