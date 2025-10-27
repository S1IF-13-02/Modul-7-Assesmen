package main 

import "fmt"

func main(){
	var x,y int
	var bakteri int = 1
	fmt.Scan(&x, &y)
	for i := x ; i <= y ; i++{
		bakteri = bakteri * i
	}
	fmt.Print(bakteri)
}