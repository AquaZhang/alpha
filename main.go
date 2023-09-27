package main

import (
	"main/handlers"
	"main/routes"
)

func main() {
	handlers.Setup()
	r := routes.Setup()
	r.Run(":8000")
}
