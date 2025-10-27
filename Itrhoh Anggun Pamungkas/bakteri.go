package main

import "fmt"

func main() {
	var x, y int

	fmt.Print("masukkan nilai x dan y: ")
	fmt.Scan(&x, &y)

	jumlah := 1
	for i := x; i <= y; i++ {
		jumlah *= i
	}
	fmt.Print(jumlah)
}
