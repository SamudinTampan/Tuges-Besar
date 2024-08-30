package main

import "fmt"

func TampilanMenu(namaAkun string) {
	var selesai = false
	var input, idAkun int
	var namaKontakBaru string
	idAkun = cariUserID(namaAkun)
	for !selesai {

		fmt.Println("+-------------------------------------+")
		fmt.Println("|             K O N T A K      	      |")
		fmt.Println("|-------------------------------------+")
		fmt.Println("|-- KONTAK ---------------------------+")
		if User[idAkun].TotalTeman != 0 {
			DaftarKontak(idAkun)
		} else {
			fmt.Println("Tidak Ada Kontak")
		}
		fmt.Println("+-- GRUP -----------------------------+")
		if User[idAkun].TotalGrup != 0 {
			DaftarGrup(idAkun)
		} else {
			fmt.Println("Tidak ada grup")
		}
		fmt.Println("+-------------------------------------+")
		fmt.Println("| 1. Menu Pesan Pribadi               |")
		fmt.Println("| 2. Menu Pesan Grup                  |")
		fmt.Println("| 3. Tambah Kontak                    |")
		fmt.Println("| 4. Buat Grup                        |")
		fmt.Println("| 5. Kembali                          |")
		fmt.Println("+-------------------------------------+")
		fmt.Print("Pilih (1/2/3/4/5)?: ")
		fmt.Scan(&input)
		fmt.Print("\n")

		if input == 0 {
			selesai = true
		} else if input == 1 {
			menuPesan(namaAkun, idAkun)
		} else if input == 2 {
			menuGrup(namaAkun, idAkun)
		} else if input == 3 {
			TampilanTambahKontak(&namaKontakBaru)
			TambahKontak(namaAkun, namaKontakBaru, idAkun)
		} else if input == 4 {
			buatGrup(namaAkun, idAkun)
		} else if input == 5 {
			selesai = true
		}
	}
}

func DaftarKontak(idAkun int) {
	// Menampilkan daftar teman user dengan id = idAkun
	for i := 1; i <= User[idAkun].TotalTeman; i++ {
		fmt.Println(i, ". ", User[idAkun].Teman[i].nama)
	}
}

func DaftarGrup(idAkun int) {
	// Menampilkan daftar grup user dengan id = idAkun
	for i := 1; i <= User[idAkun].TotalGrup; i++ {
		fmt.Println(i, ". ", User[idAkun].Grup[i].namaGrup)
	}
}

func buatGrup(namaAkun string, idAkun int) {
	var namaGrup, YoN string

	fmt.Println("-- Buat Grup ---=--------------------")
	fmt.Print("Masukan Nama Grup : ")
	fmt.Scan(&namaGrup)
	fmt.Println("-------------------------------------")
	fmt.Print("\n")

	if CekStatusGrup(idAkun, namaGrup) {
		fmt.Println("-------------------------------------")
		fmt.Printf("Anda SUdah Memiliki Grup Dengan Nama %s!\n", namaGrup)
		fmt.Println("-------------------------------------")
		fmt.Print("\n")
	} else {
		fmt.Println("-- Buat Grup ---=--------------------")
		fmt.Printf("Tambahkan Grup Dengan Nama '%s'?\n", namaGrup)
		fmt.Print("(Y/N)?: ")
		fmt.Scan(&YoN)
		fmt.Println("-------------------------------------")

		if YoN == "Y" {
			User[idAkun].TotalGrup += 1
			User[idAkun].Grup[User[idAkun].TotalGrup].totalAnggota = 1
			User[idAkun].Grup[User[idAkun].TotalGrup].anggotaGrup[1] = namaAkun
			User[idAkun].Grup[User[idAkun].TotalGrup].namaGrup = namaGrup
		}

	}
}

