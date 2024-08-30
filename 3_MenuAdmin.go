package main

import "fmt"

func TampilanAwalAdmin() {
	//show admin display containt option of "Daftar Akun" and "Approve Account"
	var input int
	var selesai bool = false
	for !selesai {
		TampilanAdmin()
		fmt.Print("Pilih?: ")
		fmt.Scan(&input)
		fmt.Print("\n")
		if input == 1 {
			daftarAkun()
			fmt.Print("\n")
		} else if input == 2 {
			TerimaAKun()
		} else if input == 3 {
			TolakAkun()
		} else if input == 0 {
			selesai = true
			fmt.Print("\n")
		}
	}
}

func daftarAkun() {
	// Print data acc that approved
	fmt.Println("+--- DAFTAR AKUN ---------------------------------------------+")
	fmt.Printf("| %2s | %20s | %15s | %13s |\n", "ID", "Nama", "Password", "Status")
	fmt.Println("+-------------------------------------------------------------+")
	for i := 0; i < User[0].TotalUser; i++ {
		fmt.Printf("| %2d | %20s | %15s | %13s |\n", User[i].ID, User[i].Username, User[i].Password, User[i].StatusAkun)
	}
	fmt.Println("+-------------------------------------------------------------+")
}

func TerimaAKun() {
	var idAkun int
	fmt.Println("-- ADMIN ----------------------------")
	fmt.Print("Masukan ID akun yang ingin diterima: ")
	fmt.Scan(&idAkun)
	fmt.Println("-------------------------------------")
	if idAkun <= User[0].TotalUser {
		User[idAkun].StatusAkun = "Diterima"
		fmt.Printf("AKun '%s' Dengan ID %d Diterima\n", User[idAkun].Username, idAkun)
	} else {
		fmt.Println("ID Yang Anda Masukan Tidak Ada")
	}
	fmt.Print("\n")
}

func TolakAkun() {
	var idAkun int
	fmt.Println("-- ADMIN ----------------------------")
	fmt.Print("Masukan ID akun yang ingin ditolak: ")
	fmt.Scan(&idAkun)
	fmt.Println("-------------------------------------")
	if idAkun <= User[0].TotalUser {
		fmt.Printf("AKun '%s' Dengan ID %d Ditolak\n", User[idAkun].Username, idAkun)
		HapusAkun(idAkun)
		urutUser()
	} else {
		fmt.Println("ID Yang Anda Masukan Tidak Ada")
	}
	fmt.Print("\n")
}

func HapusAkun(idAkun int) {
	for i := idAkun; i < User[0].TotalUser; i++ {
		User[i] = User[i+1]
	}
	User[0].TotalUser -= 1
}
