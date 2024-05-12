package main

import (
	"touristSpotApp/modules"
	_ "touristSpotApp/seeds"
)

func main() {
	for {
		modules.MainMenu()
	}
}
