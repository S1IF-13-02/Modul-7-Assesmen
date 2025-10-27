package main

import "fmt"

func main() {
	var x, y, n int
	var hasil int
	fmt.Print("Masukkan bilangan bulat: ")
	fmt.Scan(&x, &y)

	for n = x; n <= y; n++ {
		hasil = y * (x + y)
	}
	fmt.Println(hasil)
}
