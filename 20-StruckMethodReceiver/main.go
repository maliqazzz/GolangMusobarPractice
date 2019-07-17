package main

import (
	"fmt"
)

func main() {

	p := &Person{
		Id:   1,
		Name: "Malik",
		Age:  22,
	}

	// fmt.Println(p.GetId())
	// fmt.Println(p.GetName())
	// fmt.Println(p.GetAge()

	fmt.Println(p.GetName())
	p.changeName("Habibie")
	fmt.Println(p.GetName())

}

type Person struct {
	Id   int
	Name string
	Age  int
}

func (p *Person) changeName(newName string) {
	p.Name = newName
}

func (p *Person) GetId() int {
	return p.Id
}

func (p *Person) GetName() string {
	return p.Name
}

func (p *Person) GetAge() int {
	return p.Age
}
