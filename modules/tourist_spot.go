package modules

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var (
	FavoriteDB    = make(map[int]Favorite)
	touristSpotId = 0
)

type TouristSpot struct {
	Id          int
	Name        string
	Location    string
	Category    string
	Description string
}

func SetTouristSpotId(id int) {
	if id > touristSpotId {
		touristSpotId = id
	}
}

func GetAllTouristSpots() {
	ids := make([]int, 0, len(TouristSpotDB))

	for id := range TouristSpotDB {
		ids = append(ids, id)
	}

	sort.Ints(ids)

	for _, id := range ids {
		touristSpot := TouristSpotDB[id]
		fmt.Println("=========================================")
		fmt.Println("ID:", touristSpot.Id)
		fmt.Println("Nama:", touristSpot.Name)
		fmt.Println("Lokasi:", touristSpot.Location)
		fmt.Println("Kategori:", touristSpot.Category)
		fmt.Println("Deskripsi:", touristSpot.Description)

		fmt.Println("Wahana:")
		for _, attraction := range AttractionDB {
			if attraction.TouristSpotId == touristSpot.Id {
				fmt.Println("  - Nama Atraksi:", attraction.Name)
				fmt.Println("    Deskripsi:", attraction.Description)
			}
		}

		fmt.Println("Fasilitas:")
		for _, facility := range FacilityDB {
			if facility.TouristSpotId == touristSpot.Id {
				fmt.Println("  - Nama Fasilitas:", facility.Name)
				fmt.Println("    Deskripsi:", facility.Description)
			}
		}

		fmt.Println()
	}
}

func SearchTouristSpots() {
	fmt.Println("=== Pencarian Tempat Wisata ===")
	fmt.Println("1. Berdasarkan Nama")
	fmt.Println("2. Berdasarkan Lokasi")
	fmt.Println("3. Berdasarkan Wahana")
	fmt.Println("4. Berdasarkan Fasilitas")
	fmt.Println("5. Kembali ke Menu Utama")

	var choice int
	fmt.Print("Masukkan pilihan Anda: ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		SearchByName()
	case 2:
		SearchByLocation()
	case 3:
		SearchByAttraction()
	case 4:
		SearchByFacility()
	case 5:
		fmt.Println("Kembali ke Menu Utama.")
		MainMenu()
	default:
		fmt.Println("Pilihan tidak valid. Mohon masukkan opsi yang benar.")
		SearchTouristSpots()
	}
}

func SearchByName() {
	fmt.Print("Masukkan nama tempat wisata yang ingin dicari: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		query := scanner.Text()

		found := false
		for _, touristSpot := range TouristSpotDB {
			if strings.Contains(strings.ToLower(touristSpot.Name), strings.ToLower(query)) {
				fmt.Println("=========================================")
				fmt.Println("ID:", touristSpot.Id)
				fmt.Println("Nama:", touristSpot.Name)
				fmt.Println("Lokasi:", touristSpot.Location)
				fmt.Println("Kategori:", touristSpot.Category)
				fmt.Println("Deskripsi:", touristSpot.Description)

				fmt.Println("Wahana:")
				for _, attraction := range AttractionDB {
					if attraction.TouristSpotId == touristSpot.Id {
						fmt.Println("  - Nama Atraksi:", attraction.Name)
						fmt.Println("    Deskripsi:", attraction.Description)
					}
				}

				fmt.Println("Fasilitas:")
				for _, facility := range FacilityDB {
					if facility.TouristSpotId == touristSpot.Id {
						fmt.Println("  - Nama Fasilitas:", facility.Name)
						fmt.Println("    Deskripsi:", facility.Description)
					}
				}
				found = true
			}
		}

		if !found {
			fmt.Println("Tempat wisata dengan nama tersebut tidak ditemukan.")
		}
	}

	fmt.Println()
	MainMenu()
}

func SearchByLocation() {
	fmt.Print("Masukkan lokasi tempat wisata yang ingin dicari: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		query := scanner.Text()

		found := false
		for _, touristSpot := range TouristSpotDB {
			if strings.Contains(strings.ToLower(touristSpot.Location), strings.ToLower(query)) {
				fmt.Println("=========================================")
				fmt.Println("ID:", touristSpot.Id)
				fmt.Println("Nama:", touristSpot.Name)
				fmt.Println("Lokasi:", touristSpot.Location)
				fmt.Println("Kategori:", touristSpot.Category)
				fmt.Println("Deskripsi:", touristSpot.Description)

				fmt.Println("Wahana:")
				for _, attraction := range AttractionDB {
					if attraction.TouristSpotId == touristSpot.Id {
						fmt.Println("  - Nama Atraksi:", attraction.Name)
						fmt.Println("    Deskripsi:", attraction.Description)
					}
				}

				fmt.Println("Fasilitas:")
				for _, facility := range FacilityDB {
					if facility.TouristSpotId == touristSpot.Id {
						fmt.Println("  - Nama Fasilitas:", facility.Name)
						fmt.Println("    Deskripsi:", facility.Description)
					}
				}
				found = true
			}
		}

		if !found {
			fmt.Println("Tempat wisata dengan lokasi tersebut tidak ditemukan.")
		}
	}

	fmt.Println()
	MainMenu()
}

