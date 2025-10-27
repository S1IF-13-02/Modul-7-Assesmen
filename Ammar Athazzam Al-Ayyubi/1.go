package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("jumlah bilangan genap: ")
	fmt.Scan(&n)

	fmt.Println("bilangan genap")
	for i := 1; i <= n; i++ {
		fmt.Printf("%d ", i*2)
	}

}
