package main

import "fmt"

func main() {
	var n int

	fmt.Print("Masukkan n: ")
	fmt.Scan(&n)

	for i := 1; i <= (n * 2); i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
