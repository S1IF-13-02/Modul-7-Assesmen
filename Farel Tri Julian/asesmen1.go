package main

import "fmt"

func main() {
	var a int
	fmt.Print("masukan angka: ")
	fmt.Scan(&a)

	fmt.Print("Output: ")
	for i := 1; i <= a*2; i += 2 {
		fmt.Print(i, " ")
	}
}
