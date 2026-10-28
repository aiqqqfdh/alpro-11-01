package main

import "fmt"

func main() {
	var t int
	fmt.Print("Input t(jumlah transaksi yang akan dilakukan): ")
	fmt.Scan(&t)

	totalP := 0

	for i := 1; i <= t; i++ {
		var film, hari string
		var tTiket int

		fmt.Printf("Input transaksi ke-%d (film, hari, tiket)", i)
		fmt.Scan(&film, &hari, &tTiket)

		hargaS := 0

		if hari == "Senin" || hari == "Selasa" || hari == "Rabu" || hari == "Kamis" {
			if film == "Reguler" {
				hargaS = 40000
			} else if film == "IMAX" {
				hargaS = 50000
			} else if film == "3D" {
				hargaS = 60000
			}
		}

		if hari == "Jumat" || hari == "Sabtu" || hari == "Minggu" {
			if film == "Reguler" {
				hargaS = 45000
			} else if film == "IMAX" {
				hargaS = 55000
			} else if film == "3D" {
				hargaS = 65000
			}
		}
		totalH := hargaS * tTiket

		if tTiket >= 5 {
			diskon := (totalH * 10) / 100
			totalH = totalH - diskon
		}

		fmt.Printf("output total harga transaksi ke-%d: %d\n", i, totalH)

		totalP = totalP + totalH
	}

	fmt.Println("output total pendapatan:", totalP)
}

	

