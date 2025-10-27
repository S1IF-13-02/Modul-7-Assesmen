package main

import "fmt"

func main() {
	var n int
	var i int
	var bilanganGenap int

	fmt.Print("Masukkan jumlah bilangan genap yang ingin ditampilkan: ")
	fmt.Scan(&n)

	fmt.Print("Output: ")
	for i = 1; i <= n; i++ {
		bilanganGenap = i * 2
		fmt.Print(bilanganGenap)
		if i < n {
			fmt.Print(", ")
		}
	}
	fmt.Println()
}
