package migration

import (
	"ofspace_be/config"
	user "ofspace_be/features/users/data"
)

func AutoMigrate() {
	db := config.DB
	if err := db.Exec("SET FOREIGN_KEY_CHECKS = 0").Error; err != nil {
		panic(err)
	}
	//	pake rds aws sebelumnya dulu aja
	if err := db.Exec("DROP TABLE IF EXISTS user_course_videos").Error; err != nil {
		panic(err)
	}
	if err := db.Exec("DROP TABLE IF EXISTS user_courses").Error; err != nil {
		panic(err)
	}
	if err := db.Exec("DROP TABLE IF EXISTS user_notes").Error; err != nil {
		panic(err)
	}
	if err := db.Exec("DROP TABLE IF EXISTS user_tasks").Error; err != nil {
		panic(err)
	}
	if err := db.Exec("DROP TABLE IF EXISTS users").Error; err != nil {
		panic(err)
	}
	if err := db.Exec("DROP TABLE IF EXISTS tasks").Error; err != nil {
		panic(err)
	}
	if err := db.Exec("DROP TABLE IF EXISTS videos").Error; err != nil {
		panic(err)
	}
	if err := db.Exec("DROP TABLE IF EXISTS courses").Error; err != nil {
		panic(err)
	}

	err := db.AutoMigrate(
		&user.User{},
	)
	if err != nil {
		panic(err)
	}
}
