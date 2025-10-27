package main

import "fmt"

func main() {
	var a, b, jumlahbakteri, j int

	fmt.Print("masukan jumlah bakteri awal : ")
	fmt.Scan(&a)
	fmt.Print("masukan jumlah bakteri terakhir : ")
	fmt.Scan(&b)
	jumlahbakteri = 1
	for j = a; j <= b; j++ {
		jumlahbakteri = jumlahbakteri * j

	}
	fmt.Print("total bakteri : ", jumlahbakteri)

}
