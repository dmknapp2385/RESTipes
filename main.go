package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	s "winners.com/recipes/Server"
)

var controller = RecipeController{BaseURL: "http://localhost:3000"}
var reader = bufio.NewReader(os.Stdin)


func main(){
    
	go s.StartServer() // go routine so tests can be done concurrently

	// temporary for testing controller
	// getRecipesTest(&controller)

   for{
    
    fmt.Println("What would you like to do:\n1. Get all recipes\n2. Get Recipe by name\n3. Add Recipe\n4. Update recipe\n5. Delete all recipes? Press any other key to exit.")
    input, _ := reader.ReadString('\n')
    input = strings.TrimSpace(input)


    //loop for input
    if input == "1"{
        recipes, _ := controller.GetRecipes()
        output(recipes)
    }else if input == "2"{
        recipe := getRecipe()
        if recipe != nil{
            printRecipe(recipe)
        }        
    }else if input == "3"{
        addRecipe()
    }else if input == "4"{
        updateRecipe()
    }else if input =="5"{
        controller.DeleteAllRecipes()
    } else{
        os.Exit(3)
    }
   }

   //this is never reached but seems to be working without it???
//    select{} // blocks to goroutine from exiting so the server can continue running

}

func getRecipe() (*s.Recipe){
    fmt.Println("Which recipe would you like to get?")

    input, _ := reader.ReadString('\n')
    input = strings.TrimSpace(input)
    
    recipe,err := controller.GetRecipeByName(input)

    if err != nil{
        fmt.Println("Could not find recipe by that name")
        return nil
    } 
    return recipe
}

func addRecipe(){
    recipe := s.Recipe{}
    fmt.Println("Title: ")
    title, _ := reader.ReadString('\n')
    title = strings.TrimSpace(title)
    recipe.Title = title
    fmt.Println("Author: ")
    author, _ := reader.ReadString('\n')
    author = strings.TrimSpace(author)
    recipe.Author = author
    fmt.Println("Ingredients (endter a line separated list): ")
    ingredients, _ := reader.ReadString('\n')
    ingredients = strings.TrimSpace(ingredients)
    ingredList := strings.Split(ingredients,",")
    recipe.Ingredients = ingredList
    fmt.Println("Steps (endter a line separated list): ")
    steps, _ := reader.ReadString('\n')
    steps = strings.TrimSpace(steps)
    stepList := strings.Split(steps,",")
    recipe.Steps = stepList
    
}

func updateRecipe(){
    recipe:= getRecipe()
    if recipe == nil{
        return
    }
    name := recipe.Title
    fmt.Println("What would you like to update\n1.Title\n2.Author\n3.Ingredients\n4.Steps?")
    input, _ := reader.ReadString('\n')
    input = strings.TrimSpace(input)
    if input == "1"{
        fmt.Println("New Title: ")
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)
        recipe.Title=input
    }else if input == "2"{
        fmt.Println("New Author: ")
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)
        recipe.Author=input
    }else if input == "3"{
        fmt.Println("Enter comma separated list of ingredients: ")
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)
        ingredients:= strings.Split(input,",")
        recipe.Ingredients = ingredients
    }else if input == "4"{
        fmt.Println("Enter comma separated list of steps: ")
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)
        steps:= strings.Split(input,",")
        recipe.Ingredients = steps
    }else {
        fmt.Println("Did not recoginze command")
        return
    }

    controller.UpdateRecipe(name,*recipe)
}


//funciton prints singlular recipe with all information
func printRecipe(r *s.Recipe){
    outputStr := "Title: " + r.Title + "\nAuthor: " + r.Author + "\nIngredients:\n"

        for _,ingredient := range r.Ingredients{
            outputStr += "\t\u2022" + ingredient + " \n"
        }
        outputStr += "Steps:\n"
        for i,step := range r.Steps{
            outputStr += "\t" + strconv.Itoa(i+1) + ": " + step + " \n"
        }
        outputStr += "Rating: " + string(r.Rating)
        fmt.Println(outputStr)
}

//function takes a recipe slice and outputs in readable format
func output (recipies []s.Recipe){
    for _, recipe:=range recipies{
        fmt.Println(recipe.Title)
    }
    
}