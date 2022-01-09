package data

import (
	"context"
	"fmt"
	"ofspace_be/features/users"

	"gorm.io/gorm"
)

type UserData struct {
	Connect *gorm.DB
}

func NewUserData(connect *gorm.DB) users.Data {
	return &UserData{
		Connect: connect,
	}
}

func (rep *UserData) LoginUser(ctx context.Context, domain users.Core) (users.Core, error) {
	var user User
	result := rep.Connect.First(&user, "email = ? ", domain.Email)
	if result.Error != nil {
		fmt.Println(result.Error)
		return users.Core{}, result.Error
	}
	return toUserCore(user), nil
}

func (rep *UserData) RegisterUser(ctx context.Context, domain users.Core) (users.Core, error) {
	user := FromCore(domain)

	// hashedPassword, err := encrypt.Hash(domain.Password)
	// if err != nil {
	// 	return users.Core{}, err
	// }

	// user.Password = hashedPassword

	result := rep.Connect.Create(&user)

	if result.Error != nil {
		return users.Core{}, result.Error
	}

	return toUserCore(user), nil
}

func (rep *UserData) GetUserByID(ctx context.Context, id uint) (users.Core, error) {
	var user User
	result := rep.Connect.First(&user, "id= ?", id)
	if result.Error != nil {
		return users.Core{}, result.Error
	}
	return toUserCore(user), nil
}
