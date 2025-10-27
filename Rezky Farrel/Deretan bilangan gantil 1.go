package main
import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var bilanganGanjil = 1
	for i := 0; i < n; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(bilanganGanjil)
		bilanganGanjil += 2
	}
	fmt.Println()
}