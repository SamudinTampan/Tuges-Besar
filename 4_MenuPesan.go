package main

import (
	"bufio"
	"fmt"
	"os"
)

func menuPesan(namaAkun string, idAkun int) {
	var selesai bool = false
	var namaAkunDituju string
	var input int
	for !selesai {
		fmt.Println("-------------------------------------")
		fmt.Println("-             MENU PESAN            -")
		fmt.Println("-------------------------------------")
		DaftarKontak(idAkun)
		fmt.Println("-------------------------------------")
		fmt.Println("Masukan '0' Untuk Kembali")
		fmt.Print("Pilih? ")
		fmt.Scan(&input)
		fmt.Print("\n")

		if input == 0 {
			selesai = true
		} else {
			namaAkunDituju = User[idAkun].Teman[input].nama
			RuangPesan(namaAkun, namaAkunDituju)
		}
	}
}

func RuangPesan(namaAkun, namaAkunDituju string) {
	var selesai bool = false
	var idAkunDituju, idTemanAkunDituju, idAkun, totalPesan int
	var pesan string
	idAkun = cariUserID(namaAkun)
	idTemanAkunDituju = cariUserIDteman(idAkun, namaAkunDituju)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	for !selesai {
		fmt.Println("+-------------------------------------------+")
		fmt.Printf("| %10s                                |\n", namaAkun)
		fmt.Println("+-------------------------------------------+")
		if User[idAkun].Teman[idTemanAkunDituju].totalPesan > 0 {
			DaftarPesan(idAkun, idTemanAkunDituju)
		} else {
			fmt.Println("-- Tidak ada pesan --")
		}
		fmt.Println("+-------------------------------------------+")
		fmt.Println("|  0. Kembali                               |")
		fmt.Println("+-------------------------------------------+")
		fmt.Print("Krim Pesan : ")
		pesan = input()

		if pesan == "0" {
			selesai = true
		} else {
			User[idAkun].Teman[idTemanAkunDituju].totalPesan += 1
			totalPesan = User[idAkun].Teman[idTemanAkunDituju].totalPesan
			User[idAkun].Teman[idTemanAkunDituju].pesan[totalPesan].isiPesan = pesan
			User[idAkun].Teman[idTemanAkunDituju].pesan[totalPesan].pengirim = namaAkun

			idAkunDituju = cariUserID(namaAkunDituju)
			User[idAkunDituju].Teman[idTemanAkunDituju].totalPesan += 1
			totalPesan = User[idAkunDituju].Teman[idTemanAkunDituju].totalPesan
			User[idAkunDituju].Teman[idTemanAkunDituju].pesan[totalPesan].isiPesan = pesan
			User[idAkunDituju].Teman[idTemanAkunDituju].pesan[totalPesan].pengirim = namaAkun
		}
	}
}

func DaftarPesan(idAkun, idTemanAkunDituju int) {
	for i := 1; i <= User[idAkun].Teman[idTemanAkunDituju].totalPesan; i++ {
		fmt.Printf(" %8s  :  %s                            \n", User[idAkun].Teman[idTemanAkunDituju].pesan[i].pengirim, User[idAkun].Teman[idTemanAkunDituju].pesan[i].isiPesan)
	}
}
