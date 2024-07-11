package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()
	handler := NewRecipesHandler(NewInMemoryStorage())
	group := router.Group("/recipes")
	{
		group.GET("/", handler.ListRecipes)
		group.GET("/:id", handler.GetRecipe)
		group.POST("/", handler.CreateRecipe)
		group.DELETE("/:id", handler.DeleteRecipe)
		group.PUT("/:id", handler.UpdateRecipe)
	}

	log.Fatal("Some error when try to run server", router.Run("localhost:8080"))
}
