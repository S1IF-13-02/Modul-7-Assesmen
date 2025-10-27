package main

import "fmt"

func main() {
	var n int
	fmt.Print("masukan bilangan habis dibagi dua: ")
	fmt.Scan(&n)

	fmt.Print("Output: ")
	for i := 1; i <= n; i++ {
		fmt.Print(i * 2)
		if i < n {
			fmt.Printf(", ")
		}
	}
	fmt.Println()
}
