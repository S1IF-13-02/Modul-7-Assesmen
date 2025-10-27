package main

import "fmt"

func main() {
	var b int = 5
	fmt.Print("masukan nilai a: ")
	fmt.Scan(&b)

	for i := 1; i <= b; i += 2 {
		fmt.Print(i)

	}

}
