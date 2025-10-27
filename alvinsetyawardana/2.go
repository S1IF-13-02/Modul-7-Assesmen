package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("Masukkan hari awal: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan hari akhir: ")
	fmt.Scan(&y)

	if x > y || x <= 0 {
		fmt.Println("Input tidak valid. Pastikan x ≤ y dan x > 0.")
		return
	}

	jumlah := 1
	for i := x; i <= y; i++ {
		jumlah *= i
	}

	fmt.Printf("Jumlah bakteri pada hari ke-%d sampai ke-%d adalah: %d\n", x, y, jumlah)
}
