package migration

import (
	"ofspace-be/config"
	accessibility "ofspace-be/features/accessibility/data"
	building "ofspace-be/features/building/data"
	complex "ofspace-be/features/complex/data"
	facility "ofspace-be/features/facility/data"
	user "ofspace-be/features/users/data"
)

func AutoMigrate() {
	db := config.DB
	//if err := db.Exec("SET FOREIGN_KEY_CHECKS = 0").Error; err != nil {
	//	panic(err)
	//}
	//if err := db.Exec("DROP TABLE IF EXISTS users").Error; err != nil {
	//	panic(err)
	//}
	//if err := db.Exec("DROP TABLE IF EXISTS complexs").Error; err != nil {
	//	panic(err)
	//}
	//if err := db.Exec("DROP TABLE IF EXISTS complex").Error; err != nil {
	//	panic(err)
	//}
	//if err := db.Exec("DROP TABLE IF EXISTS accessibility").Error; err != nil {
	//	panic(err)
	//}
	//if err := db.Exec("DROP TABLE IF EXISTS accessibilitys").Error; err != nil {
	//	panic(err)
	//}
	//if err := db.Exec("DROP TABLE IF EXISTS accessibilities").Error; err != nil {
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
		&building.ExteriorPhoto{},
	)
	if err != nil {
		panic(err)
	}
}
