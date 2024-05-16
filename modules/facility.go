package modules

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	FacilityDB = make(map[int]Facility)
	facilityId = 0
)

type Facility struct {
	ID            int
	Name          string
	Description   string
	TouristSpotId int
}

func SetFacilityId(id int) {
	if id > facilityId {
		facilityId = id
	}
}

func CreateFacility() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan ID tempat wisata: ")
	var touristSpotId int
	fmt.Scanln(&touristSpotId)

	_, ok := TouristSpotDB[touristSpotId]
	if !ok {
		fmt.Println("Tempat wisata dengan ID tersebut tidak ditemukan.")
		return
	}

	fmt.Print("Masukkan nama fasilitas: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Masukkan deskripsi fasilitas: ")
	description, _ := reader.ReadString('\n')
	description = strings.TrimSpace(description)

	facilityId++
	newFacility := Facility{ID: facilityId, Name: name, Description: description, TouristSpotId: touristSpotId}
	FacilityDB[facilityId] = newFacility

	fmt.Println("Fasilitas baru berhasil ditambahkan!")
}

func UpdateFacility() {
	fmt.Print("Masukkan ID fasilitas yang ingin diubah: ")
	var id int
	fmt.Scanln(&id)

	_, ok := FacilityDB[id]
	if !ok {
		fmt.Println("Fasilitas dengan ID tersebut tidak ditemukan.")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan nama fasilitas: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Masukkan deskripsi fasilitas: ")
	description, _ := reader.ReadString('\n')
	description = strings.TrimSpace(description)

	newFacility := Facility{ID: id, Name: name, Description: description}
	FacilityDB[id] = newFacility

	fmt.Println("Fasilitas berhasil diubah!")
}

func DeleteFacility() {
	fmt.Print("Masukkan ID fasilitas yang ingin dihapus: ")
	var id int
	fmt.Scanln(&id)

	_, ok := FacilityDB[id]
	if !ok {
		fmt.Println("Fasilitas dengan ID tersebut tidak ditemukan.")
		return
	}

	delete(FacilityDB, id)
	fmt.Println("Fasilitas berhasil dihapus!")
}
