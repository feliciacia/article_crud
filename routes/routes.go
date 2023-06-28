package routes

import (
	"article_crud/controllers"
	"article_crud/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ArtRoute(router *gin.Engine) {
	router.POST("/art", controllers.CreateArticle())
	router.GET("/art/:id", controllers.GetArticle())
	router.GET("/art", controllers.GetArticles())
	router.PUT("/art/:id", controllers.UpdateArticle())
	router.DELETE("/art/:id", controllers.DeleteArticle())
	router.GET("art/api", middleware.Authenticate(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "access granted"})
	})
}
