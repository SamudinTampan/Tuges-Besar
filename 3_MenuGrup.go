package main

import (
	"bufio"
	"fmt"
	"os"
)

func menuGrup(namaAkun string, idAkun int) {
	var selesai bool = false
	var namaGrup string
	var input int
	for !selesai {
		fmt.Println("-------------------------------------")
		fmt.Println("-             MENU GRUP             -")
		fmt.Println("-------------------------------------")
		DaftarGrup(idAkun)
		fmt.Println("-------------------------------------")
		fmt.Println("Masukan '0' Untuk Kembali")
		fmt.Print("Pilih? ")
		fmt.Scan(&input)
		fmt.Print("\n")

		if input == 0 {
			selesai = true
		} else {
			namaGrup = User[idAkun].Grup[input].namaGrup
			RuangPesanGrup(namaAkun, namaGrup)
		}
	}
}

func RuangPesanGrup(namaAkun, namaGrup string) {
	var selesai bool = false
	var idGrup, idAkunDituju, idAkun, totalPesan int
	var pesan, namaKontakBaru string
	idAkun = cariUserID(namaAkun)
	idGrup = cariGrupID(idAkun, namaGrup)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	for !selesai {
		fmt.Println("+-------------------------------------------+") //37
		fmt.Printf("| %10s                                |\n", namaAkun)
		fmt.Println("+-------------------------------------------+")
		if User[idAkun].Grup[idGrup].totalPesan != 0 {
			DaftarPesanGrup(idAkun, idGrup)
		} else {
			fmt.Println("-- Tidak ada pesan --")
		}
		fmt.Println("+-------------------------------------------+")
		fmt.Println("| 0. Kembali                                |")
		fmt.Println("| 1. Deskripsi Grup                         |")
		fmt.Println("+-------------------------------------------+")
		fmt.Print("Krim Pesan : ")
		pesan = input()

		if pesan == "0" {
			selesai = true
		} else if pesan == "1" {
			fmt.Print("\n")
			DeskripsiGrup(idAkun, idGrup, namaAkun, namaKontakBaru, namaGrup)
		} else if pesan == "2" {
			fmt.Print("\n")
			DaftarAnggotaGrup(idAkun, idGrup)
		} else {
			for i := 1; i <= User[idAkun].Grup[idGrup].totalAnggota; i++ {
				idAkunDituju = cariUserID(User[idAkun].Grup[idGrup].anggotaGrup[i])
				idGrup = cariGrupID(idAkunDituju, namaGrup)
				User[idAkunDituju].Grup[idGrup].totalPesan += 1
				totalPesan = User[idAkunDituju].Grup[idGrup].totalPesan
				User[idAkunDituju].Grup[idGrup].pesanGrup[totalPesan].isipesan = pesan
				User[idAkunDituju].Grup[idGrup].pesanGrup[totalPesan].pengirim = namaAkun
			}
		}
	}
}

func DaftarPesanGrup(idAkun, idGrup int) {
	for i := 1; i <= User[idAkun].Grup[idGrup].totalPesan; i++ {
		fmt.Printf(" %8s  :  %s                            \n", User[idAkun].Grup[idGrup].pesanGrup[i].pengirim, User[idAkun].Grup[idGrup].pesanGrup[i].isipesan)
	}
}

func DaftarAnggotaGrup(idAkun, idGrup int) {
	fmt.Println("+-- Anggota Grup -------------------+")
	for i := 1; i <= User[idAkun].Grup[idGrup].totalAnggota; i++ {
		fmt.Printf(" %d. %s \n", i, User[idAkun].Grup[idGrup].anggotaGrup[i])
	}
	fmt.Println("+-----------------------------------+")
	fmt.Print("\n")
}

func DeskripsiGrup(idAkun, idGrup int, namaAkun, namaKontakBaru, namaGrup string) {
	var selesai bool = false
	var pesan string

	for !selesai {
		fmt.Println("+-- Deskripsi Grup -----------------+")
		fmt.Println("| 0. Kembali                        |")
		fmt.Println("| 1. Tambah Anggota Grup            |")
		fmt.Println("| 2. Anggota Grup                   |")
		fmt.Println("+-----------------------------------+")
		fmt.Print("Krim Pesan : ")
		fmt.Scan(&pesan)
		if pesan == "0" {
			selesai = true
		} else if pesan == "1" {
			fmt.Print("\n")
			TampilanTambahKontakkeGrup(&namaKontakBaru)
			TambahKontakkeGrup(namaAkun, namaGrup, namaKontakBaru, idAkun)
		} else if pesan == "2" {
			fmt.Print("\n")
			DaftarAnggotaGrup(idAkun, idGrup)
		}
	}
	pesan = input()
}
