package main

import "fmt"

func main() {
    var keping int

    fmt.Print("Masukkan keping: ")
    fmt.Scan(&keping)

    peti := keping / 800
    sisa := keping % 800

    karung := sisa / 100
    sisa = sisa % 100

    ikat := sisa / 10
    sisa = sisa % 10

    k := sisa

    fmt.Println(peti, karung, ikat, k)
}