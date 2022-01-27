package business

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	complex2 "ofspace-be/features/complex"
	"ofspace-be/features/users"
	mocks "ofspace-be/features/users/mock"
	"os"
	"testing"
	"time"
)

var (
	compData     mocks.Data
	compBusiness users.Business
	compValue    users.Core
	compList     []users.Core
)

func TestMain(m *testing.M) {
	compBusiness = NewUserBusiness(&compData, time.Hour*1)
	compValue = users.Core{
		ID:          1,
		Name:        "test",
		Password:    "",
		Role:        "",
		Email:       "",
		Phone:       "",
		Token:       "",
		AdminStatus: "",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}

	compList = []users.Core{
		{
			ID:          1,
			Name:        "test",
			Password:    "",
			Role:        "",
			Email:       "",
			Phone:       "",
			Token:       "",
			AdminStatus: "",
			CreatedAt:   time.Time{},
			UpdatedAt:   time.Time{},
		},
	}
	os.Exit(m.Run())
}

func TestRegisterUser(t *testing.T) {
	t.Run("RegisterUser sukses", func(t *testing.T) {
		compData.On("RegisterUser", mock.Anything, mock.Anything).Return(compValue, nil).Once()
		data, err := compBusiness.RegisterUser(context.Background(), compValue)
		assert.NotEqual(t, data, err)
	})
	t.Run("RegisterUser gagal", func(t *testing.T) {
		compData.On("RegisterUser", mock.Anything, mock.Anything).Return(compValue, errors.New("err")).Once()
		data, err := compBusiness.RegisterUser(context.Background(), compValue)
		assert.NotEqual(t, data, err)
		assert.Error(t, err)
	})
}

func TestLoginUser(t *testing.T) {
	t.Run("LoginUser sukses", func(t *testing.T) {
		compData.On("LoginUser", mock.Anything, mock.Anything).Return(compValue, nil).Once()
		data, err := compBusiness.LoginUser(context.Background(), compValue)
		assert.NotEqual(t, data, err)
	})
	t.Run("LoginUser gagal", func(t *testing.T) {
		compData.On("LoginUser", mock.Anything, mock.Anything).Return(compValue, errors.New("err")).Once()
		data, err := compBusiness.LoginUser(context.Background(), compValue)
		assert.NotEqual(t, data, err)
		assert.Error(t, err)
	})
}

//
func TestGetUserByID(t *testing.T) {
	t.Run("GetUserByID sukses", func(t *testing.T) {
		compData.On("GetUserByID", mock.Anything, mock.AnythingOfType("uint")).Return(compValue, nil).Once()
		data, err := compBusiness.GetUserByID(context.Background(), 1)
		assert.Equal(t, data, compValue)
		assert.Nil(t, err)
	})
	t.Run("GetUserByID gagal", func(t *testing.T) {
		compData.On("GetUserByID", mock.Anything, mock.AnythingOfType("uint")).Return(compValue, errors.New("err")).Once()
		data, err := compBusiness.GetUserByID(context.Background(), 2)
		assert.NotEqual(t, data, complex2.Core{})
		assert.Error(t, err)
	})
}

//func TestGetAllComplex(t *testing.T) {
//	t.Run("get all complex sukses", func(t *testing.T) {
//		compData.On("GetAllComplex", mock.Anything).Return(compList, nil).Once()
//		data, err := compBusiness.GetAllComplex(context.Background())
//		assert.Equal(t, data, compList)
//		assert.Nil(t, err)
//	})
//	t.Run("get all complex gagal", func(t *testing.T) {
//		compData.On("GetAllComplex", mock.Anything).Return(compList, errors.New("err")).Once()
//		data, err := compBusiness.GetAllComplex(context.Background())
//		assert.Equal(t, data, []complex2.Core{})
//		assert.Error(t, err)
//	})
//}
//
func TestSearchUserByName(t *testing.T) {
	t.Run("SearchUserByName sukses", func(t *testing.T) {
		compData.On("SearchUserByName", mock.Anything, mock.AnythingOfType("string")).Return(compList, nil).Once()
		data, err := compBusiness.SearchUserByName(context.Background(), compList[0].Name)
		assert.Equal(t, data, compList)
		assert.Nil(t, err)
	})

	t.Run("SearchUserByName gagal", func(t *testing.T) {
		compData.On("SearchUserByName", mock.Anything, mock.AnythingOfType("string")).Return(compList, errors.New("err")).Once()
		data, err := compBusiness.SearchUserByName(context.Background(), "test")
		assert.NotEqual(t, data, users.Core{})
		assert.Error(t, err)
	})
}

