package main

import "fmt"

func main() {
	var keping, peti, sisa, karung, ikat int
	fmt.Print("Masukkan jumlah keping: ")
	fmt.Scan(&keping)

	peti = keping / 800
	sisa = keping % 800

	karung = sisa / 80
	sisa = sisa % 80

	ikat = sisa / 8
	sisa = sisa % 8

	fmt.Printf("%d peti, %d karung, %d ikat, %d keping\n", peti, karung, ikat, sisa)
}
