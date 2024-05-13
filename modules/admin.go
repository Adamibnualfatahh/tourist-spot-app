package modules

import "fmt"

func UserList() {
	fmt.Println("=== Daftar Semua Pengguna ===")

	if len(UserDB) == 0 {
		fmt.Println("Belum ada pengguna yang terdaftar.")
	} else {
		fmt.Println("Daftar Pengguna:")
		for _, user := range UserDB {
			fmt.Printf("- %d %s\n", user.Id, user.Email)
		}
	}

	backToAdminMenu()
}

func backToAdminMenu() {
	fmt.Println("\nTekan x untuk kembali ke Menu Admin...")
	var input string
	fmt.Scanln(&input)

	if input != "x" {
		fmt.Println("Opsi tidak ditemukan.")
		fmt.Println("\nTekan enter untuk kembali ke Menu Admin...")
		fmt.Scanln()
	}

	AdminMenu()
}
