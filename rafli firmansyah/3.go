package main

import "fmt"

func main() {
	var peti, karung, ikat, keping int
	fmt.Print("masukan keping:")
	fmt.Scan(&keping)

	peti = keping / 800
	karung = peti % 10
	ikat = karung % 10
	keping = peti % 100

	fmt.Println("%d peti, %d karung, %d ikat, %d keping", peti, karung, ikat, keping)
}
