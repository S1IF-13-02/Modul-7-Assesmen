package main

import "fmt"

func main() {
	var n int

	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		bilanganGanjil := (2 * i) + 1

		fmt.Print(bilanganGanjil, " ")
	}
	fmt.Println()
}
