package main

import (
	"ofspace_be/config"
	"ofspace_be/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	e.Start(":8080")
}
