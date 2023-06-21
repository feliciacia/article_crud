package main

import (
	"article_crud/database"
	"article_crud/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.ArtRoute(router)
	database.ConnectDB()
	router.Run(":6000")
}
