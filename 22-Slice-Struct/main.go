package main

import (
	"fmt"
)

func main() {
	var players []Player
	players = []Player{Player{Id: 1, Name: "Farhan"},
		Player{Id: 2, Name: "Alifiannisa"},
		Player{Id: 3, Name: "Apid"},
		Player{Id: 4, Name: "Tacella"},
		Player{Id: 5, Name: "Yunita"}}

	players = append(players, Player{Id: 6, Name: "Malik"})
	players = append(players, Player{Id: 7, Name: "Nurul"})

	for _, v := range players {
		fmt.Println(v)
	}
}

type Player struct {
	Id   int
	Name string
}

// package main

// import (
// 	"fmt"
// )

// func main() {

// 	var players []Player
// 	players = []Player{Player{ID: 1, Name: "Maldini"}, Player{ID: 2, Name: "Nesta"},
// 		Player{ID: 3, Name: "Gattuso"}, Player{ID: 4, Name: "Pirlo"}}

// 	players = append(players, Player{ID: 5, Name: "Kaka"})
// 	players = append(players, Player{ID: 6, Name: "Pipo"})
// 	players = append(players, Player{ID: 7, Name: "Ronaldinho"})
// 	players = append(players, Player{ID: 8, Name: "Seedorf"})

// 	for _, v := range players {
// 		fmt.Println(v)
// 	}
// }

// type Player struct {
// 	ID   int
// 	Name string
// }
