package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("Masukan x dan y :")
	fmt.Scan(&x, &y)

	hasil := x
	fmt.Print("Logika  = ", x)     
	for i := x + 1; i <= y; i++ {  
		fmt.Print("x", i)
		hasil *= i
	}
	fmt.Println()                 

	fmt.Println("Output  =", hasil) 
}
