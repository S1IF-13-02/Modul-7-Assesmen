package main

import "fmt"

func main() {
	var x, y int
	
	fmt.Print("Masukkan hari awal (x) dan hari akhir (y): ")
	
	fmt.Scan(&x, &y)

	var jumlahBakteri int64 = 1 

	for i := x; i <= y; i++ {
		jumlahBakteri *= int64(i)
	}

	fmt.Printf("Jumlah bakteri terakhir (dari hari %d sampai %d) adalah: %d\n", x, y, jumlahBakteri)
}