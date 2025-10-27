package main

import "fmt"

func main() {
	var n int
	fmt.Print("masukkan bilangan: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("input harus bilangan positif")
		return
	}

	fmt.Printf("deret %d bilangan ganjil:\n", n)

	for i := 0; i < n; i++ {
		bilanganGanjil := 2*i + 1
		fmt.Print(bilanganGanjil)

		if i < n-1 {
			fmt.Print(", ")
		}
	}
}
