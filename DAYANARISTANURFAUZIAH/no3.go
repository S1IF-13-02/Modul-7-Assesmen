package main

import "fmt"

func main() {
	var peti, karung, ikat int
	fmt.Print("Masukkan bilangan bulat: ")
	fmt.Scan(&peti)
	karung = ikat % 10
	ikat = peti % 8
	peti = karung % 8
	fmt.Printf("%d peti, %d karung, %d ikat\n", peti, karung, ikat)
}
