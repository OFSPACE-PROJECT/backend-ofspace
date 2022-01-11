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
	compData := accessibilityData.NewAccessibilityData(config.DB)
	compBusiness := accessibilityBusiness.NewAccessibilityBusiness(compData, timeoutContext)
	compPresentation := accessibilityPresentation.NewAccessibilityPresentation(compBusiness)

	return &Presenter{
		compPresentation,
	}
}
