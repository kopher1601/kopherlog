package main

import (
	"kopherlog/router"
)

func main() {
	router := router.Setup()

	router.Run(":8080")
}
