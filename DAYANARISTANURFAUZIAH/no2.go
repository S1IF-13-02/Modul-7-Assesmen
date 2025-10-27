package main

import "fmt"

func main() {
	var x, y, hasil int
	fmt.Scan(&x, &y)

	for i := 1; i <= y; i++ {

		hasil = x * y 
	}

	fmt.Print("x", x,"y" y,"hasil" hasil)
}
