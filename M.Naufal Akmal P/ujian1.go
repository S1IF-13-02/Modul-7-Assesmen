package main

import "fmt"

func main () {
	var n int
	fmt.Print("masukan bilangan bulat n: ")
	fmt.Scan(&n)
	fmt.Println("n harus bilangan bulat positif.")
	

	for i := 1; i <= n; i++ {
		bilanganGenap := 2 * i
		fmt.Print(bilanganGenap)

		if i < n {
			fmt.Print(", ")
		}
	}
	fmt.Println()
}
