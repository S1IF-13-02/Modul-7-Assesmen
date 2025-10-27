package main

import "fmt"

func main() {
	var totalKeping int

	keping per ikat = 10
	keping per karung = 100
	keping per peti = 800

	fmt.Print("Masukkan total Keping: ")
	fmt.Scanln(&totalKeping)

	peti := totalKeping / keping per peti
	sisaSetelahPeti := totalKeping % keping per peti

	karung := sisaSetelahPeti / keping per karung
	sisaSetelahKarung := sisaSetelahPeti % keping per karung

	
