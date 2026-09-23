package main

import "fmt"

func main() {
	var uang int

	uang = 7000

	l10 := uang / 10000
	s10 := uang % 10000

	l5 := s10 / 5000
	s5 := s10 % 5000

	l1 := s5 / 1000

	fmt.Println(l10)

	fmt.Println(l5)

	fmt.Println(l1)


}
