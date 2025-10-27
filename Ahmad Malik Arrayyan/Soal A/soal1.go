package main

import "fmt"

func main() {
	var r int
	fmt.Print("Masukkan jumlah bilangan genap: ")
	fmt.Scan(&r)

	fmt.Print("Output: ")
	for i := 1; i <= r; i++ {
		fmt.Print(i*2, " ")
	}
}
