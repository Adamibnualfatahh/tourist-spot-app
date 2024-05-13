package seeds

import (
	"touristSpotApp/modules"
	_ "touristSpotApp/modules"
)

type TouristSpot struct {
	ID          int
	Name        string
	Location    string
	Category    string
	Description string
}

func SeedTouristSpots() {
	touristSpots1 = TouristSpot{ID: 1, Name: "Pantai Bali", Location: "Bali", Category: "Pantai", Description: "pantai yang cantik dan bersih."}
	touristSpots2 = TouristSpot{ID: 2, Name: "Musium Jakarta", Location: "Jakarta", Category: "Musium", Description: "Musium yang menyimpan sejarah jakarta."}
	touristSpots3 = TouristSpot{ID: 3, Name: "Pantai Kelingking", Location: "Nusa Peninda", Category: "Pantai", Description: "Pantai yang indah dan bersih."}
	touristSpots3 = TouristSpot{ID: 3, Name: "Gunung Bromo", Location: "Jawa Timur", Category: "Gunung", Description: "Gunung berapi aktif yang terkenal dengan pemandangannya yang dramatis."}
	touristSpots4 = TouristSpot{ID: 4, Name: "Taman Safari", Location: "Bogor", Category: "Kebun Binatang", Description: "Taman Safari dengan berbagai jenis hewan dan atraksi yang menarik untuk keluarga."}
	touristSpots5 = TouristSpot{ID: 5, Name: "Danau Toba", Location: "Sumatra Utara", Category: "Danau", Description: "Danau vulkanik terbesar di Indonesia dengan pemandangan alam yang memukau."}
	touristSpots6 = TouristSpot{ID: 6, Name: "Candi Borobudur", Location: "Magelang", Category: "Candi", Description: "Candi Budha terbesar di dunia yang merupakan situs Warisan Dunia UNESCO."}
	touristSpots7 = TouristSpot{ID: 7, Name: "Pulau Komodo", Location: "Nusa tenggara ", Category: "Pantai", Description: "Pantai yang sangat indah yang terkenal akan pasir merah dan habitat asli hewan komodo."}
	touristSpots8 = TouristSpot{ID: 8, Name: "Ubud", Location: "Bali", Category: "Kota Seni", Description: "Pusat seni dan budaya Bali, dikenal dengan galeri seni dan pertunjukan tari tradisional."}
	touristSpots9 = TouristSpot{ID: 9, Name: "Raja Ampat", Location: "Papua Barat", Category: "Taman Laut", Description: "Salah satu destinasi selam terbaik di dunia dengan keanekaragaman hayati yang tinggi."}
	touristSpots10 = TouristSpot{ID: 10, Name: "Bunaken", Location: "Sulawesi Utara", Category: "Taman Laut", Description: "Taman Nasional Laut yang terkenal dengan terumbu karang yang berwarna-warni dan kehidupan lautnya yang kaya."}

	modules.TouristSpotDB[1] = touristSpots1
	modules.TouristSpotDB[2] = touristSpots2
	modules.TouristSpotDB[3] = touristSpots3
	modules.TouristSpotDB[4] = touristSpots4
	modules.TouristSpotDB[5] = touristSpots5
	modules.TouristSpotDB[6] = touristSpots6
	modules.TouristSpotDB[7] = touristSpots7
	modules.TouristSpotDB[8] = touristSpots8
	modules.TouristSpotDB[9] = touristSpots9
	modules.TouristSpotDB[10] = touristSpots10

	modules.SetTouristSpotId(10)
}

func init() {
	SeedTouristSpots()
}
