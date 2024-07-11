package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"net/http"
)

type RecipesHandler struct {
	store recipeStore
}

func NewRecipesHandler(s recipeStore) *RecipesHandler {
	return &RecipesHandler{store: s}
}

func (r RecipesHandler) CreateRecipe(c *gin.Context) {
	var recipe Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := slug.Make(recipe.Name)
	if err := r.store.Add(id, recipe); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func (r RecipesHandler) ListRecipes(c *gin.Context) {
	list, _ := r.store.List()
	c.JSON(http.StatusOK, list)
}

func (r RecipesHandler) GetRecipe(c *gin.Context) {
	recipe, _ := r.store.Get(c.Param("id"))
	c.JSON(http.StatusOK, recipe)
}

func (r RecipesHandler) UpdateRecipe(c *gin.Context) {
	id := c.Param("id")
	var recipe Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := r.store.Update(id, recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func (r RecipesHandler) DeleteRecipe(c *gin.Context) {
	if err := r.store.Remove(c.Param("id")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}
