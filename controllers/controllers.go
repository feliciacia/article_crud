package controllers

import (
	"article_crud/database"
	"article_crud/models"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var artcollection *mongo.Collection = database.GetCollection(database.DB, "article")

func CreateArticle() gin.HandlerFunc {
	return func(c *gin.Context) {
		var art models.Article
		if err := c.BindJSON(&art); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newArticle := models.Article{
			ID:      xid.New().String(), //xid for generate unique id
			Title:   art.Title,
			Content: art.Content,
		}

		result, err := artcollection.InsertOne(context.Background(), newArticle)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, result)
	}
}

func GetArticle() gin.HandlerFunc {
	return func(c *gin.Context) {
		var art models.Article
		artid := c.Param("id")
		err := artcollection.FindOne(context.Background(), bson.M{"_id": artid}).Decode(&art)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, art)
	}
}

func UpdateArticle() gin.HandlerFunc {
	return func(c *gin.Context) {
		artid := c.Param("id")
		var input models.Article
		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var art models.Article
		err := artcollection.FindOne(context.Background(), bson.M{"_id": artid}).Decode(&art)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if input.Title != "" {
			art.Title = input.Title
		}
		if input.Content != "" {
			art.Content = input.Content
		}
		_, err = artcollection.UpdateOne(context.Background(), bson.M{"_id": artid}, bson.M{"$set": art})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, art)
	}
}
