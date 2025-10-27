package main
import "fmt"
func main() {
	var n, i int
	var hasil int
	hasil = 0
	
	fmt.Print("input:")
	fmt.Scanln(&n)
	for i = 1; i <= n; i++ {
		hasil += i
		fmt.Print(i/2, " ")
	}
}