package main

import "fmt"

func main() {
	var keping int
	fmt.Print("Masukkan jumlah keping: ")
	fmt.Scan(&keping)

	peti := keping / (8 * 10 * 10)
	keping %= 8 * 10 * 10

	karung := keping / (10 * 10)
	keping %= 10 * 10

	ikat := keping / 10
	keping %= 10

	fmt.Printf("%d peti, %d karung, %d ikat, dan %d keping\n", peti, karung, ikat, keping)
}
