package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	jumlahBakteri := 1
	for i := x; i <= y; i++ {
		jumlahBakteri *= i
	}

	fmt.Println(jumlahBakteri)
}