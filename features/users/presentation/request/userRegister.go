package request

import (
	"ofspace-be/features/users"
)

type UserRegister struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Phone    string `json:"phone"`
}

func (user *UserRegister) ToCore() users.Core {
	return users.Core{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
		Phone:    user.Phone,
	}
}
