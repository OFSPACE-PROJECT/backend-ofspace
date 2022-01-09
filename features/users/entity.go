package users

import (
	"context"
	"time"
)

type Core struct {
	ID          uint
	Name        string
	Password    string
	Role        string
	Email       string
	Phone       string
	Token       string
	AdminStatus string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Business interface {
	LoginUser(ctx context.Context, data Core) (Core, error)
	RegisterUser(ctx context.Context, data Core) (Core, error)
	// GetAllUsers(ctx context.Context) ([]Core, error)
	UpdateUser(ctx context.Context, data Core) (Core, error)
	GetUserByID(ctx context.Context, id uint) (Core, error)
	// GetUserByName(ctx context.Context, name string) (Core, error)
	// GetUserByPhone(ctx context.Context, id uint, phone string) (Core, error)
	// GetUserByAdminStatus(ctx context.Context, status string) (Core, error)
}

type Data interface {
	LoginUser(ctx context.Context, data Core) (Core, error)
	RegisterUser(ctx context.Context, data Core) (Core, error)
	// GetAllUsers(ctx context.Context) ([]Core, error)
	UpdateUser(ctx context.Context, data Core) (Core, error)
	GetUserByID(ctx context.Context, id uint) (Core, error)
	// GetUserByName(ctx context.Context, name string) (Core, error)
	// GetUserByPhone(ctx context.Context, id uint, phone string) (Core, error)
	// GetUserByAdminStatus(ctx context.Context, status string) ([]Core, error)
}
