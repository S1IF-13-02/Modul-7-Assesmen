package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan jumlah bilangan genap: ")
	fmt.Scan(&n)

	fmt.Print("Hasil: ")
	for i := 1; i <= n; i++ {
		fmt.Print(2*i)
            fmt.Print(", ")

		}
	}
