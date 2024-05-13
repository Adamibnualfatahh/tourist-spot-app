package modules

import "fmt"

func AdminMenu() {
	clearScreen()
	fmt.Println("=== Menu Admin ===")
	fmt.Println("1. Lihat Semua User")
	fmt.Println("2. Tambah Tempat Pariwisata")
	fmt.Println("3. Ubah Tempat Pariwisata")
	fmt.Println("4. Hapus Tempat Pariwisata")
	fmt.Println("5. Kembali ke Menu Utama")

	var choice int
	fmt.Print("Masukkan pilihan Anda: ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		UserList()
	case 2:
		MainMenu()
	default:
		fmt.Println("Pilihan tidak valid. Mohon masukkan opsi yang benar.")
	}
}
