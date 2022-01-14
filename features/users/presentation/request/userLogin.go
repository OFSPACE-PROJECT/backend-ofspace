package request

import (
	"ofspace-be/features/users"
)

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (user *UserLogin) ToCore() users.Core {
	return users.Core{
		Email:    user.Email,
		Password: user.Password,
	}
}