//func TestSearchComplex(t *testing.T) {
//	t.Run("search complex sukses", func(t *testing.T) {
//		compData.On("SearchComplex", mock.Anything, mock.AnythingOfType("string")).Return(compList, nil).Once()
//		data, err := compBusiness.SearchComplex(context.Background(), "test")
//		assert.Equal(t, data, compList)
//		assert.Nil(t, err)
//	})
//	t.Run("search complex gagal", func(t *testing.T) {
//		compData.On("SearchComplex", mock.Anything, mock.AnythingOfType("string")).Return(compList, errors.New("err")).Once()
//		data, err := compBusiness.SearchComplex(context.Background(), "gagal")
//		assert.Equal(t, data, []complex2.Core{})
//		assert.Error(t, err)
//	})
//}
//
func TestUpdateUser(t *testing.T) {
	t.Run("UpdateUser sukses", func(t *testing.T) {
		compData.On("GetUserByID", mock.Anything, mock.AnythingOfType("uint")).Return(compValue, nil).Once()
		compData.On("UpdateUser", mock.Anything, mock.Anything).Return(compValue, nil).Once()
		data, err := compBusiness.UpdateUser(context.Background(), compValue)
		assert.Equal(t, data, compValue)
		assert.Nil(t, err)
	})
	t.Run("UpdateUser gagal", func(t *testing.T) {
		compData.On("GetUserByID", mock.Anything, mock.AnythingOfType("uint")).Return(users.Core{}, errors.New("err")).Once()
		compData.On("UpdateUser", mock.Anything, mock.Anything).Return(compValue, errors.New("err")).Once()
		data, err := compBusiness.UpdateUser(context.Background(), users.Core{})
		assert.Equal(t, data, users.Core{})
		assert.Error(t, err)
	})
}

func TestDeleteUser(t *testing.T) {
	t.Run("DeleteUser sukses", func(t *testing.T) {
		compData.On("GetUserByID", mock.Anything, mock.AnythingOfType("uint")).Return(compValue, nil).Once()
		compData.On("DeleteUser", mock.Anything, mock.Anything).Return(compValue, nil).Once()
		data, err := compBusiness.DeleteUser(context.Background(), 1)
		assert.Equal(t, data, compValue)
		assert.Nil(t, err)
	})
	t.Run("DeleteUser gagal", func(t *testing.T) {
		compData.On("GetUserByID", mock.Anything, mock.AnythingOfType("uint")).Return(users.Core{}, errors.New("err")).Once()
		compData.On("UpdateUser", mock.Anything, mock.Anything).Return(compValue, errors.New("err")).Once()
		data, err := compBusiness.DeleteUser(context.Background(), 1)
		assert.Equal(t, data, users.Core{})
		assert.Error(t, err)
	})
}

func TestGetUserByAdminStatus(t *testing.T) {
	t.Run("GetUserByAdminStatus sukses", func(t *testing.T) {
		compData.On("GetUserByAdminStatus", mock.Anything, mock.AnythingOfType("string")).Return(compList, nil).Once()
		data, err := compBusiness.GetUserByAdminStatus(context.Background(), "verified")
		assert.Equal(t, data, compList)
		assert.Nil(t, err)
	})
	t.Run("GetUserByAdminStatus gagal", func(t *testing.T) {
		compData.On("GetUserByAdminStatus", mock.Anything, mock.AnythingOfType("string")).Return([]users.Core{}, errors.New("err")).Once()
		data, err := compBusiness.GetUserByAdminStatus(context.Background(), "haha")
		assert.Equal(t, data, []users.Core{})
		assert.Error(t, err)
	})
}
