package data

import (
	"ofspace_be/features/users"
	"time"
)

type User struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Password    string
	Role        string
	Email       string
	Phone       string
	AdminStatus string
	CreatedAt   time.Time
	UpdatedAt   time.Time
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
		Password:    core.Password,
		Phone:       core.Phone,
		AdminStatus: core.AdminStatus,
		CreatedAt:   core.CreatedAt,
		UpdatedAt:   core.UpdatedAt,
	}
}
