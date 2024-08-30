package main

import "fmt"

func TampilanTambahKontakkeGrup(namaKontakBaru *string) {
	fmt.Println("-- Tambah Kontak ke Grup ------------")
	fmt.Print("Cari nama : ")
	fmt.Scan(namaKontakBaru)
	fmt.Println("-------------------------------------")
	fmt.Print("\n")
}

func TampilanTambahKontak(namaKontakBaru *string) {
	fmt.Println("-- Tambah Kontak --------------------")
	fmt.Print("Cari nama : ")
	fmt.Scan(namaKontakBaru)
	fmt.Println("-------------------------------------")
	fmt.Print("\n")
}

func TampilanAdmin() {
	fmt.Println("+--- ADMIN -------------------------+")
	fmt.Println("| 1. Daftar Akun                    |")
	fmt.Println("| 2. Terima Akun                    |")
	fmt.Println("| 3. Tolak Akun                     |")
	fmt.Println("| 0. Kembail                        |")
	fmt.Println("+-----------------------------------+")
}

func TampilanUser() {
	fmt.Println("+-----------------------------------+")
	fmt.Println("|             MAIN MENU             |")
	fmt.Println("+-----------------------------------+")
	fmt.Println("| 1. Registrasi                     |")
	fmt.Println("| 2. Login                          |")
	fmt.Println("| 3. Keluar                         |")
	fmt.Println("+-----------------------------------+")
}
