package main

import "fmt"

func main() {
	var y, x, sisa int

	fmt.Scan(&y)
	fmt.Scan(&x)
	sisa = (y % x)
	fmt.Println(sisa)
}
