package main

import "fmt"
func main (){
var  keping int
	fmt.Print("Masukan Keping : ")
	fmt.Scan(&keping)

	peti := keping / 800
	keping = keping % 800

	karung := keping / 80
	keping = keping % 80

	ikat := keping / 8
	keping = keping % 8

	fmt.Printf("%d peti, %d karung, %d ikat, dan %d keping\n",peti, karung, ikat, keping)
}