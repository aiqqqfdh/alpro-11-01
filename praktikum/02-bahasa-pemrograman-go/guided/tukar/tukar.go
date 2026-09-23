package main

import "fmt"

func main() {
	var a,b int

	//membaca input
	fmt.Scan(&a,&b)

	//menukar nilai a dan b

	a,b = b,a

	//menampilkan Output

	fmt.Println(a)
	fmt.Println(b)
}