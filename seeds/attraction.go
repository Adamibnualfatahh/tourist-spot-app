package seeds

import (
	"touristSpotApp/modules"
	_ "touristSpotApp/modules"
)

func SeedAttractions() {
	attraction1 := modules.Attraction{Id: 1, Name: "Kora Kora", Description: "Permainan yang menantang", TouristSpotId: 1}
	attraction2 := modules.Attraction{Id: 2, Name: "Arung Jeram", Description: "Arung jeram yang seru", TouristSpotId: 2}
	attraction3 := modules.Attraction{Id: 3, Name: "Kembang Api", Description: "Kembang api yang indah", TouristSpotId: 3}
	attraction4 := modules.Attraction{Id: 4, Name: "Tari Kecak", Description: "Tarian ritual kuno Bali", TouristSpotId: 4}
	attraction5 := modules.Attraction{Id: 5, Name: "Banana Boat", Description: "kapal karet yang ditarik dengan speedboat dalam kecepatan tinggi", TouristSpotId: 5}
	attraction6 := modules.Attraction{Id: 6, Name: "Snorkeling", Description: "Menyelam dilaut dengan pemandangan trumbu karang yang indah", TouristSpotId: 6}
	attraction7 := modules.Attraction{Id: 7, Name: "Selancar", Description: "berselancar diombak yang memacu adrenalin", TouristSpotId: 7}

	modules.AttractionDB[1] = attraction1
	modules.AttractionDB[2] = attraction2
	modules.AttractionDB[3] = attraction3
	modules.AttractionDB[4] = attraction4
	modules.AttractionDB[5] = attraction5
	modules.AttractionDB[6] = attraction6
	modules.AttractionDB[7] = attraction7

	modules.SetAttractionId(7)
}

func init() {
	SeedAttractions()
}
