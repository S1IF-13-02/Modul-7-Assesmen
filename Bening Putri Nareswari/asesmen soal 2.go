package main

import "fmt"

func main() {
	var hariAwal, hariAkhir int
	var jumlahBakteri int = 1

	fmt.Print("Masukkan hari pertama (x): ")
	fmt.Scanln(&hariAwal)

	fmt.Print("Masukkan batas hari (y): ")
	fmt.Scanln(&hariAkhir)

	faktorPerkembangan := hariAwal + 1

	for hari := hariAwal; hari <= hariAkhir; hari++ {

		{
			jumlahBakteri += faktorPerkembangan
		}

		fmt.Printf("Hari ke-%d: Jumlah Bakteri = %d\n", hari, jumlahBakteri)
	}
}
