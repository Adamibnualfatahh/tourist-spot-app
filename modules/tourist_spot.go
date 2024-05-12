package modules

import (
	"fmt"
	"os"
)

func MainMenu() {
	fmt.Println("=== Selamat Datang di Aplikasi Tempat Wisata ===")
	fmt.Println("1. Daftar")
	fmt.Println("2. Masuk")
	fmt.Println("3. Cari Tempat Wisata")
	fmt.Println("4. Lihat Semua Tempat Wisata")
	fmt.Println("5. Keluar")

	var choice int
	fmt.Print("Masukkan pilihan Anda: ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		RegisterUser()
	case 2:
		LoginUser()
	case 3:
		fmt.Println("Anda memilih Cari Tempat Wisata.")
	case 4:
		fmt.Println("Anda memilih Lihat Semua Tempat Wisata.")
	case 5:
		fmt.Println("Keluar dari program.")
		os.Exit(0)
	default:
		fmt.Println("Pilihan tidak valid. Mohon masukkan opsi yang benar.")
	}
}
