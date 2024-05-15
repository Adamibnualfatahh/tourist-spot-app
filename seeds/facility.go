package seeds

import (
	"touristSpotApp/modules"
	_ "touristSpotApp/modules"
)

func SeedFacilities() {
	facility1 := modules.Facility{ID: 1, Name: "Kamar Mandi", Description: "Kamar mandi yang nyaman", TouristSpotId: 1}
	facility2 := modules.Facility{ID: 2, Name: "Kolam Renang", Description: "Kolam renang yang luas", TouristSpotId: 2}
	facility3 := modules.Facility{ID: 3, Name: "Tanaman Hias", Description: "Tanaman hias yang indah", TouristSpotId: 3}
	facility4 := modules.Facility{ID: 4, Name: "Restoran", Description: "Restoran dengan makanan lezat", TouristSpotId: 4}
	facility5 := modules.Facility{ID: 5, Name: "Toko Souvenir", Description: "Toko souvenir dengan berbagai pilihan", TouristSpotId: 5}
	facility6 := modules.Facility{ID: 6, Name: "Pusat Informasi", Description: "Pusat informasi wisata", TouristSpotId: 6}
	facility7 := modules.Facility{ID: 7, Name: "Area Parkir", Description: "Area parkir yang luas", TouristSpotId: 7}
	facility8 := modules.Facility{ID: 8, Name: "Area Bermain Anak", Description: "Area bermain untuk anak-anak", TouristSpotId: 8}
	facility9 := modules.Facility{ID: 9, Name: "Mushola", Description: "Tempat ibadah yang nyaman", TouristSpotId: 9}
	facility10 := modules.Facility{ID: 10, Name: "ATM Center", Description: "Pusat ATM untuk kemudahan transaksi", TouristSpotId: 10}
	facility11 := modules.Facility{ID: 11, Name: "Kolam Renang", Description: "Kolam renang yang luas", TouristSpotId: 1}

	modules.FacilityDB[1] = facility1
	modules.FacilityDB[2] = facility2
	modules.FacilityDB[3] = facility3
	modules.FacilityDB[4] = facility4
	modules.FacilityDB[5] = facility5
	modules.FacilityDB[6] = facility6
	modules.FacilityDB[7] = facility7
	modules.FacilityDB[8] = facility8
	modules.FacilityDB[9] = facility9
	modules.FacilityDB[10] = facility10
	modules.FacilityDB[11] = facility11

	modules.SetFacilityId(11)
}

func init() {
	SeedFacilities()
}
