package main

import "fmt"

// fungsi Harga setelah diskon
func hargaSetelahDiskon(harga []float64, diskon float64) []float64 {
	hargaSetelahDiskon := make([]float64, len(harga))

	for i, h := range harga {
		hargaSetelahDiskon[i] = h * (1 - diskon/100)
	}

	return hargaSetelahDiskon
}

func main() {
	hargaJual := []float64{100000, 250000, 150000, 300000, 200000}
	var diskon float64

	fmt.Print("Masukkan persentase diskon: ")
	fmt.Scan(&diskon)

	// Menghitung harga setelah diskon
	hargaSetelahDiskon := hargaSetelahDiskon(hargaJual, diskon)

	// Menampilkan hasil
	fmt.Println("Harga setelah diskon:")
	for i, h := range hargaSetelahDiskon {
		fmt.Printf("Harga barang %d: %.2f\n", i+1, h)
	}
}
