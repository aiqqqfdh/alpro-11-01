package main

import "fmt"

func main() {
	var hari int

	fmt.Scan(&hari)

	var tahun int = hari / 360
	var sisa int = hari % 360

	var bulan int = sisa / 30
	sisa = sisa % 30

	var minggu int = sisa / 7
	var hariSisa int = sisa % 7



	fmt.Println(tahun, bulan, minggu, hariSisa)
}