package main

import (
	"ofspace-be/config"
	"ofspace-be/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	err := e.Start(":8080")
	if err != nil {
		return
	}
}
