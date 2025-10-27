package main
import "fmt"

func main() {
	var x, y int
	fmt.Print("Masukan x dan y: ")
	fmt.Scan(&x, &y)

	total := 1

	for i := x; i <= y; i++ {
		total *= i	
	}
	fmt.Println("Output:", total)
}