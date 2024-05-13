package seeds

import (
	"touristSpotApp/modules"
	_ "touristSpotApp/modules"
)

type Favorite struct {
	Id            int
	UserId        int
	TouristSpotId string
	SavedDate     string
}

func SeedFavorites() {
	favorite1 := Favorite{Id: 1, UserId: 2, TouristSpotId: "1", SavedDate: "2024-05-01"}
	favorite2 := Favorite{Id: 2, UserId: 2, TouristSpotId: "2", SavedDate: "2024-05-02"}
	favorite3 := Favorite{Id: 3, UserId: 2, TouristSpotId: "3", SavedDate: "2024-05-03"}

	modules.FavoriteDB[1] = favorite1
	modules.FavoriteDB[2] = favorite2
	modules.FavoriteDB[3] = favorite3

	modules.SetFavoriteId(3)
}

func init() {
	SeedFavorites()
}
