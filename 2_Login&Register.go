package main

import "fmt"

func RegistrationAcc() {
	//User registrasi akun
	var temp dataUser
	var selesai bool

	for !selesai {
		fmt.Println("+-----------------------------------+")
		fmt.Println("|   	      REGISTRASI            |")
		fmt.Println("+-----------------------------------+")
		fmt.Println("| 0. Kembali                        |")
		fmt.Println("+-----------------------------------+")
		fmt.Print("  Username  : ")
		fmt.Scan(&temp.Username)
		if temp.Username == "0" {
			selesai = true
		} else {
			fmt.Print("  Password  : ")
			fmt.Scan(&temp.Password)
			fmt.Print("\n")
			if temp.Username == "0" {
				selesai = true
			} else if CekStatusAkun(temp.Username, temp.Password, "Reg") {
				fmt.Println("-- Username sudah terdaftar --")
			} else {
				User[User[0].TotalUser].Username = temp.Username
				User[User[0].TotalUser].Password = temp.Password
				fmt.Println("-- Akun berhasil didaftarkan, menunggu persetujuan admin --")
				User[User[0].TotalUser].StatusAkun = "BelumDiterima"
				User[User[0].TotalUser].TotalTeman = 0
				User[User[0].TotalUser].TotalGrup = 0
				User[0].TotalUser += 1
				urutUser()
			}
		}
		fmt.Print("\n")
	}
}

func urutUser() {
	//Mengurutkan username akun user secara menaik
	var i int = User[0].TotalUser - 1
	temp := User[i]
	for i > 1 && temp.Username < User[i-1].Username {
		User[i] = User[i-1]
		i--
	}
	User[i] = temp
	for j := 1; j <= User[0].TotalUser; j++ {
		User[j].ID = j
	}
}

func LoginAcc() {
	//User login akun
	var nama, password string
	var selesai bool = false

	for !selesai {
		fmt.Println("+-----------------------------------+")
		fmt.Println("|   	      L O G I N             |")
		fmt.Println("+-----------------------------------+")
		fmt.Println("| 0. Kembali                        |")
		fmt.Println("+-----------------------------------+")
		fmt.Print("  Username: ")
		fmt.Scan(&nama)
		if nama == "0" {
			selesai = true
		} else {
			fmt.Print("  Password: ")
			fmt.Scan(&password)
			if password == "0" {
				selesai = true
				continue
			} else if nama == "admin" && password == "123" {
				fmt.Print("\n")
				TampilanAwalAdmin()
			} else if CekStatusAkun(nama, password, "Log") {
				fmt.Print("\n")
				TampilanMenu(nama)
			} else {
				fmt.Print("\n")
				fmt.Println("Username atau Passord Salah, atau akun belum disetujui admin")
			}
		}
		fmt.Print("\n")
	}
}
