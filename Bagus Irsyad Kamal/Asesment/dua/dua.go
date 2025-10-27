package main

import "fmt"

func main() {
	var awal, akhir int
	fmt.Print("Masukkan hari awal dan hari akhir : ")
	fmt.Scan(&awal, &akhir)

	hasil := 1
	for i := awal; i <= akhir; i++ {
		hasil *= i
	}

	fmt.Println("Jumlah bakteri adalah:", hasil)
}
