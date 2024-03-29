// Recipes API
//
// This is a sample recipes API. You can find out more about the API at  https://github.com/rodrigo-galba/go-service
//
// Schemes: http
// Host: localhost:8080
// BasePath: /
// Version: 1.0.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
// swagger:meta
package main

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	handlers "github.com/rodrigo-galba/go-service/internal/go-service/handlers"
	"github.com/rodrigo-galba/go-service/internal/go-service/models"
	"io/ioutil"
	"log"
	"os"
)

var recipesHandler *handlers.RecipesHandler

func init() {
	log.Println("Initializing service")
	godotenv.Load()
	ctx := context.Background()
	file, _ := ioutil.ReadFile(os.Getenv("RECIPES_DATASET"))
	recipesList := make([]models.Recipe, 0)
	_ = json.Unmarshal(file, &recipesList)

	recipesHandler = handlers.NewRecipesHandler(ctx, recipesList)
}

func setupRouter() *gin.Engine {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	// Health check test
	router.GET("/", handlers.HealthcheckHandler)
	router.GET("/health", handlers.HealthcheckHandler)
	router.POST("/recipes", recipesHandler.NewRecipeHandler)
	router.GET("/recipes", recipesHandler.ListRecipesHandler)
	router.PUT("/recipes/:id", recipesHandler.UpdateRecipeHandler)
	router.DELETE("/recipes/:id", recipesHandler.DeleteRecipeHandler)
	router.GET("/recipes/search", recipesHandler.SearchRecipesHandler)
	router.GET("/recipes/:id", recipesHandler.GetRecipeHandler)

	return router
}

func main() {
	r := setupRouter()
	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	r.Run(":5000")
}
