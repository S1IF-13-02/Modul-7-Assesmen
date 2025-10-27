package main

import "fmt"

func main() {
	var totalKeping int 

	var kepingPerIkat int = 8
	var ikatPerKarung int = 10
	var karungPerPeti int = 10
	
	fmt.Print("Masukkan total uang dalam satuan keping: ")
	fmt.Scan(&totalKeping)

	var kepingPerKarung int = ikatPerKarung * kepingPerIkat 
	var kepingPerPeti int = karungPerPeti * kepingPerKarung 
	var peti, karung, ikat, sisaKeping, finalKeping int

	peti = totalKeping / kepingPerPeti
	sisaKeping = totalKeping % kepingPerPeti 

	karung = sisaKeping / kepingPerKarung
	sisaKeping = sisaKeping % kepingPerKarung 

	ikat = sisaKeping / kepingPerIkat
	sisaKeping = sisaKeping % kepingPerIkat 
	
	finalKeping = sisaKeping
	
	fmt.Printf("\nKeluaran: %d peti, %d karung, %d ikat, dan %d keping\n", 
		peti, karung, ikat, finalKeping)
}