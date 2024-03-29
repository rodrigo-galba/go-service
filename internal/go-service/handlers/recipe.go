package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	models "github.com/rodrigo-galba/go-service/internal/go-service/models"
	"github.com/rs/xid"
	"log"
	"net/http"
	"strings"
	"time"
)

type RecipesHandler struct {
	ctx         context.Context
	recipes[]	models.Recipe
}

func NewRecipesHandler(ctx context.Context, recipesList []models.Recipe) *RecipesHandler {
	return &RecipesHandler{
		ctx:      ctx,
		recipes:  recipesList,
	}
}

// swagger:operation POST /recipes recipes newRecipe
// Create a new recipe
// ---
// produces:
// - application/json
// responses:
//     '200':
//         description: Successful operation
//     '400':
//         description: Invalid input
func (handler *RecipesHandler) NewRecipeHandler(ctx *gin.Context) {
	log.Println("creating a recipe")

	var recipe models.Recipe
	var recipes = handler.recipes
	if err := ctx.ShouldBindJSON(&recipe); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	recipe.ID = xid.New().String()
	recipe.PublishedAt = time.Now()
	recipes = append(recipes, recipe)
	ctx.JSON(http.StatusOK, recipe)
	log.Printf("Recipe created ID:%d ", recipe.ID)
}


// swagger:operation GET /recipes recipes listRecipes
// Returns list of recipes
// ---
// produces:
// - application/json
// responses:
//     '200':
//         description: Successful operation
func (handler *RecipesHandler) ListRecipesHandler(c *gin.Context) {
	log.Println("listing recipes")
	c.JSON(http.StatusOK, handler.recipes)
}

// swagger:operation PUT /recipes/{id} recipes updateRecipe
// Update an existing recipe
// ---
// parameters:
// - name: id
//   in: path
//   description: ID of the recipe
//   required: true
//   type: string
// produces:
// - application/json
// responses:
//     '200':
//         description: Successful operation
//     '400':
//         description: Invalid input
//     '404':
//         description: Invalid recipe ID
func (handler *RecipesHandler) UpdateRecipeHandler(c *gin.Context) {
	id := c.Param("id")
	var recipe models.Recipe
	var recipes = handler.recipes

	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
	}

	index := -1
	for i := 0; i < len(recipes); i++ {
		if recipes[i].ID == id {
			index = i
		}
	}

	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Recipe not found"})
		return
	}

	handler.recipes[index] = recipe
	c.JSON(http.StatusOK, recipe)
}

// swagger:operation DELETE /recipes/{id} recipes deleteRecipe
// Delete an existing recipe
// ---
// produces:
// - application/json
// parameters:
//   - name: id
//     in: path
//     description: ID of the recipe
//     required: true
//     type: string
// responses:
//     '200':
//         description: Successful operation
//     '404':
//         description: Invalid recipe ID
func (handler *RecipesHandler) DeleteRecipeHandler(c *gin.Context) {
	id := c.Param("id")
	log.Println("deleting recipe by id: " + id)
	defer log.Println("Recipe by id: " + id + " deleted.")
	var recipes = handler.recipes

	index := -1
	for i := 0; i < len(handler.recipes); i++ {
		if handler.recipes[i].ID == id {
			index = i
		}
	}
	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Recipe not found"})
		return
	}
	recipes = append(recipes[:index], recipes[index+1:]...)
	c.JSON(http.StatusOK, gin.H{
		"message": "Recipe has been deleted"})
}

// swagger:operation GET /recipes/search recipes findRecipe
// Search recipes based on tags
// ---
// produces:
// - application/json
// parameters:
//   - name: tag
//     in: query
//     description: recipe tag
//     required: true
//     type: string
// responses:
//     '200':
//         description: Successful operation
func (handler *RecipesHandler) SearchRecipesHandler(c *gin.Context) {
	tag := c.Query("tag")
	log.Println("Searching recipes by tag: " + tag)
	var recipes = handler.recipes

	listOfRecipes := make([]models.Recipe, 0)
	for i := 0; i < len(recipes); i++ {
		found := false
		for _, t := range recipes[i].Tags {
			found = strings.EqualFold(t, tag)
		}
		if found {
			listOfRecipes = append(listOfRecipes, recipes[i])
			log.Println("Tag found: " + tag)
		}
	}

	c.JSON(http.StatusOK, listOfRecipes)
}

// swagger:operation GET /recipes/{id} recipes getRecipe
// Get one recipe
// ---
// produces:
// - application/json
// parameters:
//   - name: id
//     in: path
//     description: ID of the recipe
//     required: true
//     type: string
// responses:
//     '200':
//         description: Successful operation
//     '404':
//         description: Invalid recipe ID
func (handler *RecipesHandler) GetRecipeHandler(c *gin.Context) {
	var recipes = handler.recipes

	id := c.Param("id")
	for i := 0; i < len(recipes); i++ {
		if recipes[i].ID == id {
			c.JSON(http.StatusOK, recipes[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Recipe not found"})
}
