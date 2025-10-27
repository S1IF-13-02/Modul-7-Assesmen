package main

import "fmt"
func main (){
	var x, y, total int
	fmt.Print("Masukan angka awal hari : ")
	fmt.Scan(&x)
	fmt.Print("Masukan angka akhir hari : ")
	fmt.Scan(&y)

	total = 1
	for i := x ; i<=y ; i++{
		total = total * i
	}
	fmt.Println(total)
}