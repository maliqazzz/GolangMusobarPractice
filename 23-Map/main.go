package main

import (
	"fmt"
)

func main() {
	var mapPerson map[int]string
	mapPerson = make(map[int]string)

	mapPerson[1] = "Alex"
	mapPerson[2] = "Nurdin"
	mapPerson[3] = "Tejo"
	mapPerson[4] = "Malik"

	for k, v := range mapPerson {
		fmt.Println(k, v)
	}

	malik, ok := mapPerson[4]
	if !ok {
		fmt.Println("Malik Tidak ada di mapPerson")
	} else {
		fmt.Println(malik)
	}
}
