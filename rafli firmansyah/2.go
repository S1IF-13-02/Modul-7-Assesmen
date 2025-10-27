package main

import "fmt"

func main() {
	var x, y, bakteri int
	fmt.Print("masukan dua bilangan: ")
	fmt.Scan(&x, &y)

	for i := 1; i <= y; i++ {
		bakteri += x
	}
	fmt.Println("hasil perkalian", bakteri)
}
