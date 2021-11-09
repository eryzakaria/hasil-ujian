package main

import "fmt"

func main() {
	fmt.Printf("Masukan Nama Anda: ")
	var nama string
	fmt.Scan(&nama)
	fmt.Printf("Masukan Nama NIM Anda: ")
	var nim string
	fmt.Scan(&nim)
	fmt.Printf("Masukan Nilai Tugas: ")
	var tugas int
	fmt.Scan(&tugas)
	fmt.Printf("Masukan Nilai UTS: ")
	var uts int
	fmt.Scan(&uts)
	fmt.Printf("Masukan Nilai UAS: ")
	var uas int
	fmt.Scan(&uas)
	fmt.Printf("Masukan Absensi : ")
	var absensi int
	fmt.Scan(&absensi)

	fmt.Println(nama, nim)
	result := uts + uas + tugas
	result = result / 3
	if result > 101 && absensi > 101 {
		fmt.Println("Masukan Nilai Yang Benar")
	} else if result >= 90 && absensi >= 90 {
		fmt.Println("Anda Dinyatakan Lulus (Cumlaude)", result)
	} else if result >= 70 && absensi >= 70 {
		fmt.Println("Anda Dinyatakan Lulus", result)
	} else {
		fmt.Println("Anda Belum Lulus", result)
	}
}
