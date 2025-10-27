package main

import "fmt"

func main() {
	var a int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&a)

	for i := 1; i <= a; i++ {
		fmt.Printf("%d ", i*2)
	}
}
