package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("Nilai a:", a)
	fmt.Scan(&a)
	fmt.Print("Nilai b:", b)
	fmt.Scan(&b)
		
	
	tambah := a + b
	kurang := a - b
	kali := a * b
	bagi := a / b
	sisa := a % b

	fmt.Println(tambah)
	fmt.Println(kurang)
	fmt.Println(kali)
	fmt.Println(bagi)
	fmt.Println(sisa)
}
