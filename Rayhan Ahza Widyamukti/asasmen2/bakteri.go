package main

import "fmt"

func main() {
	var x, y int
	var i int
	var jumlahBakteri int

	fmt.Print("Masukkan hari awal (x): ")
	fmt.Scan(&x)
	fmt.Print("Masukkan hari akhir (y): ")
	fmt.Scan(&y)
	
	jumlahBakteri = x 

	for i = x + 1; i <= y; i++ {
		jumlahBakteri = jumlahBakteri * i
	}

	fmt.Println("Jumlah bakteri pada hari ke-", y, "adalah:", jumlahBakteri)
}
