package main

import "fmt"

func main() {
	var x, y int

	fmt.Print("Masukan Hari x : ")
	fmt.Scan(&x)
	fmt.Print("Masukan Hari y : ")
	fmt.Scan(&y)

	total := 1
	
	for i := x; i <= y; i++ {
		total *= i
	}
	fmt.Println("Jumalah bakteri adalah : ",total)
}