package main

import "fmt"

func main() {
	var n int
	fmt.Print("masukkan: ")
	fmt.Scan(&n)

	fmt.Print("keluaran: ")
	for i := 1; i <= n; i++ {
		fmt.Printf("%d ", 2*i)
	}
}
