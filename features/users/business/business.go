package business

import (
	"ofspace_be/features/users"
	"time"

	// "ofspace_be/middleware"
	"context"
)

type usersBusiness struct {
	userData       users.Data
	contextTimeout time.Duration
}

func NewUserBusiness(userData users.Data, timeout time.Duration) users.Business {
	return &usersBusiness{userData: userData, contextTimeout: timeout}
}

func (ub *usersBusiness) LoginUser(c context.Context, data users.Core) (users.Core, error) {
	userData, err := ub.userData.LoginUser(c, data)
	if err != nil {
		return users.Core{}, err
	}

	// userData.Token, err = middleware.CreateTokens(userData.ID, userData.Name)
	// if err != nil {
	// 	return userData, err
	// }
	return userData, nil
}

func (ub *usersBusiness) RegisterUser(c context.Context, data users.Core) (users.Core, error) {

	// if data.Name == "" || data.Email == "" || data.Password == "" {
	// 	return users.Core{}, error
	// }

	ctx, error := context.WithTimeout(c, ub.contextTimeout)
	defer error()

	data.UpdatedAt = time.Now()
	// data.Password, _ = encrypt.Hash(data.Password)
	user, err := ub.userData.RegisterUser(ctx, data)
	if err != nil {
		return users.Core{}, err
	}

	return user, nil

}
