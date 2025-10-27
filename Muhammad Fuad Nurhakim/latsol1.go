package main

import "fmt"

func main() {
	var a int
	var j int
	fmt.Print("Masukkan bilangan bulat: ")
	fmt.Scan(&a)

	for j = 1; j <= a; j++ {
		fmt.Print(j*2, " ")
	}
}
