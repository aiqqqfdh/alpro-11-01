package main

import "fmt"

func main() {
	var suhu int
	fmt.Scan(&suhu)

	switch {
	case suhu <= 0:
		fmt.Println("Dingin Sekali")
	case suhu >= 10 && suhu <= 20:
		fmt.Println("Dingin")
	case suhu > 20 && suhu <= 30:
		fmt.Println("Normal")
	case suhu > 30 && suhu <= 40:
		fmt.Println("Panas")
	}
}