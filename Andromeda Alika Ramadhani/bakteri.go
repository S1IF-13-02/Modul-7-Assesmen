package main

import "fmt"

func main() {
	var x, y, i, bakteri int

	fmt.Print("Hari ke: ")
	fmt.Scan(&x, &y)

	bakteri = 1
	for i = x; i <= y; i += 1 {
		bakteri = bakteri * i
	}
	fmt.Println(bakteri)
}
