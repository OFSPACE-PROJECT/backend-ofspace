package unit

import (
	"context"
	_facility "ofspace_be/features/facility"
	"time"
)

type Core struct {
	Id             uint
	ComplexId      uint
	MainPhoto      string
	Type           string
	InitPrice      float32
	TotalUnit      int
	RemainingUnit  int
	UnitFacilities []_facility.Core
	InteriorPhotos []InteriorPhotos
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type InteriorPhotos struct {
	Id          uint
	UnitId      uint
	PhotoURL    string
	Description string
}

type Business interface {
	CreateUnit(ctx context.Context, unit Core) (Core, error)
	GetAllUnit(ctx context.Context, buildingId uint) ([]Core, error)
	GetUnitByType(ctx context.Context, unitType string) (Core, error)
	UpdateUnit(ctx context.Context, id uint, unit Core) (Core, error)
	DeleteUnit(ctx context.Context, id uint) error
	//	exterior and floor photo
	CreateInteriorPhoto(ctx context.Context, unitId uint, photo InteriorPhotos) (InteriorPhotos, error)
	UpdateInteriorPhoto(ctx context.Context, photo InteriorPhotos) (InteriorPhotos, error)
	GetInteriorPhoto(ctx context.Context, UnitId int) ([]InteriorPhotos, error)
	DeleteInteriorPhoto(ctx context.Context, UnitId int, photoId uint) error

	//	for manage facility
	// AddFacilityToUnit(ctx context.Context, unitId uint, facilityId uint) ([]UnitFacilities, error)
	SearchFacility(ctx context.Context, name string) (_facility.Core, error)
	GetFacility(ctx context.Context, facilityId uint) (_facility.Core, error)
	DeleteFacility(ctx context.Context, unitId uint, facilityId uint) error
}

type Data interface {
	CreateUnit(ctx context.Context, unit Core) (Core, error)
	GetAllUnit(ctx context.Context, complexId uint) ([]Core, error)
	GetAllVerifiedUnit(ctx context.Context, complexId uint, unitStatus string) ([]Core, error)
	SearchUnitByName(ctx context.Context, name string) ([]Core, error)
	GetUnitById(ctx context.Context, id uint) (Core, error)
	UpdateUnit(ctx context.Context, id uint) (Core, error)
	RequestUnit(ctx context.Context, id uint, name string) (Core, error)
	//	exterior and floor photo
	CreateInteriorPhoto(ctx context.Context, unitId uint, photo InteriorPhotos) (InteriorPhotos, error)
	UpdateInteriorPhoto(ctx context.Context, photo InteriorPhotos) (InteriorPhotos, error)
	GetInteriorPhoto(ctx context.Context, UnitId int) ([]InteriorPhotos, error)
	DeleteInteriorPhoto(ctx context.Context, UnitId int, photoId uint) error

	//	for manage facility
	// AddFacilityToUnit(ctx context.Context, unitId uint, facilityId uint) ([]UnitFacilities, error)
	SearchFacility(ctx context.Context, name string) (_facility.Core, error)
	GetFacility(ctx context.Context, facilityId uint) (_facility.Core, error)
	DeleteFacility(ctx context.Context, unitId uint, facilityId uint) error
}
