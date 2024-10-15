package main

import (
	"fmt"
)

// Fungsi untuk menghitung total harga setelah diskon
func totalHargaSetelahDiskon(sepatu1, sepatu2 string, namaSepatu []string, hargaSepatu []float64) float64 {
	var totalHarga, diskon float64

	//ini buat seleksi sepatu nanti tentukan harganya
	for i, nama := range namaSepatu {
		if nama == sepatu1 || nama == sepatu2 {
			totalHarga += hargaSepatu[i]
		}
	}

	// Menentukan diskon setelah pilih sepatu
	if (sepatu1 == "adidas" && sepatu2 == "puma") || (sepatu1 == "puma" && sepatu2 == "adidas") {
		diskon = 50000.0
	} else if (sepatu1 == "puma" && sepatu2 == "kappa") || (sepatu1 == "kappa" && sepatu2 == "puma") {
		diskon = 150000.0
	} else if (sepatu1 == "adidas" && sepatu2 == "kappa") || (sepatu1 == "kappa" && sepatu2 == "adidas") {
		diskon = 75000.0
	} else {
		fmt.Println("Sepatu yang dipilih tidak memenuhi syarat untuk diskon")
		return 0
	}

	// total harga setelah diskon
	totalHargaSetelahDiskon := totalHarga - diskon
	return totalHargaSetelahDiskon
}

func main() {
	namaSepatu := []string{"adidas", "puma", "kappa"}
	hargaSepatu := []float64{200000, 150000, 600000}

	//jangan lupa input 2 sepatu yg mau dibeli
	var sepatu1 string
	fmt.Println("Masukkan sepatu pertama (adidas, puma, kappa):")
	fmt.Scanln(&sepatu1)

	var sepatu2 string
	fmt.Println("Masukkan sepatu kedua (adidas, puma, kappa):")
	fmt.Scanln(&sepatu2)

	// hitung total harga setelah diskon
	total := totalHargaSetelahDiskon(sepatu1, sepatu2, namaSepatu, hargaSepatu)

	// Jika total bukan 0, cetak hasilnya atau pengecekan bug
	if total != 0 {
		fmt.Printf("Total harga setelah diskon: Rp%.2f\n", total)
	}
}
