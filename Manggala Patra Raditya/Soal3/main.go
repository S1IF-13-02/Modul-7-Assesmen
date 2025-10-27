package main

import "fmt"

func main() {
	var keping int
	fmt.Scan(&keping)

	peti := keping / 800
	sisa := keping % 800

	karung := sisa / 100
	sisa %= 100

	ikat := sisa / 10
	keping = sisa % 10

	fmt.Println(peti, "peti,", karung, "karung,", ikat, "ikat,", "dan", keping, "keping")
}
