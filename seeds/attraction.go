package seeds

import (
	"touristSpotApp/modules"
	_ "touristSpotApp/modules"
)

func SeedAttractions() {
	attraction1 := modules.Attraction{Id: 1, Name: "Kora Kora", Description: "Permainan yang menantang", TouristSpotId: 1}
	attraction2 := modules.Attraction{Id: 2, Name: "Arung Jeram", Description: "Arung jeram yang seru", TouristSpotId: 2}
	attraction3 := modules.Attraction{Id: 3, Name: "Kembang Api", Description: "Kembang api yang indah", TouristSpotId: 3}

	modules.AttractionDB[1] = attraction1
	modules.AttractionDB[2] = attraction2
	modules.AttractionDB[3] = attraction3

	modules.SetAttractionId(3)
}

func init() {
	SeedAttractions()
}
