package modules

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	AttractionDB = make(map[int]Attraction)
	attractionId = 0
)

type Attraction struct {
	Id            int
	Name          string
	Description   string
	TouristSpotId int
}

func SetAttractionId(id int) {
	if id > attractionId {
		attractionId = id
	}
}

func CreateAttraction() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan ID tempat wisata: ")
	var touristSpotId int
	fmt.Scanln(&touristSpotId)

	_, ok := TouristSpotDB[touristSpotId]
	if !ok {
		fmt.Println("Tempat wisata dengan ID tersebut tidak ditemukan.")
		return
	}

	fmt.Print("Masukkan nama wahana: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Masukkan deskripsi wahana: ")
	description, _ := reader.ReadString('\n')
	description = strings.TrimSpace(description)

	attractionId++
	newAttraction := Attraction{Id: attractionId, Name: name, Description: description, TouristSpotId: touristSpotId}
	AttractionDB[attractionId] = newAttraction

	fmt.Println("Wahana baru berhasil ditambahkan!")
}

func UpdateAttraction() {
	fmt.Print("Masukkan ID wahana yang ingin diubah: ")
	var id int
	fmt.Scanln(&id)

	_, ok := AttractionDB[id]
	if !ok {
		fmt.Println("Wahana dengan ID tersebut tidak ditemukan.")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan nama wahana: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Masukkan deskripsi wahana: ")
	description, _ := reader.ReadString('\n')
	description = strings.TrimSpace(description)

	newAttraction := Attraction{Id: id, Name: name, Description: description}
	AttractionDB[id] = newAttraction

	fmt.Println("Wahana berhasil diubah!")
}

func DeleteAttraction() {
	fmt.Print("Masukkan ID wahana yang ingin dihapus: ")
	var id int
	fmt.Scanln(&id)

	_, ok := AttractionDB[id]
	if !ok {
		fmt.Println("Wahana dengan ID tersebut tidak ditemukan.")
		return
	}

	delete(AttractionDB, id)
	fmt.Println("Wahana berhasil dihapus!")
}
