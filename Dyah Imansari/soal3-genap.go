package main
import "fmt"
func main() {
	var Peti, Karung, Ikat, Keping int

	fmt.Scan(&Keping)
	Peti = Keping / 800
	Karung = (Keping % 800) / 80
	Ikat = (Keping % 800 % 80) / 8
	Keping = Keping % 800 % 80 % 8

	fmt.Println(Peti, "peti,", Karung, "karung,", Ikat, "ikat,", Keping, "dan keping")
}