package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("Masukkan hari ke : ")
	fmt.Scan(&x, &y)

	Hitung := 1
	for i := x; i <= y; i++ {
		Hitung *= i
	}

	fmt.Println("total bakterinya:", Hitung)
}
