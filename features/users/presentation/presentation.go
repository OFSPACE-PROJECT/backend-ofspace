package presentation

import (
	"net/http"
	"ofspace-be/features/users"
	"ofspace-be/features/users/presentation/request"
	"ofspace-be/features/users/presentation/response"

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
	err := c.Bind(&user)
	if err != nil {
		return err
	}
	ctx := c.Request().Context()
	data, err2 := up.usersBusiness.LoginUser(ctx, user.ToCore())
	if err2 != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err2.Error(),
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

func (up *UsersPresentation) DeleteUser(c echo.Context) error {

	Id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err.Error(),
		})
	}
	ctx := c.Request().Context()
	_, err = up.usersBusiness.DeleteUser(ctx, uint(Id))
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
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

func (up *UsersPresentation) SearchUserByName(c echo.Context) error {
	name := c.QueryParam("name")
	ctx := c.Request().Context()
	comp, err := up.usersBusiness.SearchUserByName(ctx, name)
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.ToListUserCore(comp),
	})
}

func (up *UsersPresentation) SearchUserByAdminStatus(c echo.Context) error {
	status := c.QueryParam("status")
	ctx := c.Request().Context()
	comp, err := up.usersBusiness.GetUserByAdminStatus(ctx, status)
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.ToListUserCore(comp),
	})
}

func (up *UsersPresentation) GetAllUser(c echo.Context) error {
	ctx := c.Request().Context()
	comp, err := up.usersBusiness.GetAllUser(ctx)
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.ToListUserCore(comp),
	})
}
