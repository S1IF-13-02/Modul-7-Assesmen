package main

import "fmt"

func main() {

	var (
		barang int
	)

	fmt.Printf("Masukan jumlah barang :")
	fmt.Scan(&barang)

	peti := barang / barang
	barang = barang % 10
	karung := barang / 10
	barang = barang % 10
	ikat := barang / 8
	barang = barang % 8
	keping := barang / 8
	barang = barang % 8

	fmt.Println(peti, karung, ikat, keping)
}
