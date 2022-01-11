package factory

import (
	config2 "ofspace-be/config"
	facilityBusiness "ofspace-be/features/facility/business"
	facilityData "ofspace-be/features/facility/data"
	facilityPresentation "ofspace-be/features/facility/presentation"

	"os"
	"strconv"

	"github.com/joho/godotenv"

	"time"
)

type Presenter struct {
	FacilityPresentation *facilityPresentation.FacilityPresentation
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
	//facility

	facData := facilityData.NewFacilityData(config2.DB)
	facBusiness := facilityBusiness.NewFacilityBusiness(facData, timeoutContext)
	facPresentation := facilityPresentation.NewFacilityPresentation(facBusiness)

	return &Presenter{
		facPresentation,
	}
}