func SearchByAttraction() {
	fmt.Print("Masukkan wahana wisata yang ingin dicari: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		query := scanner.Text()

		found := false
		for _, attraction := range AttractionDB {
			if strings.Contains(strings.ToLower(attraction.Name), strings.ToLower(query)) {
				touristSpot := TouristSpotDB[attraction.TouristSpotId]
				fmt.Println("=========================================")
				fmt.Println("ID:", touristSpot.Id)
				fmt.Println("Nama:", touristSpot.Name)
				fmt.Println("Lokasi:", touristSpot.Location)
				fmt.Println("Kategori:", touristSpot.Category)
				fmt.Println("Deskripsi:", touristSpot.Description)

				fmt.Println("Wahana:")
				for _, attraction := range AttractionDB {
					if attraction.TouristSpotId == touristSpot.Id {
						fmt.Println("  - Nama Atraksi:", attraction.Name)
						fmt.Println("    Deskripsi:", attraction.Description)
					}
				}

				fmt.Println("Fasilitas:")
				for _, facility := range FacilityDB {
					if facility.TouristSpotId == touristSpot.Id {
						fmt.Println("  - Nama Fasilitas:", facility.Name)
						fmt.Println("    Deskripsi:", facility.Description)
					}
				}
				found = true
			}
		}

		if !found {
			fmt.Println("Tempat wisata dengan wahana tersebut tidak ditemukan.")
		}
	}
}

func SearchByFacility() {
	fmt.Print("Masukkan fasilitas wisata yang ingin dicari: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		query := scanner.Text()

		found := false
		for _, facility := range FacilityDB {
			if strings.Contains(strings.ToLower(facility.Name), strings.ToLower(query)) {
				touristSpot := TouristSpotDB[facility.TouristSpotId]
				fmt.Println("=========================================")
				fmt.Println("ID:", touristSpot.Id)
				fmt.Println("Nama:", touristSpot.Name)
				fmt.Println("Lokasi:", touristSpot.Location)
				fmt.Println("Kategori:", touristSpot.Category)
				fmt.Println("Deskripsi:", touristSpot.Description)

				fmt.Println("Wahana:")
				for _, attraction := range AttractionDB {
					if attraction.TouristSpotId == touristSpot.Id {
						fmt.Println("  - Nama Atraksi:", attraction.Name)
						fmt.Println("    Deskripsi:", attraction.Description)
					}
				}

				fmt.Println("Fasilitas:")
				for _, facility := range FacilityDB {
					if facility.TouristSpotId == touristSpot.Id {
						fmt.Println("  - Nama Fasilitas:", facility.Name)
						fmt.Println("    Deskripsi:", facility.Description)
					}
				}
				found = true
			}
		}

		if !found {
			fmt.Println("Tempat wisata dengan fasilitas tersebut tidak ditemukan.")
		}
	}
}

func CreateTouristSpot() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan nama tempat wisata: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Masukkan lokasi tempat wisata: ")
	location, _ := reader.ReadString('\n')
	location = strings.TrimSpace(location)

	fmt.Print("Masukkan kategori tempat wisata: ")
	category, _ := reader.ReadString('\n')
	category = strings.TrimSpace(category)

	fmt.Print("Masukkan deskripsi tempat wisata: ")
	description, _ := reader.ReadString('\n')
	description = strings.TrimSpace(description)

	touristSpotId++
	newTouristSpot := TouristSpot{Id: touristSpotId, Name: name, Location: location, Category: category, Description: description}
	TouristSpotDB[touristSpotId] = newTouristSpot

	fmt.Println("Tempat wisata baru berhasil ditambahkan!")
}

func UpdateTouristSpot() {
	fmt.Print("Masukkan ID tempat wisata yang ingin diubah: ")
	var id int
	fmt.Scanln(&id)

	_, ok := TouristSpotDB[id]
	if !ok {
		fmt.Println("Tempat wisata dengan ID tersebut tidak ditemukan.")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan nama tempat wisata: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Masukkan lokasi tempat wisata: ")
	location, _ := reader.ReadString('\n')
	location = strings.TrimSpace(location)

	fmt.Print("Masukkan kategori tempat wisata: ")
	category, _ := reader.ReadString('\n')
	category = strings.TrimSpace(category)

	fmt.Print("Masukkan deskripsi tempat wisata: ")
	description, _ := reader.ReadString('\n')
	description = strings.TrimSpace(description)

	newTouristSpot := TouristSpot{Id: id, Name: name, Location: location, Category: category, Description: description}
	TouristSpotDB[id] = newTouristSpot

	fmt.Println("Tempat wisata berhasil diubah!")
}

func DeleteTouristSpot() {
	fmt.Print("Masukkan ID tempat wisata yang ingin dihapus: ")
	var id int
	fmt.Scanln(&id)

	_, ok := TouristSpotDB[id]
	if !ok {
		fmt.Println("Tempat wisata dengan ID tersebut tidak ditemukan.")
		return
	}

	delete(TouristSpotDB, id)
	fmt.Println("Tempat wisata berhasil dihapus!")
}
