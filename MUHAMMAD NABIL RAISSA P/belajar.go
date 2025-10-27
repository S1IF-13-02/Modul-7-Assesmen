package main

import "fmt"

func main() {
	var n int
	fmt.Print("masukan: ")
	fmt.Scan(&n)

	fmt.Print("hasil: ")
	for i := 1; i <= n*2; i += 2 {
		fmt.Print(i, " ")
	}
}
