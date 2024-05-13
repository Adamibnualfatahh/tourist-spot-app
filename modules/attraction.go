package modules

var (
	AttractionDB = make(map[int]Attraction)
	attractionId = 0
)

type Attraction struct {
	Id            int
	Name          string
	Description   string
	TouristSpotId int
}

func SetAttractionId(id int) {
	if id > attractionId {
		attractionId = id
	}
}
