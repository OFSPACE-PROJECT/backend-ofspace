package migration

import (
	"ofspace-be/config"
	acc "ofspace-be/features/accessibility/data"
)

func AutoMigrate() {
	db := config.DB
	if err := db.Exec("SET FOREIGN_KEY_CHECKS = 0").Error; err != nil {
		panic(err)
	}

	if err := db.Exec("DROP TABLE IF EXISTS users").Error; err != nil {
		panic(err)
	}

	err := db.AutoMigrate(
		//&user.User{},
		&acc.Accessibility{},
	)
	if err != nil {
		panic(err)
	}

}
