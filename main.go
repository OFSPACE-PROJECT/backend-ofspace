package main

import (
	"ofspace-be/config"
	"ofspace-be/migration"
	"ofspace-be/routes"
)

func main() {
	config.InitDB()
	migration.AutoMigrate()
	e := routes.New()
	err := e.Start(":8080")
	if err != nil {
		return
	}
}
