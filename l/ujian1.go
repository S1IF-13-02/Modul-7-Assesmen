package main

import "fmt"

func main() {
	var n int

	fmt.Print("masukkan bilangan: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i = i + 2 {
		fmt.Print(i, " ")
	}
}
