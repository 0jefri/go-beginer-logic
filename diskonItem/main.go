package main

import (
	"fmt"
)

// Fungsi untuk menghitung harga setelah diskon
func hargaSetelahDiskon(harga []float64, jumlahBeli int) float64 {
	// hargaSetelahDiskon := make([]float64, jumlahBeli)
	// var diskon float64

	// // Menentukan persentase diskon berdasarkan jumlah item yang dibeli
	// switch {
	// case jumlahBeli == 1:
	// 	diskon = 0
	// case jumlahBeli == 2 || jumlahBeli == 3:
	// 	diskon = 10
	// case jumlahBeli == 4:
	// 	diskon = 50
	// case jumlahBeli > 4:
	// 	diskon = 75
	// }

	// // Menghitung harga setelah diskon
	// for i, h := range harga[:jumlahBeli] {
	// 	hargaSetelahDiskon[i] = h * (1 - diskon/100)
	// }

	// return hargaSetelahDiskon

	var totalHarga float64
	var diskon float64

	// Menentukan persentase diskon berdasarkan jumlah item yang dibeli
	switch {
	case jumlahBeli == 1:
		diskon = 0 // Tidak ada diskon
	case jumlahBeli == 2 || jumlahBeli == 3:
		diskon = 10 // Diskon 10% untuk 2 atau 3 item
	case jumlahBeli == 4:
		diskon = 50 // Diskon 50% untuk 4 item
	case jumlahBeli > 4:
		diskon = 75 // Diskon 75% untuk lebih dari 4 item
	}

	// Menghitung total harga dari item yang dibeli
	for _, h := range harga[:jumlahBeli] {
		totalHarga += h
	}

	// Menghitung total harga setelah diskon
	totalHargaSetelahDiskon := totalHarga * (1 - diskon/100)

	return totalHargaSetelahDiskon
}

func main() {
	// Harga item yang dijual
	// hargaJual := []float64{100000, 200000, 150000, 300000, 250000, 500000}
	// var jumlahBeli int

	// // Meminta input jumlah item yang dibeli
	// fmt.Print("Masukkan jumlah item yang dibeli: ")
	// fmt.Scan(&jumlahBeli)

	// // Mengecek apakah jumlah beli valid
	// if jumlahBeli < 1 {
	// 	fmt.Println("Jumlah item yang dibeli tidak valid.")
	// 	return
	// }

	// // Menghitung harga setelah diskon
	// hargaSetelahDiskon := hargaSetelahDiskon(hargaJual, jumlahBeli)

	// // Menampilkan hasil
	// fmt.Println("Harga setelah diskon:")
	// for i, h := range hargaSetelahDiskon {
	// 	fmt.Printf("Harga item %d: %.2f\n", i+1, h)
	// }

	hargaJual := []float64{100000, 200000, 150000, 300000, 250000, 500000}
	var jumlahBeli int

	// Meminta input jumlah item yang dibeli
	fmt.Print("Masukkan jumlah item yang dibeli: ")
	fmt.Scan(&jumlahBeli)

	// Mengecek apakah jumlah beli valid
	if jumlahBeli < 1 || jumlahBeli > len(hargaJual) {
		fmt.Println("Jumlah item yang dibeli tidak valid.")
		return
	}

	// Menghitung total harga setelah diskon
	totalHargaSetelahDiskon := hargaSetelahDiskon(hargaJual, jumlahBeli)

	// Menampilkan hasil
	fmt.Printf("Total harga setelah diskon: %.2f\n", totalHargaSetelahDiskon)
}
