package main

import (
	_ "github.com/go-sql-driver/mysql"
	"kopherlog/config"
	"kopherlog/router"
)

func main() {
	config.Initialize()
	client := config.SetupDB()
	r := router.Setup(client)

	r.Run(":8080")
}
