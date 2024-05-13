package seeds

import (
	"touristSpotApp/modules"
	_ "touristSpotApp/modules"
)

func SeedFacilities() {
	facility1 := modules.Facility{ID: 1, Name: "Kamar Mandi", Description: "Kamar mandi yang nyaman", TouristSpotId: 1}
	facility2 := modules.Facility{ID: 2, Name: "Kolam Renang", Description: "Kolam renang yang luas", TouristSpotId: 2}
	facility3 := modules.Facility{ID: 3, Name: "Tanaman Hias", Description: "Tanaman hias yang indah", TouristSpotId: 3}

	modules.FacilityDB[1] = facility1
	modules.FacilityDB[2] = facility2
	modules.FacilityDB[3] = facility3

	modules.SetFacilityId(3)
}

func init() {
	SeedFacilities()
}
