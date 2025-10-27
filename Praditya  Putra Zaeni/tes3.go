package main

import "fmt"

func main () {

 	var keping int
	fmt.Scan(&keping)

	peti := keping / 800
	sisa := keping % 800

 	karung :=  sisa / 80
	sisa = sisa % 80

	ikat := sisa / 8
	sisa = sisa % 8

 	fmt.Printf("%d peti, %dkeping, %dikat dan %dkarung\n", peti, karung, ikat, sisa)
}
