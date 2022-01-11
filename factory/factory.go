package factory

import (
	"github.com/joho/godotenv"
	"ofspace-be/config"
	accessibilityBusiness "ofspace-be/features/accessibility/business"
	accessibilityData "ofspace-be/features/accessibility/data"
	accessibilityPresentation "ofspace-be/features/accessibility/presentation"
	"os"
	"strconv"
	"time"
)

type Presenter struct {
	AccessibilityPresentation *accessibilityPresentation.AccessibilityPresentation
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
	//	Accessibility
	accData := accessibilityData.NewAccessibilityData(config.DB)
	accBusiness := accessibilityBusiness.NewAccessibilityBusiness(accData, timeoutContext)
	accPresentation := accessibilityPresentation.NewAccessibilityPresentation(accBusiness)

	return &Presenter{
		accPresentation,
	}
}
