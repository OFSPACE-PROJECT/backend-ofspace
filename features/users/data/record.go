package data

import (
	"ofspace-be/features/users"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"not null;type:varchar(100);"`
	Password    string `gorm:"not null;type:varchar(100);"`
	Role        string `gorm:"default:costumer"`
	Email       string `gorm:"unique;not null;type:varchar(100);"`
	Phone       string `gorm:"unique;not null;type:varchar(100);"`
	AdminStatus string `gorm:"default:unverified"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func toUserCore(u User) users.Core {
	return users.Core{
		ID:          u.ID,
		Name:        u.Name,
		Role:        u.Role,
		Email:       u.Email,
		Password:    u.Password,
		Phone:       u.Phone,
		AdminStatus: u.AdminStatus,
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
	}
}

func FromCore(core users.Core) User {
	return User{
		ID:          core.ID,
		Name:        core.Name,
		Email:       core.Email,
		Role:        core.Role,
		Password:    core.Password,
		Phone:       core.Phone,
		AdminStatus: core.AdminStatus,
		CreatedAt:   core.CreatedAt,
		UpdatedAt:   core.UpdatedAt,
	}
}

func ListToCore(core []User) (result []users.Core) {
	for _, user := range core {
		result = append(result, toUserCore(user))
	}
	return
}
