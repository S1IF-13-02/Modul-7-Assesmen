package main

import "fmt"

func main() {
	var r, p int
	fmt.Print("Masukkan hari awal dan hari akhir : ")
	fmt.Scan(&r, &p)

	hasil := 1
	for i := r; i <= p; i++ {
		hasil *= i
	}

	fmt.Println("Output:", hasil)
}