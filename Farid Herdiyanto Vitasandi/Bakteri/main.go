package main

import "fmt"

func main() {
	var hariAwal, hariAkhir int

	fmt.Print("Masukkan hari awal (x): ")
	fmt.Scan(&hariAwal)

	fmt.Print("Masukkan hari akhir (y): ")
	fmt.Scan(&hariAkhir)

	awalBakteri := 1

	for i := hariAwal; i <= hariAkhir; i++ {
		awalBakteri = awalBakteri * i
		fmt.Print(awalBakteri, " ")
	}
}

