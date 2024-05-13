package seeds

import (
	"touristSpotApp/modules"
	_ "touristSpotApp/modules"
)

type User struct {
	Id       int
	Name     string
	Email    string
	Password string
}

func SeedUsers() {
	admin := User{Id: 1, Name: "admin", Email: "admin@gmail.com", Password: "admin"}
	user := User{Id: 2, Name: "user", Email: "user@gmail.com", Password: "user"}
	adam := User{Id: 3, Name: "adam", Email: "adam", Password: "adam"}

	modules.UserDB["admin@gmail.com"] = modules.User(admin)
	modules.UserDB["user@gmail.com"] = modules.User(user)
	modules.UserDB["adam"] = modules.User(adam)

	modules.SetUserId(3)
}

func init() {
	SeedUsers()
}
