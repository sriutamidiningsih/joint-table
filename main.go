package main

import (
	"dummy/config"
	"dummy/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	db := config.Database()
	router := gin.Default()
	routes.User(db, router)
	router.Run()
}