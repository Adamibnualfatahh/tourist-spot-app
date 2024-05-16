package modules

import "fmt"

func AdminMenu() {
	ClearScreen()
	fmt.Println("=== Menu Admin ===")
	fmt.Println("1. Lihat Semua User")
	fmt.Println("2. Tambah Tempat Pariwisata")
	fmt.Println("3. Ubah Tempat Pariwisata")
	fmt.Println("4. Hapus Tempat Pariwisata")
	fmt.Println("5. Tambah Wahana")
	fmt.Println("6. Ubah Wahana")
	fmt.Println("7. Hapus Wahana")
	fmt.Println("8. Tambah Fasilitas")
	fmt.Println("9. Ubah Fasilitas")
	fmt.Println("10. Hapus Fasilitas")
	fmt.Println("11. Kembali ke Menu Utama")

	var choice int
	fmt.Print("Masukkan pilihan Anda: ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		UserList()
	case 2:
		CreateTouristSpot()
	case 3:
		UpdateTouristSpot()
	case 4:
		DeleteTouristSpot()
	case 5:
		CreateAttraction()
	case 6:
		UpdateAttraction()
	case 7:
		DeleteAttraction()
	case 8:
		CreateFacility()
	case 9:
		UpdateFacility()
	case 10:
		DeleteFacility()
	case 11:
		MainMenu()
	default:
		fmt.Println("Pilihan tidak valid. Mohon masukkan opsi yang benar.")
	}
}
