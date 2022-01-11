package migration

import (
	"ofspace-be/config"
	fac "ofspace-be/features/facility/data"
)

func AutoMigrate() {
	db := config.DB
	if err := db.Exec("SET FOREIGN_KEY_CHECKS = 0").Error; err != nil {
		panic(err)
	}
	//	pake rds aws sebelumnya dulu aja

	if err := db.Exec("DROP TABLE IF EXISTS users").Error; err != nil {
		panic(err)
	}

	err := db.AutoMigrate(
		//&user.User{},
		&fac.Facility{},
	)
	if err != nil {
		panic(err)
	}

}
