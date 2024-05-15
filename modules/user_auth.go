package modules

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	UserDB     = make(map[string]User)
	isLoggedIn bool
	userId     = 0
)

type User struct {
	Id       int
	Name     string
	Email    string
	Password string
}

func SetUserId(id int) {
	if id > userId {
		userId = id
	}
}

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

	userId++
	newUser := User{Id: userId, Name: name, Email: email, Password: password}
	UserDB[email] = newUser

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
		SetLoggedInStatus(true)
	} else {
		fmt.Println("Email atau kata sandi salah. Silakan coba lagi.")
	}
}

func ValidateLogin(email, password string) (bool, string) {
	user, ok := UserDB[email]
	if !ok {
		return false, ""
	}

	if user.Password != password {
		return false, ""
	}

	return true, user.Name
}

func SetLoggedInStatus(status bool) {
	isLoggedIn = status
}
