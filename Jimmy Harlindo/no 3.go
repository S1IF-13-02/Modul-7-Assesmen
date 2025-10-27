package main
import "fmt"

func main() {
	var keping int
	fmt.Print("Masukkan jumlah keping: ")
	fmt.Scan(&keping)

	kepingPerIkat := 10
	kepingPerKarung := 10 * kepingPerIkat
	kepingPerPeti := 8 * kepingPerKarung

	peti := keping / kepingPerPeti
	sisa := keping % kepingPerPeti
	karung := sisa / kepingPerKarung
	sisa = sisa % kepingPerKarung
	ikat := sisa / kepingPerIkat
	kepingSisa := sisa % kepingPerIkat

	fmt.Printf("%d peti, %d karung, %d ikat, dan %d keping\n", peti, karung, ikat, kepingSisa)
}
