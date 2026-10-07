package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	var sepuluh int = x / 10000
	var sisa int = x % 10000

	var lima int = sisa / 5000
	sisa = sisa % 5000

	var satu int = sisa / 1000

	fmt.Println(sepuluh, lima, satu)
}
