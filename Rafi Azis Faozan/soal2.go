package main

import "fmt"

func main() {
	var x int
	var y int
	var Bakteri int
	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan nilai y: ")
	fmt.Scan(&y)
	for i := x; i <= y; i++ {
		Bakteri = (x + i) * y
	}
	fmt.Print("Jumlah bakteri: ", Bakteri)
}
