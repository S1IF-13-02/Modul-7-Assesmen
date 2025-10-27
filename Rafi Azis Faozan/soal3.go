package main

import "fmt"

func main() {
	var peti int
	var karung int
	var ikat int
	var keping int
	fmt.Print("Masukkan nilai keping: ")
	fmt.Scan(&keping)
	peti = keping / 800
	sisa := keping % 800
	karung = sisa / 80
	sisa = sisa % 80
	ikat = sisa / 8
	kepingSisa := sisa % 8
	fmt.Println("peti:", peti)
	fmt.Println("karung:", karung)
	fmt.Println("ikat:", ikat)
	fmt.Println("keping:", kepingSisa)
}
