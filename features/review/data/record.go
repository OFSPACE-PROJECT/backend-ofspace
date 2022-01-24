package data

import (
	"ofspace-be/features/review"
	"time"

	"gorm.io/gorm"
)

type Review struct {
	Id               uint `gorm:"primaryKey"`
	CustomerId       uint
	Customer         User `gorm:"foreignKey:CustomerId"`
	UnitId           uint
	Unit             Unit `gorm:"foreignKey:UnitId"`
	BookingId        uint
	Booking          Booking `gorm:"foreignKey:BookingId"`
	RatingAccess     float32
	RatingFacility   float32
	RatingManagement float32
	RatingQuality    float32
	Comment          string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}

type Booking struct {
	ID          uint `gorm:"primaryKey"`
	TotalBought uint
	Price       float32
	DealDate    time.Time
	StartDate   time.Time
	EndDate     time.Time
}

type User struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}

type Unit struct {
	Id          uint `gorm:"primaryKey"`
	BuildingId  uint `gorm:"not null"`
	Description string
	UnitType    string
}

func toUserCore(u User) review.User {
	return review.User{
		ID:   u.ID,
		Name: u.Name,
	}
}

func toUnitCore(b Unit) review.Unit {
	return review.Unit{
		Id:          b.Id,
		BuildingId:  b.BuildingId,
		Description: b.Description,
		UnitType:    b.UnitType,
	}
}

func toReviewCore(u Review) review.Core {
	return review.Core{
		Id:               u.Id,
		CustomerId:       u.CustomerId,
		Customer:         toUserCore(u.Customer),
		UnitId:           u.UnitId,
		Unit:             toUnitCore(u.Unit),
		BookingId:        u.BookingId,
		RatingAccess:     u.RatingAccess,
		RatingFacility:   u.RatingFacility,
		RatingManagement: u.RatingManagement,
		RatingQuality:    u.RatingQuality,
		Comment:          u.Comment,
		CreatedAt:        u.CreatedAt,
		UpdatedAt:        u.UpdatedAt,
	}
}

func ToListCore(data []Review) (result []review.Core) {
	result = []review.Core{}
	for _, res := range data {
		result = append(result, toReviewCore(res))
	}
	return
}
func FromListCore(data []review.Core) (result []Review) {
	result = []Review{}
	for _, res := range data {
		result = append(result, FromCore(res))
	}
	return
}
func FromCore(core review.Core) Review {
	return Review{
		Id:               core.Id,
		CustomerId:       core.CustomerId,
		UnitId:           core.UnitId,
		BookingId:        core.BookingId,
		RatingAccess:     core.RatingAccess,
		RatingFacility:   core.RatingFacility,
		RatingManagement: core.RatingManagement,
		RatingQuality:    core.RatingQuality,
		Comment:          core.Comment,
		CreatedAt:        core.CreatedAt,
		UpdatedAt:        core.UpdatedAt,
	}
}
