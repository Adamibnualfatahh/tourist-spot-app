package modules

var (
	TouristSpotDB = make(map[int]TouristSpot)
	favoriteId    = 0
)

type Favorite struct {
	Id            int
	UserId        int
	TouristSpotId int
	SavedDate     string
}

func SetFavoriteId(id int) {
	if id > favoriteId {
		favoriteId = id
	}
}
