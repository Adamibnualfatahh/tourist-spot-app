package seeds

import (
	"touristSpotApp/modules"
	_ "touristSpotApp/modules"
)

func SeedTouristSpots() {
	touristSpot1 := modules.TouristSpot{Id: 1, Name: "Pantai Bali", Location: "Bali", Category: "Pantai", Description: "Pantai yang cantik dan bersih."}
	touristSpot2 := modules.TouristSpot{Id: 2, Name: "Museum Jakarta", Location: "Jakarta", Category: "Museum", Description: "Museum yang menyimpan sejarah Jakarta."}
	touristSpot3 := modules.TouristSpot{Id: 3, Name: "Pantai Kelingking", Location: "Nusa Penida", Category: "Pantai", Description: "Pantai yang indah dan bersih."}
	touristSpot4 := modules.TouristSpot{Id: 4, Name: "Gunung Bromo", Location: "Jawa Timur", Category: "Gunung", Description: "Gunung berapi aktif yang terkenal dengan pemandangan yang dramatis."}
	touristSpot5 := modules.TouristSpot{Id: 5, Name: "Danau Toba", Location: "Sumatra Utara", Category: "Danau", Description: "Danau vulkanik terbesar di Indonesia dengan pemandangan alam yang memukau."}
	touristSpot6 := modules.TouristSpot{Id: 6, Name: "Candi Borobudur", Location: "Magelang", Category: "Candi", Description: "Candi Budha terbesar di dunia yang merupakan situs Warisan Dunia UNESCO."}
	touristSpot7 := modules.TouristSpot{Id: 7, Name: "Pulau Komodo", Location: "Nusa Tenggara", Category: "Pulau", Description: "Pulau yang sangat indah yang terkenal akan pasir merah dan habitat asli hewan komodo."}
	touristSpot8 := modules.TouristSpot{Id: 8, Name: "Ubud", Location: "Bali", Category: "Kota Seni", Description: "Pusat seni dan budaya Bali, dikenal dengan galeri seni dan pertunjukan tari tradisional."}
	touristSpot9 := modules.TouristSpot{Id: 9, Name: "Raja Ampat", Location: "Papua Barat", Category: "Taman Laut", Description: "Salah satu destinasi selam terbaik di dunia dengan keanekaragaman hayati yang tinggi."}
	touristSpot10 := modules.TouristSpot{Id: 10, Name: "Bunaken", Location: "Sulawesi Utara", Category: "Taman Laut", Description: "Taman Nasional Laut yang terkenal dengan terumbu karang yang berwarna-warni dan kehidupan lautnya yang kaya."}

	modules.TouristSpotDB[1] = touristSpot1
	modules.TouristSpotDB[2] = touristSpot2
	modules.TouristSpotDB[3] = touristSpot3
	modules.TouristSpotDB[4] = touristSpot4
	modules.TouristSpotDB[5] = touristSpot5
	modules.TouristSpotDB[6] = touristSpot6
	modules.TouristSpotDB[7] = touristSpot7
	modules.TouristSpotDB[8] = touristSpot8
	modules.TouristSpotDB[9] = touristSpot9
	modules.TouristSpotDB[10] = touristSpot10

	modules.SetTouristSpotId(10)
}

func init() {
	SeedTouristSpots()
}
