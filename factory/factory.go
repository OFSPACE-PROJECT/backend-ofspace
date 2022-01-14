package factory

import (
	"ofspace-be/config"

	"os"
	"strconv"

	"github.com/joho/godotenv"
	//user
	userBusiness "ofspace-be/features/users/business"
	userData "ofspace-be/features/users/data"
	userPresentation "ofspace-be/features/users/presentation"

	//complex
	complexBusiness "ofspace-be/features/complex/business"
	complexData "ofspace-be/features/complex/data"
	complexPresentation "ofspace-be/features/complex/presentation"

	//accessibility
	accessibilityBusiness "ofspace-be/features/accessibility/business"
	accessibilityData "ofspace-be/features/accessibility/data"
	accessibilityPresentation "ofspace-be/features/accessibility/presentation"

	//facility
	facilityBusiness "ofspace-be/features/facility/business"
	facilityData "ofspace-be/features/facility/data"
	facilityPresentation "ofspace-be/features/facility/presentation"

	//building
	buildingBusiness "ofspace-be/features/building/business"
	buildingData "ofspace-be/features/building/data"
	buildingPresentation "ofspace-be/features/building/presentation"

	"time"
)

type Presenter struct {
	UserPresentation          *userPresentation.UsersPresentation
	ComplexPresentation       *complexPresentation.ComplexPresentation
	AccessibilityPresentation *accessibilityPresentation.AccessibilityPresentation
	FacilityPresentation      *facilityPresentation.FacilityPresentation
	BuildingPresentation      *buildingPresentation.BuildingPresentation
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

	//complex
	compData := complexData.NewComplexData(config.DB)
	compBusiness := complexBusiness.NewComplexBusiness(compData, timeoutContext)
	compPresentation := complexPresentation.NewComplexPresentation(compBusiness)

	//	Accessibility
	accData := accessibilityData.NewAccessibilityData(config.DB)
	accBusiness := accessibilityBusiness.NewAccessibilityBusiness(accData, timeoutContext)
	accPresentation := accessibilityPresentation.NewAccessibilityPresentation(accBusiness)

	//facility
	facData := facilityData.NewFacilityData(config.DB)
	facBusiness := facilityBusiness.NewFacilityBusiness(facData, timeoutContext)
	facPresentation := facilityPresentation.NewFacilityPresentation(facBusiness)

	//buildingfacilityData
	bData := buildingData.NewBuildingData(config.DB)
	bBusiness := buildingBusiness.NewBuildingBusiness(bData, timeoutContext, userBusiness, facBusiness)
	bPresentation := buildingPresentation.NewBuildingPresentation(bBusiness)
	return &Presenter{
		userPresentation,
		compPresentation,
		accPresentation,
		facPresentation,
		bPresentation,
	}
}
