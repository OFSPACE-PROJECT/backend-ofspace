package migration

import (
	"ofspace-be/config"
	accessibility "ofspace-be/features/accessibility/data"
	booking "ofspace-be/features/booking/data"
	building "ofspace-be/features/building/data"
	complex "ofspace-be/features/complex/data"
	facility "ofspace-be/features/facility/data"
	review "ofspace-be/features/review/data"
	unit "ofspace-be/features/unit/data"
	user "ofspace-be/features/users/data"
	wishlist "ofspace-be/features/wishlist/data"
)

func AutoMigrate() {
	db := config.DB
	//if err := db.Exec("SET FOREIGN_KEY_CHECKS = 0").Error; err != nil {
	//	panic(err)
	//}
	//if err := db.Exec("DROP TABLE IF EXISTS users").Error; err != nil {
	//	panic(err)
	//}
	//if err := db.Exec("DROP TABLE IF EXISTS complexes").Error; err != nil {
	//	panic(err)
	//}
	//if err := db.Exec("DROP TABLE IF EXISTS accessibility").Error; err != nil {
	//	panic(err)
	//}

	//if err := db.Exec("DROP TABLE IF EXISTS facility").Error; err != nil {
	//	panic(err)
	//}
	//if err := db.Exec("DROP TABLE IF EXISTS facilities").Error; err != nil {
	//	panic(err)
	//}
	//if err := db.Exec("DROP TABLE IF EXISTS buildings").Error; err != nil {
	//	panic(err)
	//}

	err := db.AutoMigrate(
		&user.User{},
		&complex.Complex{},
		&accessibility.Accessibility{},
		&facility.Facility{},
		&building.Building{},
		&unit.Unit{},
		&building.ExteriorPhoto{},
		&building.FloorPhoto{},
		&unit.InteriorPhoto{},
		&wishlist.Wishlist{},
		&review.Review{},
		&booking.Booking{},
	)
	if err != nil {
		panic(err)
	}
}
