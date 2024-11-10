package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	s "winners.com/recipes/Server"
)



func main(){
	go s.StartServer() // go routine so tests can be done concurrently
	// s.StartServer()

	// temporary for testing controller
	controller := RecipeController{BaseURL: "http://localhost:3000"}
	// getRecipesTest(&controller)
   

   for{
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("What would you like to do: 1. Get all recipes \n 2.Get Recipe by name \n 3.Add Recipe\n4. Update recipe\n5. Delete all recipes? Press any other key to exit.")
    input, _ := reader.ReadString('\n')
    input = strings.TrimSpace(input)

    //loop for input
    if input == "1"{
        recipes, _ := controller.GetRecipes()
        output(recipes)
    }else if input == "2"{
        recipe := getRecipe()
        output([]s.Recipe{*recipe})
    }else if input == "3"{
        addRecipe()
    }else if input == "4"{
        updateRecipe()
    }else if input =="5"{
        controller.DeleteAllRecipes()
    } else{
        break
    }
   }
    
   select{} // blocks to goroutine from exiting so the server can continue running

}

func getRecipe() (*s.Recipe){
    
}

func addRecipe(){

}

func updateRecipe(){
    recipe:= getRecipe()

}


//function takes a recipe slice and outputs in readable format
func output (recipies []s.Recipe){
    for _, recipe:=range recipies{
        outputStr := "Title: " + recipe.Title + "\n Author: " + recipe.Author + "\n Ingredients:\n"

        for _,ingredient := range recipe.Ingredients{
            outputStr += "\t\u2022" + ingredient + " \n"
        }
        outputStr += "Steps:\n"
        for i,step := range recipe.Steps{
            outputStr += "\t" + strconv.Itoa(i+1) + ": " + step + " \n"
        }
        outputStr += "Rating: " + string(recipe.Rating)
        fmt.Println(outputStr)
    }
    
}