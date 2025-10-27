package main

import "fmt"

func main() {
	var awal, akhir int
	fmt.Print("masukan dua angka: ")
	fmt.Scan(&awal, &akhir)

	total := 1

	for i := awal; i <= akhir; i++ {
		total *= i
	}

	fmt.Printf("jumlah bakteri hari terakhir: %d\n", total)
}
