package server

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

/*
	type Recipe struct {
		ID          int      `json:"id"`
		Title       string   `json:"title"`
		Ingredients []string `json:"ingredients"`
		Steps       []string `json:"steps"`
		Baketime    int      `json:"baketime"`
		Vegan       bool     `json:"vegan"`
		Author      string   `json:"author"`
		Rating      int      `json:"rating"`
	}
*/
type Recipe struct {
	ID          uint32   `json:"id"`
	Title       string   `json:"title"`
	Ingredients []string `json:"ingredients"`
	Steps       []string `json:"steps"`
	Baketime    uint8    `json:"baketime"`
	Vegan       bool     `json:"vegan"`
	Author      string   `json:"author"`
	Rating      uint8    `json:"rating"`
}

var recipeBook = []Recipe{
	{ID: 1, Title: "meat loaf", Ingredients: []string{"beef", "breadcrumbs", "spices"}, Baketime: 120, Vegan: false, Author: "elle", Rating: 5},
}

var server_running bool = false

// Function to return all recipes
func getRecipes(c *gin.Context) {
	c.JSON(http.StatusOK, dbRecipes._Recipes)
}

// Function to delete all recipes
func deleteAll(c *gin.Context) {
	recipeBook = []Recipe{}
	c.JSON(http.StatusOK, gin.H{
		"message": "All recipes deleted.",
	})
}

// Function to create a new recipe
func createRecipe(c *gin.Context) {
	var newRecipe Recipe

	// Attempt to bind newRecipe to JSON information in the Context
	if err := c.BindJSON(&newRecipe); err != nil {
		return
	}

	// Add recipe to recipe slice
	recipeBook = append(recipeBook, newRecipe) // old backend
	dbRecipes.insert(&newRecipe)               // updated backend

	// Returns status to client and the recipe just created
	c.IndentedJSON(http.StatusCreated, newRecipe)
}

// Helper to get recipe by name
func getRecipe(title string) (*Recipe, error) {
	title = strings.ToLower(title)
	for index, recipe := range recipeBook {
		if strings.ToLower(recipe.Title) == title {
			return &recipeBook[index], nil
		}
	}
	return nil, errors.New("recipe not found")
}

// Get a recipe by name
func getRecipeQuery(c *gin.Context) {
	title, title_exists := c.GetQuery("title")
	id, id_exists := c.GetQuery("id")
	if title_exists {
		recipe, err := dbRecipes.query_title(title)
		if err == nil {
			c.JSON(http.StatusOK, recipe)
			return
		}
	} else if id_exists {
		var id_uint uint64
		id_uint, _ = strconv.ParseUint(id, 10, 64)
		recipe, err := dbRecipes.query_id(id_uint)
		if err == nil {
			c.JSON(http.StatusOK, recipe)
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "missing title query"})
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "recipe not found"})
	return
}

// Get a recipe by name
func getRecipeByName(c *gin.Context) {
	title, ok := c.GetQuery("title")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "missing title query"})
		return
	}

	//recipe, err := getRecipe(title)
	recipe, err := dbRecipes.query_title(title)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Recipe not found"})
		return
	}

	c.JSON(http.StatusOK, recipe)
}

// Update recipe
func updateRecipe(c *gin.Context) {
	title, ok := c.GetQuery("title")

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "URL missing query"})
		return
	}

	var updated Recipe

	if err := c.BindJSON(&updated); err != nil {
		return
	}

	oldPtr, err := getRecipe(title)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Recipe not found"})
		return
	}

	*oldPtr = updated
	c.JSON(http.StatusOK, gin.H{
		"message": "Recipe Updated.",
	})
}

func StartServer(debug_mode bool) {
	if !debug_mode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	r.GET("/", getRecipes)
	r.DELETE("/", deleteAll)
	r.POST("/recipe", createRecipe)
	r.GET("/recipe", getRecipeQuery)
	r.PUT("/recipe", updateRecipe)

	server_running = true
	r.Run("localhost:3000") // listen and serve on port 3000
}

func ServerReady() bool {
	return server_running
}
