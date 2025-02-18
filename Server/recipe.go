package server

// recipe.go
// Model/Backend Controller for database.go

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// recipe structure implementation
type Recipe struct {
	ID          uint32   `json:"id"`
	Title       string   `json:"title"`
	Ingredients []string `json:"ingredients"`
	Steps       []string `json:"steps"`
	Baketime    uint8    `json:"baketime"`
	Vegan       bool     `json:"vegan"`
	Author      string   `json:"author"`
	Rating      int      `json:"rating"`
}

var server_running bool = false // flag indicating server has started

// Function to return all recipes
func getRecipes(c *gin.Context) {
	var recipes []Recipe
	if dbRecipes._Size == 0 {
		c.JSON(http.StatusNotFound, recipes)
	} else {
		for _, recipe := range dbRecipes._Recipes {
			recipes = append(recipes, *recipe)
		}
		c.JSON(http.StatusOK, recipes)
	}
}

// Function to delete all recipes
func deleteAll(c *gin.Context) {
	dbRecipes.deleteAll()
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
	dbRecipes.insert(&newRecipe) // updated backend

	// Returns status to client and the recipe just created
	c.IndentedJSON(http.StatusCreated, newRecipe)
}

// Get a recipe by name or id
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
	oldPtr, err := dbRecipes.query_title(title)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Recipe not found"})
		return
	}

	*oldPtr = updated
	c.JSON(http.StatusOK, gin.H{
		"message": "Recipe Updated.",
	})
}

// Delete recipe by name
func deleteRecipeByQuery(c *gin.Context) {
	title, title_ok := c.GetQuery("title")
	id, id_ok := c.GetQuery("id")
	if title_ok {
		recipe, err := dbRecipes.query_title(title)
		if err == nil {
			dbRecipes.delete(recipe)
			c.JSON(http.StatusOK, gin.H{
				"message": "Recipe deleted.",
			})
			return
		}
	} else if id_ok {
		var id_uint uint64
		id_uint, _ = strconv.ParseUint(id, 10, 64)
		recipe, err := dbRecipes.query_id(id_uint)
		if err == nil {
			dbRecipes.delete(recipe)
			c.JSON(http.StatusOK, gin.H{
				"message": "Recipe deleted",
			})
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "missing title query"})
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "recipe not found"})

}

// function to start the server and initialize routes
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
	r.DELETE("/recipe", deleteRecipeByQuery)

	server_running = true
	r.Run("localhost:3000") // starts endpoint
	// NOTE: nothing below r.Run will execute
}

// lets view.go know if server is running
func ServerReady() bool {
	return server_running
}
