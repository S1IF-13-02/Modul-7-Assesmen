package main

import "fmt"

func main() {
	var n, jumlah int

	fmt.Print("Masukkan bilangan = ")
	fmt.Scan(&n)


	for i := 1; i <= n; i++{
		jumlah += 2
		fmt.Print(jumlah, " ")
	}
}