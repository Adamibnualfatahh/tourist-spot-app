package modules

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func MainMenu() {
	fmt.Println("=== Selamat Datang di Aplikasi Tempat Wisata ===")
	fmt.Println("1. Daftar")

	if !isLoggedIn {
		fmt.Println("2. Masuk")
	}

	fmt.Println("3. Cari Tempat Wisata")
	fmt.Println("4. Lihat Semua Tempat Wisata")
	fmt.Println("5. Keluar")

	if isLoggedIn {
		fmt.Println("6. Admin")
		fmt.Println("7. Tempat Wisata Favorit")
	}

	var choice int
	fmt.Print("Masukkan pilihan Anda: ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		RegisterUser()
	case 2:
		if isLoggedIn {
			fmt.Println("Anda sudah masuk.")
			break
		}
		LoginUser()
	case 3:
		SearchTouristSpots()
	case 4:
		GetAllTouristSpots()
	case 5:
		fmt.Println("Keluar dari program.")
		os.Exit(0)
	case 6:
		if isLoggedIn {
			fmt.Println("Anda memilih Admin.")
			AdminMenu()
		} else {
			fmt.Println("Pilihan tidak valid. Mohon masukkan opsi yang benar.")
		}
	case 7:
		if isLoggedIn {
			fmt.Println("Anda memilih Tempat Wisata Favorit.")
			FavoriteMenu()
		} else {
			fmt.Println("Pilihan tidak valid. Mohon masukkan opsi yang benar.")
		}
	default:
		fmt.Println("Pilihan tidak valid. Mohon masukkan opsi yang benar.")
	}
}

func ClearScreen() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls") // wind
	} else {
		cmd = exec.Command("clear") // unix
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}
