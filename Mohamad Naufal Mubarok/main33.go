package main

import "fmt"

func main() {
    var keping int
    fmt.Print("Masukkan Nilai keping: ")
    fmt.Scan(&keping)


	//satuPeti:= 10(keping)
	//satuKarung:= 10(ikat)
	//satuIkat:= 8(keping)

	//semua di * 10*10*8= 800
	//lalu jumlah / %

    peti := keping / 800
    sisa := keping % 800

    karung := sisa / 80
    sisa = sisa % 80

    ikat := sisa / 8
    sisa = sisa % 8

    fmt.Printf("%d peti, %d karung, %d ikat, dan %d keping\n", peti, karung, ikat, sisa)
}