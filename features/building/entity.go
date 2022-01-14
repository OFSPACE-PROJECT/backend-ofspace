package building

import (
	"context"
	"time"
)

type Core struct {
	Id                 uint
	UserId             uint
	ComplexId          uint
	Name               string
	Description        string
	OfficeHours        string
	BuildingSize       string
	AverageFloorSize   string
	YearConstructed    string
	Lifts              string
	Parking            string
	Toilets            string
	BuildingStatus     string
	BuildingFacilities []Facility
	ExteriorPhotos     []ExteriorCore
	FloorPhotos        []FloorCore
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

type Facility struct {
	Id         uint
	BuildingId uint
	Name       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type ExteriorCore struct {
	Id          uint
	BuildingId  uint
	PhotoURL    string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type FloorCore struct {
	Id          uint
	BuildingId  uint
	PhotoURL    string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Business interface {
	CreateBuilding(ctx context.Context, building Core) (Core, error)
	GetAllBuilding(ctx context.Context, complexId uint) ([]Core, error)
	GetAllVerifiedBuilding(ctx context.Context, complexId uint, buildingStatus string) ([]Core, error)
	SearchBuildingByName(ctx context.Context, name string, status string) ([]Core, error)
	GetBuildingById(ctx context.Context, id uint) (Core, error)
	UpdateBuilding(ctx context.Context, building Core) (Core, error)
	//RequestBuilding(ctx context.Context, id uint, name string) (Core, error)
	//	exterior and floor photo
	//CreateExteriorPhoto(ctx context.Context, buildingId uint, photo ExteriorCore) (ExteriorCore, error)
	CreateExteriorPhoto(ctx context.Context, photo ExteriorCore) (ExteriorCore, error)
	UpdateExteriorPhoto(ctx context.Context, photo ExteriorCore) (ExteriorCore, error)
	GetExteriorPhoto(ctx context.Context, BuildingId uint, photoId uint) (ExteriorCore, error)
	GetAllExteriorPhoto(ctx context.Context, BuildingId uint) ([]ExteriorCore, error)
	DeleteExteriorPhoto(ctx context.Context, BuildingId uint, photoId uint) (ExteriorCore, error)

	CreateFloorPhoto(ctx context.Context, photo FloorCore) (FloorCore, error)
	UpdateFloorPhoto(ctx context.Context, photo FloorCore) (FloorCore, error)
	GetAllFloorPhoto(ctx context.Context, BuildingId uint) ([]FloorCore, error)
	GetFloorPhoto(ctx context.Context, BuildingId uint, photoId uint) (FloorCore, error)
	DeleteFloorPhoto(ctx context.Context, BuildingId uint, photoId uint) (FloorCore, error)

	//	for manage facility
	AddFacilityToBuilding(c context.Context, facilityId uint, buildingId uint) (Facility, error)
	GetAllBuildingFacility(c context.Context, buildingId uint) ([]Facility, error)
	GetBuildingFacility(c context.Context, buildingId uint, facilityId uint) (Facility, error)
	DeleteFacility(c context.Context, buildingId uint, facilityId uint) (Facility, error)
}

type Data interface {
	CreateBuilding(ctx context.Context, building Core) (Core, error)
	GetAllBuilding(ctx context.Context, complexId uint) ([]Core, error)
	GetAllVerifiedBuilding(ctx context.Context, complexId uint, buildingStatus string) ([]Core, error)
	SearchBuildingByName(ctx context.Context, name string, status string) ([]Core, error)
	GetBuildingById(ctx context.Context, id uint) (Core, error)
	UpdateBuilding(ctx context.Context, building Core) (Core, error)
	//RequestBuilding(ctx context.Context, id uint, name string) (Core, error)
	//	exterior and floor photo
	//CreateExteriorPhoto(ctx context.Context, buildingId uint, photo ExteriorCore) (ExteriorCore, error)
	CreateExteriorPhoto(ctx context.Context, photo ExteriorCore) (ExteriorCore, error)
	UpdateExteriorPhoto(ctx context.Context, photo ExteriorCore) (ExteriorCore, error)
	GetAllExteriorPhoto(ctx context.Context, BuildingId uint) ([]ExteriorCore, error)
	GetExteriorPhoto(ctx context.Context, buildingId uint, photoId uint) (ExteriorCore, error)
	DeleteExteriorPhoto(ctx context.Context, BuildingId uint, photoId uint) (ExteriorCore, error)

	CreateFloorPhoto(ctx context.Context, photo FloorCore) (FloorCore, error)
	UpdateFloorPhoto(ctx context.Context, photo FloorCore) (FloorCore, error)
	GetAllFloorPhoto(ctx context.Context, BuildingId uint) ([]FloorCore, error)
	GetFloorPhoto(ctx context.Context, BuildingId uint, photoId uint) (FloorCore, error)
	DeleteFloorPhoto(ctx context.Context, BuildingId uint, photoId uint) (FloorCore, error)

	//	for manage facility
	AddFacilityToBuilding(c context.Context, facilityId uint, buildingId uint) (Facility, error)
	GetAllBuildingFacility(c context.Context, buildingId uint) ([]Facility, error)
	GetBuildingFacility(c context.Context, buildingId uint, facilityId uint) (Facility, error)
	DeleteFacility(c context.Context, buildingId uint, facilityId uint) (Facility, error)
}
