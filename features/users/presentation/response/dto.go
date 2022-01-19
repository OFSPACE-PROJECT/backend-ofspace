package response

import (
	"ofspace-be/features/users"
	"time"
)

type User struct {
	ID          uint   `json:"id" form:"id"`
	Name        string `json:"name" form:"name"`
	Email       string `json:"Email" form:"username"`
	Role        string `json:"role" form:"username"`
	Password    string `json:"password" form:"password"`
	Phone       string `json:"phone" form:"phone"`
	AdminStatus string `json:"admin_status"`
}

type UserResponse struct {
	ID          uint   `json:"id" form:"id"`
	Name        string `json:"name" form:"name"`
	Role        string `json:"role" form:"username"`
	Phone       string `json:"phone" form:"phone"`
	AdminStatus string `json:"admin_status"`
	Token       string `json:"token"`
}

func ToUserLoginResponse(user users.Core) UserResponse {
	return UserResponse{
		ID:          user.ID,
		Name:        user.Name,
		Role:        user.Role,
		Phone:       user.Phone,
		AdminStatus: user.AdminStatus,
		Token:       user.Token,
	}
}
func ToUserRegisterResponse(user users.Core) User {
	return User{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		Role:        user.Role,
		Password:    user.Password,
		Phone:       user.Phone,
		AdminStatus: user.AdminStatus,
	}
}

type GetUserResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Role        string    `json:"role"`
	Phone       string    `json:"phone"`
	AdminStatus string    `json:"admin_status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func ToGetUserResponse(user users.Core) GetUserResponse {
	return GetUserResponse{
		ID:          user.ID,
		Name:        user.Name,
		Role:        user.Role,
		Phone:       user.Phone,
		AdminStatus: user.AdminStatus,
	}
}

func ToListUserCore(core []users.Core) (response []GetUserResponse) {
	for _, comp := range core {
		response = append(response, ToGetUserResponse(comp))
	}
	return
}
