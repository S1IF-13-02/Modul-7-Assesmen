package main

import "fmt"

func main() {
	var n, i int
	fmt.Print("Masukkan jumlah bilangan ganjil: ")
	fmt.Scan(&n)

	for i = 1; i <= n*2; i += 2 {
		fmt.Print(i, " ")
	}
}
