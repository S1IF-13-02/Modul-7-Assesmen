package main

import "fmt"

func main() {
	var n int
	fmt.Print("input: ")
	fmt.Scan(&n)
	
	fmt.Print("output: ")
	for i := 1; i < n; i++ {
		fmt.Print(i*2, ",")
	}
	fmt.Print(n * 2)
	fmt.Println()
}
