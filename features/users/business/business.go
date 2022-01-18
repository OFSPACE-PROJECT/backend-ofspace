package business

import (
	"ofspace-be/features/users"
	"ofspace-be/helpers/encrypt"
	"ofspace-be/middleware"
	"time"

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

	err = encrypt.CheckPassword(data.Password, userData.Password)
	if err != nil {
		return users.Core{}, err
	}
	userData.Token, err = middleware.CreateTokens(userData.ID, userData.Name, userData.Role)
	if err != nil {
		return userData, err
	}
	return userData, nil
}

func (ub *usersBusiness) RegisterUser(c context.Context, data users.Core) (users.Core, error) {
	// if data.Name == "" || data.Email == "" || data.Password == "" {
	// 	return users.Core{}, error
	// }

	ctx, error := context.WithTimeout(c, ub.contextTimeout)
	defer error()

	data.UpdatedAt = time.Now()
	data.Password, _ = encrypt.Hash(data.Password)
	user, err := ub.userData.RegisterUser(ctx, data)
	if err != nil {
		return users.Core{}, err
	}

	return user, nil

}

func (ub *usersBusiness) GetUserByID(c context.Context, id uint) (users.Core, error) {
	ctx, error := context.WithTimeout(c, ub.contextTimeout)
	defer error()

	// if id == 0 {
	// 	return users.Core{}, error
	// }

	user, err := ub.userData.GetUserByID(ctx, id)
	if err != nil {
		return users.Core{}, err
	}

	return user, nil

}

func (ub *usersBusiness) SearchUserByName(c context.Context, name string) ([]users.Core, error) {
	ctx, error1 := context.WithTimeout(c, ub.contextTimeout)
	defer error1()

	user, err := ub.userData.SearchUserByName(ctx, name)
	if err != nil {
		return []users.Core{}, err
	}

	return user, nil

}

func (ub *usersBusiness) UpdateUser(c context.Context, data users.Core) (users.Core, error) {
	// if Id == 0 {
	// 	return User{}, resp.ErrFillData
	// }

	ctx, error := context.WithTimeout(c, ub.contextTimeout)
	defer error()
	_, err := ub.userData.GetUserByID(ctx, data.ID)
	if err != nil {
		return users.Core{}, err
	}
	data.UpdatedAt = time.Now()

	up, err := ub.userData.UpdateUser(ctx, data)
	if err != nil {
		return users.Core{}, err
	}

	return up, nil

}

func (ub *usersBusiness) DeleteUser(c context.Context, id uint) (users.Core, error) {

	// if id == 0 {
	// 	return User{}, resp.ErrFillData
	// }

	ctx, error := context.WithTimeout(c, ub.contextTimeout)
	defer error()

	_, err := ub.userData.GetUserByID(ctx, id)
	if err != nil {
		return users.Core{}, err
	}

	del, err := ub.userData.DeleteUser(ctx, id)
	if err != nil {
		return users.Core{}, err
	}

	return del, nil
}

func (ub *usersBusiness) GetUserByAdminStatus(c context.Context, status string) ([]users.Core, error) {
	ctx, error1 := context.WithTimeout(c, ub.contextTimeout)
	defer error1()

	user, err := ub.userData.SearchUserByName(ctx, status)
	if err != nil {
		return []users.Core{}, err
	}

	return user, nil
}
