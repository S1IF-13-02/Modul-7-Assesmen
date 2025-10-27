package main

import "fmt"

func main() {
	var peti, karung, ikat, keping int

	fmt.Print("Masukkan jumlah dalam satuan keping: ")
	fmt.Scan(&keping)

	peti = keping / 800
	karung = (keping % 800) / 80
	ikat = (keping % 80) / 8
	keping = keping % 8

	fmt.Printf("Setara dengan %d peti, %d karung, %d ikat, dan %d keping.\n",
		peti, karung, ikat, keping)
}
