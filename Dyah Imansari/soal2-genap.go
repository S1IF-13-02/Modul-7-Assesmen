package main
import "fmt"
func main() {
	var x, y, i int
	var hasil int
	hasil = 1

	fmt.Scan(&x, &y)
	for i = x; i <= y; i++ {
		hasil = hasil * i
	}
	fmt.Println("Jumlah bakteri terakhir = ", hasil)
}