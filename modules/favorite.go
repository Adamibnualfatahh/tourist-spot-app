package modules

import "fmt"

var (
	TouristSpotDB   = make(map[int]TouristSpot)
	favoriteId      = 0
	UserFavoritesDB = make(map[int][]Favorite)
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

func FavoriteMenu() {
	fmt.Println("=== Tempat Wisata Favorit ===")
	fmt.Println("1. Lihat Tempat Wisata Favorit")
	fmt.Println("2. Tambah Tempat Wisata Favorit")
	fmt.Println("3. Hapus Tempat Wisata Favorit")
	fmt.Println("4. Kembali ke Menu Utama")
	fmt.Print("Masukkan pilihan Anda: ")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		touristSpots := GetFavoriteTouristSpots(CurrentUser)
		for _, spot := range touristSpots {
			fmt.Printf("ID: %d\n", spot.Id)
			fmt.Printf("Nama: %s\n", spot.Name)
			fmt.Printf("Lokasi: %s\n", spot.Location)
			fmt.Printf("Kategori: %s\n", spot.Category)
			fmt.Printf("Deskripsi: %s\n", spot.Description)
			fmt.Println()
		}
	case 2:
		fmt.Print("Masukkan ID Tempat Wisata: ")
		var touristSpotId int
		fmt.Scanln(&touristSpotId)

		if _, exists := TouristSpotDB[touristSpotId]; exists {
			favorite := AddFavorite(CurrentUser, touristSpotId, "2024-05-04")
			fmt.Printf("Tempat Wisata %s telah ditambahkan ke favorit Anda.\n", TouristSpotDB[favorite.TouristSpotId].Name)
		} else {
			fmt.Println("ID Tempat Wisata tidak valid.")
		}
	case 3:
		fmt.Print("Masukkan ID Tempat Wisata yang ingin dihapus: ")
		var touristSpotId int
		fmt.Scanln(&touristSpotId)

		if _, exists := TouristSpotDB[touristSpotId]; exists {
			favorites := UserFavoritesDB[CurrentUser]
			for i, favorite := range favorites {
				if favorite.TouristSpotId == touristSpotId {
					UserFavoritesDB[CurrentUser] = append(favorites[:i], favorites[i+1:]...)
					fmt.Printf("Tempat Wisata %s telah dihapus dari favorit Anda.\n", TouristSpotDB[favorite.TouristSpotId].Name)
					return
				}
			}
			fmt.Println("Tempat Wisata tidak ditemukan di daftar favorit Anda.")
		} else {
			fmt.Println("ID Tempat Wisata tidak valid.")
		}
	case 4:
		MainMenu()
	default:
		fmt.Println("Pilihan tidak valid. Mohon masukkan opsi yang benar.")
	}
	FavoriteMenu()
}

func AddFavorite(userId int, touristSpotId int, savedDate string) Favorite {
	favoriteId++
	favorite := Favorite{
		Id:            favoriteId,
		UserId:        userId,
		TouristSpotId: touristSpotId,
		SavedDate:     savedDate,
	}

	UserFavoritesDB[userId] = append(UserFavoritesDB[userId], favorite)
	return favorite
}

func GetFavoriteTouristSpots(userId int) []TouristSpot {
	favorites := UserFavoritesDB[userId]
	touristSpots := make([]TouristSpot, 0, len(favorites))

	for _, favorite := range favorites {
		if spot, exists := TouristSpotDB[favorite.TouristSpotId]; exists {
			touristSpots = append(touristSpots, spot)
		}
	}

	return touristSpots
}
