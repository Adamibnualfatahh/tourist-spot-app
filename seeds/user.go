package seeds

import (
	"touristSpotApp/modules"
	_ "touristSpotApp/modules"
)

type User struct {
	Name     string
	Email    string
	Password string
}

func SeedUsers() {
	admin := User{Name: "admin", Email: "admin@gmail.com", Password: "admin"}
	user := User{Name: "user", Email: "user@gmail.com", Password: "user"}
	adam := User{Name: "adam", Email: "adam", Password: "adam"}

	modules.UserDB["admin@gmail.com"] = modules.User(admin)
	modules.UserDB["user@gmail.com"] = modules.User(user)
	modules.UserDB["adam"] = modules.User(adam)
}

func init() {
	SeedUsers()
}
