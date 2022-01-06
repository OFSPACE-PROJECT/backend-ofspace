package main

import (
	"ofspace_be/config"
	"ofspace_be/migration"
	"ofspace_be/routes"
)

func main() {
	config.InitDB()
	migration.AutoMigrate()
	e := routes.New()
	e.Start(":8080")
}
