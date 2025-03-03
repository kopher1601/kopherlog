package main

import (
	"kopherlog/config"
	"kopherlog/router"
)

func main() {
	config.Initialize()
	client := config.SetupDB(nil)
	r := router.Setup(client)

	r.Run(":8080")
}
