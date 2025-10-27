package main

import "fmt"

func main () {

 	var  x, y int
	fmt.Scan(&x, &y)
	total := 1

 	for i := x; i <= y; i++ {
		total *= i
 
 	fmt.Println(total)
}

}
