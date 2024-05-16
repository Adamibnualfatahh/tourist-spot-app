package main

import (
	"fmt"
	"touristSpotApp/modules"
	_ "touristSpotApp/seeds"
)

func main() {
	for {
		fmt.Println("===================")
		fmt.Println("Aplikasi Tempat Wisata")
		fmt.Println("===================")
		fmt.Println("Tekan [Enter] untuk melanjutkan")
		fmt.Println("Tekan [q] untuk keluar")
		fmt.Println("===================")

		var input string
		fmt.Scanln(&input)

		if input == "q" {
			break
		} else if input == "" {
			modules.ClearScreen()
			modules.MainMenu()
		}
	}
}
