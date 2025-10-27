package main

import "fmt"

func main() {
	var x, y int

	fmt.Print("masukan x: ")
	fmt.Scan(&x)
	fmt.Print("masukan y: ")
	fmt.Scan(&y)

	hasil := 1
	for i := x; i <= y; i++ {
		hasil *= i
	}

	fmt.Println("Output: ", hasil)
}
