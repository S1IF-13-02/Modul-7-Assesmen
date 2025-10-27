package main

import "fmt"

func main() {

	var n int
	var jumlah int

	fmt.Print("Masukkan bilangan negatif (n): ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {

		jumlah = jumlah - i

	}

	fmt.Printf(" ", n, jumlah)
}
