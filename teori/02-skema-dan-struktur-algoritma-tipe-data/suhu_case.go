package main

import "fmt"

func main(){
	var umur int8
	var suhu float32

	umur = 20
	suhu = 36.5

	fmt.Println("Umur:", umur)
	fmt.Println("Suhu:", suhu)
	fmt.Println("Alamat memori Umur:", &umur)
	fmt.Println("Alamat memori Suhu:", &suhu)
} 