package main

import (
	"go-docker/internal/db"
	"go-docker/pkg/router"
)

func main() {
	db.InitDB()
	r := router.SetupRouter()
	r.Run()
}
