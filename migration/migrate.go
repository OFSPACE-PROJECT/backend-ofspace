package migration

import (
	"ofspace-be/config"
	complex2 "ofspace-be/features/complex/data"
)

func AutoMigrate() {
	db := config.DB
	if err := db.Exec("SET FOREIGN_KEY_CHECKS = 0").Error; err != nil {
		panic(err)
	}

	if err := db.Exec("DROP TABLE IF EXISTS users").Error; err != nil {
		panic(err)
	}

	if err := db.Exec("DROP TABLE IF EXISTS complex").Error; err != nil {
		panic(err)
	}

	err := db.AutoMigrate(
		//&user.User{},
		&complex2.Complex{},
	)
	if err != nil {
		panic(err)
	}

}
