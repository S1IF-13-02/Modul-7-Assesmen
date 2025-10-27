package main

import "fmt"

func main(){
	var n, hasil int

	fmt.Print("Masukan Bilangan: ")
	fmt.Scan(&n)

	for i := 0 ; i < n ; i++ {
		hasil = hasil + 2
		fmt.Print(hasil, " ")
	}

}