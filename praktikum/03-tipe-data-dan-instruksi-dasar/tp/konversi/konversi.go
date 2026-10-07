package main

import (
	"fmt"

)

func main() {
	var mil float64
	var kilo float64

	fmt.Scan(&mil)

	kilo = (mil * 1.6)

	fmt.Printf("%.1f\n", kilo)
}
