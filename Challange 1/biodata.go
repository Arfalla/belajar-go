package main

import (
	"fmt"
	"os"
)

// Struct untuk menyimpan data teman
type dataTeman struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

// Data teman dan no absennya
var daftarTeman = map[int]dataTeman{
	1:  {"Ardi", "Bekasi", "Software Engineer", "Ingin mempelajari bahasa pemrograman Go untuk meningkatkan skill."},
	2:  {"Bayu", "Semarang", "Flutter Developer", "Belajar hal baru untuk integrasi sistem dari flutter ke back end"},
	3:  {"Daffa", "Cikarang", "Web Developer", "Membutuhkan bahasa pemrograman yang efisien untuk pengembangan aplikasi web."},
	4:  {"Ilham", "Karawang", "Back End", "Memperdalam bahasa baru untuk ranah back end."},
	5:  {"Kurnia", "Bandung", "Freelancer", "Mempelajari hal baru untuk pengembangan diri."},
	6:  {"Miftahul", "Depok", "UI Designer", "Mempersiapkan skill baru untuk dunia kerja."},
	7:  {"Nurul", "Solo", "Cyber Security", "Mempelajari mekanisme dan kelebihan/ kekurangan dari backend golang untuk keamanan sistem nantinya."},
	8:  {"Rafi", "Jogja", "Data Engineer", "Menambah wawasan tentang bahasa pemrograman yang sedang naik daun."},
	9:  {"Rahmat", "Medan", "Front End", "Ingin mempelajari siklus dibalik back end."},
	10: {"Rizki", "Jakarta Timur", "Data Scientist", "Tertarik dengan kecepatan dan kesederhanaan Go dalam pengembangan perangkat lunak."},
}

// Function untuk menampilkan data teman berdasarkan nomor absen
func tampilkanData(absen int) {
	teman, temukan := daftarTeman[absen]
	if !temukan {
		fmt.Println("Data dengan nomor absen", absen, "tidak ditemukan.")
		return
	}
	fmt.Println("Menampilkan Informasi Absen:", absen)
	fmt.Println("Nama:", teman.Nama)
	fmt.Println("Alamat:", teman.Alamat)
	fmt.Println("Pekerjaan:", teman.Pekerjaan)
	fmt.Println("Alasan memilih kelas Golang:", teman.Alasan)
}

func main() {
	// Agar tidak panic saat debug
	if len(os.Args) < 2 {
		fmt.Println("Gunakanlah: go run biodata.go <nomor_absennya>")
		return
	}
	absen := os.Args[1]
	var absenInt int
	fmt.Sscanf(absen, "%d", &absenInt)

	tampilkanData(absenInt)
}

