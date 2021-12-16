package main

import (
	"tryout_echo/routes"
)

func main() {
	e := routes.Api()
	e.Start(":8181")
}