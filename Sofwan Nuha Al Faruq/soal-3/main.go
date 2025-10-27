package main

import "fmt"

func main() {
	var keping int
	fmt.Print("Masukkan jumlah keping: ")
	fmt.Scan(&keping)

	const kepingPerIkat = 8
	const kepingPerKarung = 10 * kepingPerIkat  
	const kepingPerPeti = 10 * kepingPerKarung  

	peti := keping / kepingPerPeti
	sisa := keping % kepingPerPeti

	karung := sisa / kepingPerKarung
	sisa = sisa % kepingPerKarung

	ikat := sisa / kepingPerIkat
	sisa = sisa % kepingPerIkat

	fmt.Printf("%d peti, %d karung, %d ikat, dan %d keping\n", peti, karung, ikat, sisa)
}
