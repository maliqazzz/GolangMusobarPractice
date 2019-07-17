// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	go helloGo()

// 	time.Sleep(3 * time.Second)
// 	fmt.Println("Ini fungsi main")
// }
// func helloGo() {
// 	fmt.Println("Hello Go")
// }

package main

import (
	"fmt"
)

func main() {

	done := make(chan bool)

	go helloGo(done)

	<-done

	fmt.Println("Ini fungsi main")
}
func helloGo(done chan bool) {
	fmt.Println("Hello Go")
	done <- true
}
