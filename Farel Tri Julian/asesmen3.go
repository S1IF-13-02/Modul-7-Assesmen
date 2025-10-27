package main

import "fmt"

func main() {
	var totalkeping int

	fmt.Print("masukan keping: ")
	fmt.Scan(&totalkeping)

	peti := totalkeping / 800
	sisa := totalkeping % 800

	karung := sisa / 100
	sisa = sisa % 100

	ikat := sisa / 10
	keping := sisa % 10

	fmt.Printf("%d peti, %d karung, %d ikat, %d keping\n", peti, karung, ikat, keping)
}
