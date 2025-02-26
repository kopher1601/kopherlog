package main

import (
	"kopherlog/router"
)

func main() {
	r := router.Setup()

	r.Run(":8080")
}
