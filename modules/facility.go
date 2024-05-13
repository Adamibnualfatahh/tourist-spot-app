package modules

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
