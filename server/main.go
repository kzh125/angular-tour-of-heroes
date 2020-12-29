package main

import (
	"heroes/server/router"
)

func main() {
	r := router.SetupRouter()
	r.Run(":4201")
}