func TambahKontak(namaAkun, namaKontakBaru string, idAkun int) {
	var YoN string
	var idAkunDituju int

	if CekStatusAkun(namaKontakBaru, "0", "cariNama") && namaKontakBaru != namaAkun {
		fmt.Println("-- Tambah Kontak --------------------")
		fmt.Printf("Tambahkan Akun Dengan Nama '%s'?\n", namaKontakBaru)
		fmt.Print("Tambahkan Akun? (Y/N)?: ")
		fmt.Scan(&YoN)
		fmt.Println("-------------------------------------")

		if YoN == "Y" {
			User[idAkun].TotalTeman += 1
			User[idAkun].Teman[User[idAkun].TotalTeman].nama = namaKontakBaru
			User[idAkun].Teman[User[idAkun].TotalTeman].totalPesan = 0
			idAkunDituju = cariUserID(namaKontakBaru)
			User[idAkunDituju].TotalTeman += 1
			User[idAkunDituju].Teman[User[idAkunDituju].TotalTeman].nama = namaAkun
			User[idAkunDituju].Teman[User[idAkunDituju].TotalTeman].totalPesan = 0
		}
	} else {
		fmt.Println("Akun Tidak Ditemukan")
	}
}

//DIGUNAKAN
func TambahKontakkeGrup(namaAkun, namaGrup, namaKontakBaru string, idAkun int) {
	var YoN string
	var idAkunDituju, idGrup, idGrupTeman, totalAnggota int
	idGrup = cariGrupID(idAkun, namaGrup)

	if CekStatusAkun(namaAkun, namaKontakBaru, "cariNamaTeman") && CekStatusAkun(namaKontakBaru, "0", "cariNama") {
		fmt.Println("-- Tambah Kontak ke Grup ------------")
		fmt.Printf("Tambahkan Akun Dengan Nama '%s'?\n", namaKontakBaru)
		fmt.Print("Tambahkan Akun? (Y/N)?: ")
		fmt.Scan(&YoN)
		fmt.Println("-------------------------------------")

		if YoN == "Y" {
			User[idAkun].Grup[idGrup].totalAnggota += 1
			totalAnggota = User[idAkun].Grup[idGrup].totalAnggota

			for i := 1; i < User[idAkun].Grup[idGrup].totalAnggota; i++ {
				idAkunDituju = cariUserID(User[idAkun].Grup[idGrup].anggotaGrup[i])
				idGrupTeman = cariGrupID(idAkunDituju, namaGrup)
				User[idAkunDituju].Grup[idGrupTeman].totalAnggota = totalAnggota
				User[idAkunDituju].Grup[idGrupTeman].anggotaGrup[totalAnggota] = namaKontakBaru
			}
			idAkunDituju = cariUserID(User[idAkun].Grup[idGrup].anggotaGrup[totalAnggota])
			User[idAkunDituju].TotalGrup += 1
			User[idAkunDituju].Grup[User[idAkunDituju].TotalGrup].namaGrup = namaGrup
			User[idAkunDituju].Grup[User[idAkunDituju].TotalGrup].totalAnggota = User[idAkun].Grup[idGrup].totalAnggota
			User[idAkunDituju].Grup[User[idAkunDituju].TotalGrup].anggotaGrup = User[idAkun].Grup[idGrup].anggotaGrup
			User[idAkunDituju].Grup[User[idAkunDituju].TotalGrup].totalPesan = User[idAkun].Grup[idGrup].totalPesan
			User[idAkunDituju].Grup[User[idAkunDituju].TotalGrup].pesanGrup = User[idAkun].Grup[idGrup].pesanGrup
		}
	} else if !CekStatusAkun(namaAkun, namaKontakBaru, "cariNamaTeman") && CekStatusAkun(namaKontakBaru, "0", "cariNama") {
		fmt.Printf("Anda Belum Berteman Dengan %s. Tambahkan?\n", namaKontakBaru)
		fmt.Print("(Y/N)?: ")
		fmt.Scan(&YoN)
		fmt.Print("\n")
		if YoN == "Y" {
			TambahKontak(namaAkun, namaKontakBaru, idAkun)
			fmt.Print("\n")
			TambahKontakkeGrup(namaAkun, namaGrup, namaKontakBaru, idAkun)
			fmt.Print("\n")
		}
		fmt.Printf("Akun %s Berhasil Ditambahkan Ke Grup %s\n", namaKontakBaru, namaGrup)
	} else {
		fmt.Println("Akun Tidak Ditemukan")
	}
}
