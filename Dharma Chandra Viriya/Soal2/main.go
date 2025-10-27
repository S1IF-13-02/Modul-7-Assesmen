package main

import "fmt"

func main() {
	var x, n int
	var result int = 1

	fmt.Print("Masukkan angka x dan n: ")
	fmt.Scan(&x, &n)

	for i := x; i <= n; i++ {
		result *= i
	}

	fmt.Println(result)
}
