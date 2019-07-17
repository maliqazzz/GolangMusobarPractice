package main

import (
	"fmt"
)

type Vector struct {
	X int
	Y int
}

type Player struct {
	Id   int
	Name string
}

func main() {

	var v Vector
	v.X = 10
	v.Y = 5
	fmt.Println(v)

	fmt.Println("X =", v.X)
	fmt.Println("Y =", v.Y)

	player1 := Player{Id: 01, Name: "Ronaldinho"}

	fmt.Println(player1.Id)
	fmt.Println(player1.Name)

}
