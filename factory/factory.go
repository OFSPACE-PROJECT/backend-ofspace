package factory

import (
	"github.com/joho/godotenv"
	"ofspace-be/config"
	complexBusiness "ofspace-be/features/complex/business"
	complexData "ofspace-be/features/complex/data"
	complexPresentation "ofspace-be/features/complex/presentation"
	"os"
	"strconv"
	"time"
)

type Presenter struct {
	ComplexPresentation *complexPresentation.ComplexPresentation
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
	//	Complex
	compData := complexData.NewComplexData(config.DB)
	compBusiness := complexBusiness.NewComplexBusiness(compData, timeoutContext)
	compPresentation := complexPresentation.NewComplexPresentation(compBusiness)

	return &Presenter{
		compPresentation,
	}
}
