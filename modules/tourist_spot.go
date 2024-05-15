package modules

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
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

func MainMenu() {
	clearScreen()
	fmt.Println("=== Selamat Datang di Aplikasi Tempat Wisata ===")
	fmt.Println("1. Daftar")
	fmt.Println("2. Masuk")
	fmt.Println("3. Cari Tempat Wisata")
	fmt.Println("4. Lihat Semua Tempat Wisata")
	fmt.Println("5. Keluar")

	if isLoggedIn {
		fmt.Println("6. Admin")
	}

	var choice int
	fmt.Print("Masukkan pilihan Anda: ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		RegisterUser()
	case 2:
		LoginUser()
	case 3:
		SearchTouristSpots()
	case 4:
		GetAllTouristSpots()
	case 5:
		fmt.Println("Keluar dari program.")
		os.Exit(0)
	case 6:
		if isLoggedIn {
			fmt.Println("Anda memilih Admin.")
			AdminMenu()
		} else {
			fmt.Println("Pilihan tidak valid. Mohon masukkan opsi yang benar.")
		}
	default:
		fmt.Println("Pilihan tidak valid. Mohon masukkan opsi yang benar.")
	}
}

func clearScreen() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls") // wind
	} else {
		cmd = exec.Command("clear") // unix
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
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
