package main

import "fmt"

func main(){
	var keping int
	fmt.Print("keping: ")

	fmt.Scan(&keping)
	peti := keping / 800
	sisa := keping % 800
	karung := sisa / 80
	sisa = sisa % 80
	ikat := sisa / 10
	keping = sisa % 8
	fmt.Printf("%d peti, %d karung, %d ikat, %d keping", peti, karung, ikat, keping)

}