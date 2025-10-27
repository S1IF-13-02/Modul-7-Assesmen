package main

import "fmt"

func main() {
	var n, i int
	var genap int

	fmt.Print("Input: ")
	fmt.Scan(&n)

	for i = 1; i <= n; i += 1 {
		genap = i * 2
		fmt.Println(genap, " ")
	}
}
