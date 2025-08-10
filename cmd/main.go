package main

import (
	"github.com/ardianilyas/go-auth/internal/routes"
)

func main() {
	r := routes.Setup()

	r.Run(":8000")
}