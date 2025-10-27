package main

import "fmt"

func main() {
	var peti, karung, ikat, keping int

	fmt.Print("Masukkan uang keping = ")
	fmt.Scan(&keping)

	peti = keping / (10*10*8)
	sisa := keping % (10*10*8)

	karung = sisa / (10*8)
	sisa = sisa % (10*8)

	ikat = sisa / 8
	keping = sisa % 8

	fmt.Println(peti, "peti", karung, "karung", ikat, "ikat", "dan", keping, "keping")

}