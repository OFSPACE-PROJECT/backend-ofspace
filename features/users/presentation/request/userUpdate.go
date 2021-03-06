package request

import "ofspace-be/features/users"

type UserUpdate struct {
	Id       uint   `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Phone    string `json:"phone"`
}

func (user *UserUpdate) ToCoreUpdate() users.Core {
	return users.Core{
		ID:       user.Id,
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
		Phone:    user.Phone,
	}
}
