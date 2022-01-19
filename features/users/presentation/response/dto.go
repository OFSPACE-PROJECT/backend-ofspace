package response

import (
	"ofspace-be/features/users"
)

type User struct {
	ID          uint   `json:"id" form:"id"`
	Name        string `json:"name" form:"name"`
	Email       string `json:"Email" form:"username"`
	Password    string `json:"password" form:"password"`
	Phone       string `json:"phone" form:"phone"`
	AdminStatus string `json:"admin_status"`
}

func FromUserCore(req User) users.Core {
	return users.Core{
		ID:       req.ID,
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Phone:    req.Phone,
	}
}

type UserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Token string `json:"token"`
}

func ToUserLoginResponse(user users.Core) UserResponse {
	return UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Token: user.Token,
	}
}
func ToUserRegisterResponse(user users.Core) User {
	return User{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		Password:    user.Password,
		Phone:       user.Phone,
		AdminStatus: user.AdminStatus,
	}
}

type GetUserResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	AdminStatus string `json:"admin_status"`
}

func ToGetUserResponse(user users.Core) GetUserResponse {
	return GetUserResponse{
		ID:          user.ID,
		Name:        user.Name,
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
