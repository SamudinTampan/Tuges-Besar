package main

import (
	"bufio"
	"fmt"
	"os"
)

const MAXPESAN int = 100
const MAXAKUN int = 50
const MAXGRUP int = 10

var User ArrayDataUser
var Pesan ArrayPesan
var PesanGrup ArrayGrup

type ArrayDataUser [MAXAKUN]dataUser
type ArrayPesan [MAXPESAN]dataPesan
type ArrayGrup [MAXGRUP]dataGrup

type dataUser struct { //array User
	ID         int
	Username   string
	Password   string
	StatusAkun string
	JenisAkun  string //BELUM
	Teman      [MAXAKUN]dataTeman
	Grup       [MAXGRUP]dataGrup
	TotalUser  int
	TotalTeman int
	TotalGrup  int
}

type dataTeman struct {
	totalPesan int
	nama       string
	pesan      [MAXPESAN]dataPesan
}

type dataPesan struct {
	pengirim string
	isiPesan string
}

type dataGrup struct {
	namaGrup     string
	pesanGrup    [MAXPESAN]dataPesanGrup
	anggotaGrup  [MAXAKUN]string
	totalAnggota int
	totalPesan   int
}

type dataPesanGrup struct {
	pengirim string
	isipesan string
}

func main() {
	User[0].ID = 0
	User[0].Username = "Admin"
	User[0].Password = "123"
	User[0].StatusAkun = "Diterima"
	User[0].TotalUser = 1
	TampilanAwalUser()
}

func TampilanAwalUser() {
	//menampilkan tampilan awal dan dapat memilih registrasi,login,atau keluar
	var input int
	var selesai bool = false

	for !selesai {
		TampilanUser()
		fmt.Print("Pilih?: ")
		fmt.Scan(&input)
		fmt.Print("\n")
		if input == 1 {
			RegistrationAcc()
		} else if input == 2 {
			LoginAcc()
		} else if input == 3 {
			selesai = true
		}
	}
}

func cariUserID(namaAkun string) int {
	//
	var kiri, kanan, tengah, ID int
	kiri = 0
	kanan = User[0].TotalUser - 1
	ID = -1
	for kiri <= kanan && ID == -1 {
		tengah = (kiri + kanan) / 2
		if namaAkun < User[tengah].Username {
			kanan = tengah - 1
		} else if namaAkun > User[tengah].Username {
			kiri = tengah + 1
		} else {
			ID = User[tengah].ID
		}
	}
	return ID
}

func cariUserIDteman(idAkun int, namaAkunDituju string) int {
	var x int = 0
	for i := 1; i <= User[idAkun].TotalTeman; i++ {
		if User[idAkun].Teman[i].nama == namaAkunDituju {
			x = i
		}
	}
	return x
}

func cariGrupID(idAkun int, namaGrup string) int {
	var x int = 0
	for i := 1; i <= User[idAkun].TotalGrup; i++ {
		if User[idAkun].Grup[i].namaGrup == namaGrup {
			x = i
		}
	}
	return x
}

func CekStatusAkun(string1, string2, Code string) bool {
	//mengecek status akun sudah disetujui atau belum disetujui oleh admin. fungsi mengembalikan status
	var status bool = false
	var i, idAkun int = 0, 0
	idAkun = cariUserID(string1)
	for !status && i <= User[0].TotalUser {
		if (Code == "Log") && (User[i].Username == string1) && (User[i].Password == string2) && (User[i].StatusAkun == "Diterima") {
			return true
		} else if Code == "Reg" && User[i].Username == string1 && User[i].Password == string2 {
			return true
		} else if Code == "cariNama" && User[i].Username == string1 {
			return true
		} else if Code == "cariNamaTeman" && User[idAkun].Teman[i].nama == string2 {
			return true
		} else {
			status = false
		}
		i++
	}
	return status
}

func CekStatusGrup(idAkun int, namaGrup string) bool {
	var found bool = false
	for i := 1; i <= User[idAkun].TotalGrup; i++ {
		if User[idAkun].Grup[i].namaGrup == namaGrup {
			found = true
		}
	}
	return found
}

func input() string {
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	return input
}
