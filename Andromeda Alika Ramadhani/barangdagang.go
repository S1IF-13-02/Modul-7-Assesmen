package main

import "fmt"

func main() {
	var peti, karung, ikat, keping, sisa int

	fmt.Print("Uang Keping: ")
	fmt.Scan(&keping)

	peti = keping / 800
	sisa = keping % 800
	karung = sisa / 80
	sisa = karung % 80
	ikat = sisa / 8
	keping = ikat % 8

	fmt.Printf("%d peti, %d karung, %d ikat, dan %d keping", peti, karung, ikat, keping)
}
