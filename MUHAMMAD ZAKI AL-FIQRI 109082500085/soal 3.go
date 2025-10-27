package main

import "fmt"

func main() {
	var peti, karung, ikat, keping, totalkeping, sisa int
	fmt.Print("Masukkan jumlah keping: ")
	fmt.Scan(&totalkeping)

	peti = totalkeping / 800
	sisa = totalkeping % 800

	karung = sisa / 100
	sisa = sisa % 100

	ikat = sisa / 10
	keping = sisa % 10

	fmt.Println("peti   :", peti)
	fmt.Println("karung :", karung)
	fmt.Println("ikat   :", ikat)
	fmt.Println("keping :", keping)
}
