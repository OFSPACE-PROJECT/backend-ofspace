package factory

import (
	"ofspace_be/config"
	"os"
	"strconv"

	"github.com/joho/godotenv"

	userBusiness "ofspace_be/features/users/business"
	userData "ofspace_be/features/users/data"
	userPresentation "ofspace_be/features/users/presentation"
	"time"
)

type Presenter struct {
	UserPresentation *userPresentation.UsersPresentation
}

func New() *Presenter {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err.Error())
	}
	ctx, err := strconv.Atoi(os.Getenv("CONTEXT_TIMEOUT"))
	if err != nil {
		panic(err.Error())
	}
	timeoutContext := time.Duration(ctx) * time.Second
	//users
	userData := userData.NewUserData(config.DB)
	userBusiness := userBusiness.NewUserBusiness(userData, timeoutContext)
	userPresentation := userPresentation.NewUserPresentation(userBusiness)

	return &Presenter{
		userPresentation,
	}
}
