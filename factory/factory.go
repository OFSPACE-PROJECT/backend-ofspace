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

	//unit
	unitBusiness "ofspace-be/features/unit/business"
	unitData "ofspace-be/features/unit/data"
	unitPresentation "ofspace-be/features/unit/presentation"

	//wishlist
	wishlistBusiness "ofspace-be/features/wishlist/business"
	wishlistData "ofspace-be/features/wishlist/data"
	wishlistPresentation "ofspace-be/features/wishlist/presentation"

	//booking
	bookingBusiness "ofspace-be/features/booking/business"
	bookingData "ofspace-be/features/booking/data"
	bookingPresentation "ofspace-be/features/booking/presentation"

	"time"
)

type Presenter struct {
	UserPresentation          *userPresentation.UsersPresentation
	ComplexPresentation       *complexPresentation.ComplexPresentation
	AccessibilityPresentation *accessibilityPresentation.AccessibilityPresentation
	FacilityPresentation      *facilityPresentation.FacilityPresentation
	BuildingPresentation      *buildingPresentation.BuildingPresentation
	UnitPresentation          *unitPresentation.UnitPresentation
	WishlistPresentation      *wishlistPresentation.WishlistPresentation
	BookingPresentation       *bookingPresentation.BookingPresentation
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
	usData := userData.NewUserData(config.DB)
	usBusiness := userBusiness.NewUserBusiness(usData, timeoutContext)
	usPresentation := userPresentation.NewUserPresentation(usBusiness)

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

	//building
	bData := buildingData.NewBuildingData(config.DB)
	bBusiness := buildingBusiness.NewBuildingBusiness(bData, timeoutContext, usBusiness, facBusiness)
	bPresentation := buildingPresentation.NewBuildingPresentation(bBusiness)

	//unit
	uData := unitData.NewUnitData(config.DB)
	uBusiness := unitBusiness.NewUnitBusiness(uData, bBusiness, usBusiness, facBusiness, timeoutContext)
	uPresentation := unitPresentation.NewUnitPresentation(uBusiness)

	//wishlist
	wData := wishlistData.NewWishlistData(config.DB)
	wBusiness := wishlistBusiness.NewWishlistBusiness(wData, timeoutContext)
	wPresentation := wishlistPresentation.NewWishlistPresentation(wBusiness)

	//wishlist
	bookData := bookingData.NewBookingData(config.DB)
	bookBusiness := bookingBusiness.NewBookingBusiness(bookData, timeoutContext)
	bookPresentation := bookingPresentation.NewBookingPresentation(bookBusiness)
	return &Presenter{
		usPresentation,
		compPresentation,
		accPresentation,
		facPresentation,
		bPresentation,
		uPresentation,
		wPresentation,
		bookPresentation,
	}
}
