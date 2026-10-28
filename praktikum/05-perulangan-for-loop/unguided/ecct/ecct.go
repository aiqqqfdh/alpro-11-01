package main

import "fmt"

func main() {
	var a float64
	var lulus int
	n := 6
	b := 2

	for u := 1; u <= b; u++ {
		var hasil float64

		for i := 1; i <= n; i++ {
			hasil = hasil + a
			fmt.Scan(&a)

			if a >= 5 {
				break
			}
		}

		hasilA := hasil / 6

		if hasilA >= 3.33 {
			fmt.Println("lulus")
			lulus++
		} else {
			fmt.Println("Tidak Lulus")
		}

	}
	fmt.Println(lulus)
}



