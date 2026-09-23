package main

import "fmt"

func main() {

	var nama string
	var mtk, bIng int

	fmt.Scan(&nama, &mtk, &bIng)

	total := mtk + bIng
	rataRata := total / 2

	fmt.Println(nama)
	fmt.Println(total)
	fmt.Println(rataRata)
}