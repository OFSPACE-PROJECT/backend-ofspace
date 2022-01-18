package unit

import (
	"context"
	"time"
)

type Core struct {
	Id             uint
	UserId         uint
	BuildingId     uint
	Description    string
	UnitType       string
	Price          float32
	TotalUnit      int
	RemainingUnit  int
	UnitFacilities []Facility
	InteriorPhoto  []InteriorCore
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type Facility struct {
	Id   uint
	Name string
}

type InteriorCore struct {
	Id          uint
	UnitId      uint
	PhotoURL    string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Business interface {
	CreateUnit(ctx context.Context, unit Core) (Core, error)
	GetAllUnit(ctx context.Context, buildingId uint) ([]Core, error)
	GetUnitByType(ctx context.Context, buildingId uint, unitType string) (Core, error)
	GetUnitById(ctx context.Context, buildingId uint, facilityId uint) (Core, error)
	UpdateUnit(ctx context.Context, unit Core) (Core, error)
	//DeleteUnit(ctx context.Context, id uint) (Core, error)

	//	exterior and floor photo
	CreateInteriorPhoto(ctx context.Context, photo InteriorCore) (InteriorCore, error)
	UpdateInteriorPhoto(ctx context.Context, photo InteriorCore) (InteriorCore, error)
	GetAllInteriorPhoto(ctx context.Context, UnitId uint) ([]InteriorCore, error)
	GetInteriorPhoto(ctx context.Context, UnitId uint, photoId uint) (InteriorCore, error)
	DeleteInteriorPhoto(ctx context.Context, UnitId uint, photoId uint) (InteriorCore, error)

	//	for manage facility
	// AddFacilityToUnit(ctx context.Context, unitId uint, facilityId uint) ([]UnitFacilities, error)
	AddFacilityToUnit(c context.Context, unitId uint, facilityId uint) (Facility, error)
	GetAllUnitFacility(c context.Context, unitId uint) (Core, error)
	GetUnitFacility(ctx context.Context, unitId uint, facilityId uint) (Facility, error)
	DeleteUnitFacility(ctx context.Context, unitId uint, facilityId uint) (Core, error)
}

type Data interface {
	CreateUnit(ctx context.Context, unit Core) (Core, error)
	GetAllUnit(ctx context.Context, complexId uint) ([]Core, error)
	GetUnitById(ctx context.Context, buildingId uint, facilityId uint) (Core, error)
	GetUnitByType(ctx context.Context, buildingId uint, unitType string) (Core, error)
	UpdateUnit(ctx context.Context, unit Core) (Core, error)
	//DeleteUnit(ctx context.Context, id uint) (Core, error)
	//	exterior and floor photo
	CreateInteriorPhoto(ctx context.Context, photo InteriorCore) (InteriorCore, error)
	UpdateInteriorPhoto(ctx context.Context, photo InteriorCore) (InteriorCore, error)
	GetInteriorPhoto(ctx context.Context, UnitId uint, photoId uint) (InteriorCore, error)
	GetAllInteriorPhoto(ctx context.Context, UnitId uint) ([]InteriorCore, error)
	DeleteInteriorPhoto(ctx context.Context, UnitId uint, photoId uint) (InteriorCore, error)

	//	for manage facility
	// AddFacilityToUnit(ctx context.Context, unitId uint, facilityId uint) ([]UnitFacilities, error)
	AddFacilityToUnit(c context.Context, unitId uint, facilityId uint) (Facility, error)
	GetAllUnitFacility(c context.Context, unitId uint) (Core, error)
	GetUnitFacility(ctx context.Context, unitId uint, facilityId uint) (Facility, error)
	DeleteUnitFacility(ctx context.Context, unitId uint, facilityId uint) (Core, error)
}
