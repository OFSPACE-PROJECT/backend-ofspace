package presentation

import (
	"net/http"
	"ofspace_be/features/users"
	"ofspace_be/features/users/presentation/request"
	"ofspace_be/features/users/presentation/response"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UsersPresentation struct {
	usersBusiness users.Business
}

// type json map[string]interface{}

func NewUserPresentation(up users.Business) *UsersPresentation {
	return &UsersPresentation{usersBusiness: up}
}

func (up *UsersPresentation) LoginUser(c echo.Context) error {
	user := request.UserLogin{}
	c.Bind(&user)
	ctx := c.Request().Context()
	data, err := up.usersBusiness.LoginUser(ctx, user.ToCore())
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.ToUserLoginResponse(data),
	})
}

func (up *UsersPresentation) RegisterUser(c echo.Context) error {

	UserRegister := request.UserRegister{}
	c.Bind(&UserRegister)

	ctx := c.Request().Context()
	data, err := up.usersBusiness.RegisterUser(ctx, UserRegister.ToCore())
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.ToUserRegisterResponse(data),
	})
}

func (up *UsersPresentation) UpdateUser(c echo.Context) error {

	userUpdate := request.UserUpdate{}
	c.Bind(&userUpdate)

	ctx := c.Request().Context()
	_, err := up.usersBusiness.UpdateUser(ctx, userUpdate.ToCoreUpdate())
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		// "data":    response.ToUserRegisterResponse(data),
	})
}

func (up *UsersPresentation) GetUserByID(c echo.Context) error {
	// fmt.Println("UserDetail")

	Id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err.Error(),
		})
	}

	ctx := c.Request().Context()
	user, err := up.usersBusiness.GetUserByID(ctx, uint(Id))
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.ToGetUserResponse(user),
	})
}
