package main

import "fmt"

func main() {
	var x, y , i int
	fmt.Print("Masukkan hari awal: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan hari akhir: ")
	fmt.Scan(&y)
	hasil := 1
	for i = x; i <= y; i++ {
		hasil = hasil*i
	}

	fmt.Println("Jumlah bakteri pada hari ke-", y, "adalah:", hasil)
}
