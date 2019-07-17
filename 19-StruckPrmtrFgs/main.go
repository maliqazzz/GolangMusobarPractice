package main

import (
	"fmt"
)

func main() {
	p := Person{
		Id:   1,
		Name: "Malik",
		Age:  22,
	}
	printPerson(p)
}

type Person struct {
	Id   int
	Name string
	Age  int
}

func printPerson(p Person) {
	fmt.Println("ID =", p.Id)
	fmt.Println("Name =", p.Name)
	fmt.Println("Age =", p.Age)
}
