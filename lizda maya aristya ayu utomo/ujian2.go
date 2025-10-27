package main

import "fmt"

func main() {
	var x, y int
	var hasil int

	fmt.Print(" masukan bilagan : ")
	fmt.Scan(&x, &y)

	hasil = 1
	for i := x; i <= y; i++ {
		hasil *= i
	}

	fmt.Println(" keluaran : ", hasil)

}
