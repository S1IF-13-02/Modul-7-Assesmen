package main

import "fmt"

func main() {
	var inputKeping int

	fmt.Print("Masukkan keping: ")
	fmt.Scan(&inputKeping)

	peti := (((inputKeping / 8) / 10) / 10) % 10
	karung := ((inputKeping / 8) / 10) % 10
	ikat := (inputKeping / 8) % 10
	keping := inputKeping % 8

	fmt.Printf("%d peti, %d karung, %d ikat, dan %d keping\n", peti, karung, ikat, keping)
}
