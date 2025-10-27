// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var input int
// 	fmt.Printf("input : ")
// 	fmt.Scan(&input)

// 	n := 2

// 	for i := 1; i <= input; i++ {
// 		fmt.Printf(i * n)
// 	}
// }

package main

import "fmt"

func main() {
	var input int
	fmt.Print("input : ")
	fmt.Scan(&input)

	n := 2

	for i := 1; i <= input; i++ {
		fmt.Println(i * n)
	}
}
