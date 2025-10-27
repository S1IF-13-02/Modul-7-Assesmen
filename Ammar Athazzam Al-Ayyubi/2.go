package main 
import (
	"fmt"
)
func main() {
	var a, b int 
	fmt.Print("masukan: ")
	fmt.Scan(&a, &b)

	hasil := 1
	for i := a; i <= b; i++ {
		hasil *= i
		fmt.Print(i)
		if i < b {
			fmt.Print(" x ")
		}
	}
	fmt.Printf("\nhasil: %d\n", hasil)
}