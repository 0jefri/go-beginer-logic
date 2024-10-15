package main

import "fmt"

// Menghitung luas dan keliling bangunan datar
func luasPersegiPanjang(panjang int, lebar int) (int, int) {
	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)
	return luas, keliling
}

func main() {
	var panjang, lebar int

	fmt.Print("Masukan panjang: ")
	fmt.Scan(&panjang)

	fmt.Print("Masukkan lebar: ")
	fmt.Scan(&lebar)

	luas, keliling := luasPersegiPanjang(panjang, lebar)
	fmt.Printf("Luas Persegi Panjang: %d\n", luas)
	fmt.Printf("Keliling Persegi Panjang: %d\n", keliling)
}
