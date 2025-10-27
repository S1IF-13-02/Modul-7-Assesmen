package main

import "fmt"

func main() {
	var n int

	fmt.Print("Masukkan bilangan n: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i += 2 {
		fmt.Print(i, " ")
	}
}
