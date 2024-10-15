package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// var nilai int

	// fmt.Print("Massukan Nilai Siswa: ")
	// fmt.Scan(&nilai)

	// if nilai > 80 && nilai < 101 {
	// 	fmt.Println("Good")
	// }
	// if nilai >= 70 && nilai <= 80 {
	// 	fmt.Println("not good")
	// }

	// data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// for _, value := range data {
	// 	if value%2 != 0 {
	// 		fmt.Println(value)
	// 	}
	// }

	//menghitung hari
	// day := []int{1, 2, 3, 4, 5, 6, 7}

	// // fmt.Scan("Masukan Hari :", &day)
	// for _, value := range day {
	// 	var namaHari string

	// 	switch value {
	// 	case 1:
	// 		namaHari = "Senin"
	// 	case 2:
	// 		namaHari = "Selasa"
	// 	case 3:
	// 		namaHari = "Rabu"
	// 	case 4:
	// 		namaHari = "Kamis"
	// 	case 5:
	// 		namaHari = "Jumat"
	// 	case 6:
	// 		namaHari = "Sabtu"
	// 	case 7:
	// 		namaHari = "Minggu"
	// 	default:
	// 		namaHari = "Tidak valid"
	// 	}

	// 	fmt.Printf("Hari ke-%d: %s\n", value, namaHari)
	// }

	// //mencari index
	// name := []string{"jefri", "agistar", "putra", "andi", "budi"}

	// // Nilai yang ingin dicari
	// searchValue := "agistar"
	// var found bool
	// var index int

	// // Mencari indeks dari nilai yang dicari
	// for i, value := range name {
	// 	if value == searchValue {
	// 		index = i
	// 		found = true
	// 		break
	// 	}
	// }

	// // Menampilkan hasil
	// if found {
	// 	fmt.Printf("Nilai '%s' ditemukan pada indeks ke-%d.\n", searchValue, index)
	// } else {
	// 	fmt.Printf("Nilai '%s' tidak ditemukan dalam slice.\n", searchValue)
	// }

	// tentukan great nilai siswa dari nilai tertentu start great E = 0-5 D = 5 - 6 C = 6 - 7 B = 7 - 8 A = 80 -100

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Masukan Nilai: ")
	scanner.Scan()
	input := scanner.Text()

	nilai, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Input Tidak Valid, mohon masukan angka !!!")
		return
	}

	a := "A"
	b := "B"
	c := "C"
	d := "D"
	e := "E"
	wrongValue := "Nilai Tidak Valid, silahkan masukan nilai 0 - 100 !!!"

	if nilai < 50 {
		fmt.Printf("Nilainya Adalah : %s\n", e)
	} else if nilai <= 60 {
		fmt.Printf("Nilainya Adalah : %s\n", d)
	} else if nilai <= 70 {
		fmt.Printf("Nilainya Adalah : %s\n", c)
	} else if nilai <= 80 {
		fmt.Printf("Nilainya Adalah : %s\n", b)
	} else if nilai <= 100 {
		fmt.Printf("Nilainya Adalah : %s\n", a)
	} else {
		fmt.Printf("%s\n", wrongValue)
	}
}
