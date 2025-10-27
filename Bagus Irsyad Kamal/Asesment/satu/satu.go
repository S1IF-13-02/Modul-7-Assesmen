package main

import "fmt"

func main(){
	var ganjil int
	fmt.Print("Input berapa jumlah bilangan ganjil: ")
	fmt.Scan(&ganjil)

	fmt.Print("Output: ")
	for i := 1; i <= ganjil*2; i += 2 {
		fmt.Print(i, " ")
	}

}