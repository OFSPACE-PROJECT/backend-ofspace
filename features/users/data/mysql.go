package data

import (
	"context"
	"fmt"
	"ofspace-be/features/users"

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

func (rep *UserData) SearchUserByName(ctx context.Context, name string) ([]users.Core, error) {
	var user []User
	result := rep.Connect.Where("name LIKE ?", "%"+name+"%").Find(&user)
	if result.Error != nil {
		return []users.Core{}, result.Error
	}
	return ListToCore(user), nil
}
func (rep *UserData) GetUserByAdminStatus(ctx context.Context, status string) ([]users.Core, error) {
	var user []User
	result := rep.Connect.Where("admin_status LIKE ?", "%"+status+"%").Find(&user)
	if result.Error != nil {
		return []users.Core{}, result.Error
	}
	return ListToCore(user), nil
}
func (rep *UserData) UpdateUser(ctx context.Context, domain users.Core) (users.Core, error) {
	user := FromCore(domain)
	result := rep.Connect.Where("id = ?", user.ID).Updates(&User{Name: user.Name, Email: user.Email, Password: user.Password, Role: user.Role, Phone: user.Phone, AdminStatus: user.AdminStatus})

	if result.Error != nil {
		return users.Core{}, result.Error
	}

	return toUserCore(user), nil
}

func (rep *UserData) DeleteUser(ctx context.Context, id uint) (users.Core, error) {
	var user User
	result := rep.Connect.Delete(&user, "id= ?", id)

	if result.Error != nil {
		return users.Core{}, result.Error
	}

	return toUserCore(user), nil
}
