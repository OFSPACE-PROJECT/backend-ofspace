package response

import "ofspace_be/features/users"

type User struct {
	ID       uint   `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"Email" form:"username"`
	Password string `json:"password" form:"password"`
}

func FromUserCore(req User) users.Core {
	return users.Core{
		ID:       req.ID,
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
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
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}
