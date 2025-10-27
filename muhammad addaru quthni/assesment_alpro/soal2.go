package main

import "fmt"

func main() {
	var x, y, hasil int
	fmt.Print("input: ")
	fmt.Scan(&x, &y)
	
	hasil = 1
	for i := x; i <= y; i++ {
		hasil *= i
	}
	fmt.Println(hasil)
}