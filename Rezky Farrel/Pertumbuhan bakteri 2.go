package main
import "fmt"

func main() {
	var x, y int

	fmt.Scan(&x, &y)
	var totalBakteri int = 1

	for i := x; i <= y; i++ {
		totalBakteri *= i
	}
	fmt.Println(totalBakteri)
}
