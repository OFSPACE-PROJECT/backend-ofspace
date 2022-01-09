package building

import (
	"context"
	_facility "ofspace_be/features/facility"
	"time"
)

type Core struct {
	Id                 uint
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
	BuildingFacilities []_facility.Core
	ExteriorPhotos     []ExteriorCore
	FloorPhotos        []FloorCore
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

type ExteriorCore struct {
	Id          uint
	BuildingId  uint
	PhotoURL    string
	Description string
}

type FloorCore struct {
	Id          uint
	BuildingId  uint
	PhotoURL    string
	Description string
}

type Business interface {
	CreateBuilding(ctx context.Context, building Core) (Core, error)
	GetAllBuilding(ctx context.Context, complexId uint) ([]Core, error)
	GetAllVerifiedBuilding(ctx context.Context, complexId uint, buildingStatus string) ([]Core, error)
	SearchBuildingByName(ctx context.Context, name string) ([]Core, error)
	GetBuildingById(ctx context.Context, id uint) (Core, error)
	UpdateBuilding(ctx context.Context, id uint) (Core, error)
	RequestBuilding(ctx context.Context, id uint, name string) (Core, error)
	//	exterior and floor photo
	CreateExteriorPhoto(ctx context.Context, buildingId uint, photo ExteriorCore) (ExteriorCore, error)
	UpdateExteriorPhoto(ctx context.Context, photo ExteriorCore) (ExteriorCore, error)
	GetExteriorPhoto(ctx context.Context, BuildingId int) ([]ExteriorCore, error)
	DeleteExteriorPhoto(ctx context.Context, BuildingId int, photoId uint) error

	CreateFloorPhoto(ctx context.Context, buildingId uint, photo FloorCore) (FloorCore, error)
	UpdateFloorPhoto(ctx context.Context, photo FloorCore) (FloorCore, error)
	GetFloorPhoto(ctx context.Context, BuildingId int) ([]FloorCore, error)
	DeleteFloorPhoto(ctx context.Context, BuildingId int, photoId uint) error

	//	for manage facility
	// AddFacilityToBuilding(ctx context.Context, buildingId uint, facilityId uint) ([]BuildingFacilities, error)
	SearchFacility(ctx context.Context, name string) (_facility.Core, error)
	GetFacility(ctx context.Context, facilityId uint) (_facility.Core, error)
	DeleteFacility(ctx context.Context, buildingId uint, facilityId uint) error
}

type Data interface {
	CreateBuilding(ctx context.Context, building Core) (Core, error)
	GetAllBuilding(ctx context.Context, complexId uint) ([]Core, error)
	GetAllVerifiedBuilding(ctx context.Context, complexId uint, buildingStatus string) ([]Core, error)
	SearchBuildingByName(ctx context.Context, name string) ([]Core, error)
	GetBuildingById(ctx context.Context, id uint) (Core, error)
	UpdateBuilding(ctx context.Context, id uint) (Core, error)
	RequestBuilding(ctx context.Context, id uint, name string) (Core, error)
	//	exterior and floor photo
	CreateExteriorPhoto(ctx context.Context, buildingId uint, photo ExteriorCore) (ExteriorCore, error)
	UpdateExteriorPhoto(ctx context.Context, photo ExteriorCore) (ExteriorCore, error)
	GetExteriorPhoto(ctx context.Context, BuildingId int) ([]ExteriorCore, error)
	DeleteExteriorPhoto(ctx context.Context, BuildingId int, photoId uint) error

	CreateFloorPhoto(ctx context.Context, buildingId uint, photo FloorCore) (FloorCore, error)
	UpdateFloorPhoto(ctx context.Context, photo FloorCore) (FloorCore, error)
	GetFloorPhoto(ctx context.Context, BuildingId int) ([]FloorCore, error)
	DeleteFloorPhoto(ctx context.Context, BuildingId int, photoId uint) error

	//	for manage facility
	// AddFacilityToBuilding(ctx context.Context, buildingId uint, facilityId uint) ([]BuildingFacilities, error)
	SearchFacility(ctx context.Context, name string) (_facility.Core, error)
	GetFacility(ctx context.Context, facilityId uint) (_facility.Core, error)
	DeleteFacility(ctx context.Context, buildingId uint, facilityId uint) error
}
