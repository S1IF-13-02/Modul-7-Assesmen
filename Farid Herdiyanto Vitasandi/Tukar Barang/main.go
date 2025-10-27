package main

import "fmt"

func main() {
	var jumlahKeping int 

	fmt.Print("Masukkan jumlah keping: ")
	fmt.Scan(&jumlahKeping)

	peti := jumlahKeping / 800
	sisaSetelahPeti := jumlahKeping % 800
	
	karung := sisaSetelahPeti / 100
	sisaSetelahKarung := sisaSetelahPeti % 100
	
	ikat := sisaSetelahKarung / 10
	keping := sisaSetelahKarung % 10

	fmt.Printf("%d peti, %d karung, %d ikat, dan %d keping", peti, karung, ikat, keping)
}
