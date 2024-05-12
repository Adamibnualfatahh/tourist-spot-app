package modules

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var userDB = make(map[string]string)

func RegisterUser() {
	fmt.Println("=== Daftar Pengguna Baru ===")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan nama: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Masukkan email: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	fmt.Print("Masukkan kata sandi: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	userDB[email] = password

	fmt.Println("Pengguna baru berhasil didaftarkan!")
}

func LoginUser() {
	fmt.Println("=== Masuk ===")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan email: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	fmt.Print("Masukkan kata sandi: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	isLoggedIn, name := ValidateLogin(email, password)

	if isLoggedIn {
		fmt.Println("Selamat datang, " + name + "!")
	} else {
		fmt.Println("Email atau kata sandi salah. Silakan coba lagi.")

	}
}

func ValidateLogin(email, password string) (bool, string) {
	savedPassword, ok := userDB[email]
	if !ok {
		return false, ""
	}

	if savedPassword != password {
		return false, ""
	}

	return true, email
}
