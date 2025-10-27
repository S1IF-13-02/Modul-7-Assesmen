package main

import "fmt"

func main() {
	var totalUang int
	var peti, karung, ikat, keping int

	fmt.Print("Masukkan total uang dalam satuan keping: ")
	fmt.Scan(&totalUang)

	peti = totalUang / 800

	karung = (totalUang % 800) / 80

	ikat = (totalUang % 80) / 8

	keping = totalUang % 8

	fmt.Println(peti, "peti", karung, "karung", ikat, "ikat", keping, "keping")
}
