package main

import "fmt"

func main() {
	var umur int8
	var anakKe int8
	var saldo int64

	umur = 19
	anakKe = 1
	saldo = 15000

	fmt.Println("Saya berumur: ", &umur)
	fmt.Println("Saya anak ke: ", anakKe)
	fmt.Println("Saldo saya: ", saldo)
}
