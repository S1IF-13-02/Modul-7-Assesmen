package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("hari pertama : ")
	fmt.Scan(&a)

	fmt.Print("hari terakhir : ")
	fmt.Scan(&b)

	hasil := 1
	for i := a; i <= b; i++ {
		hasil *= i
	}

	fmt.Println("jumlah :", hasil)
}
