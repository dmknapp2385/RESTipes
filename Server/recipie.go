package server

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type recipie struct {
	ID          int      `json :"id"`
	Title       string   `json:"title"`
	Ingredients []string `json: "ingredients"`
	Steps       []string `json: "steps"`
	Baketime    int      `json : "baketime"`
	Vegan       bool     `json: "vegan"`
	Author      string   `json: "author"`
	Rating      int      `json :"rating"`

    
}

var recipieBook = []recipie{{ID:1, Title:"Meat Loaf", Ingredients: []string{"beef", "breadcrumbs", "spices"}, Baketime: 120, Vegan: false, Author: "Elle", Rating: 5 }}


//function to return all recipies
func getRecipies(c *gin.Context) {
	c.JSON(http.StatusOK,recipieBook)
}

//function to delete all recipies
func deleteAll(c *gin.Context){
	recipieBook = []recipie{}
	c.JSON(http.StatusOK, gin.H{
		"messag":"All recipies deleted.",
	})
}

//function to create a new recipie
func createRecipie(c *gin.Context){
	var newRecipie recipie

	//attemtp to bind newRecipie to JSON information in the Context
	if err := c.BindJSON(&newRecipie); err != nil{
		return // will return http errorstatus as default of BindJSON method if error arises
	}

	//add recipie to recpie slice
	recipieBook = append(recipieBook, newRecipie)

	// returns status to client and the recipie just created
	c.IndentedJSON(http.StatusCreated, newRecipie)
}

//helper to get recipie by name
func getRecipie(t string)(*recipie, error){
	for index,recipie := range recipieBook{
		if recipie.Title == t{
			return &recipieBook[index], nil
		}
	}
	return nil, errors.New("recipie not found")
}


//get a recipie by name
func getRecipieByName(c *gin.Context){
	title, ok := c.GetQuery("title")
	
	if ok == false{
		c.JSON(http.StatusBadRequest, gin.H{"message": "missing title query"})
	}

	recipie, err := getRecipie(title)

	if err != nil{
		c.JSON(http.StatusNotFound, gin.H{"messge":"Recipie not found"})
	}

	c.JSON(http.StatusOK, recipie)
}

//update recipie
func updateRecpie(c *gin.Context){
	title, ok := c.GetQuery("title")

	if ok == false {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Url missing querey"})
	}

	var updated *recipie

	if err := c.BindJSON(&updated); err != nil{
        //return http errorstatus as default of BindJSON method if err
		return
	}

	oldPtr, err := getRecipie(title)

	if err != nil{
		c.JSON(http.StatusNotFound, gin.H{"messge":"Recipie not found"})
	}
	
	 *oldPtr = *updated

    c.JSON(http.StatusOK, gin.H{
        "messge":"Recipie Updated.",
    })
}


func StartServer(){
	r:= gin.Default()
	r.GET("/", getRecipies)
	r.DELETE("/", deleteAll)
	r.POST("/recipie", createRecipie)
	r.GET("/recipie", getRecipieByName)
    r.PUT("/recipie", updateRecpie)
	

	r.Run("localhost:3000") // listen and serv on port 3000
}

