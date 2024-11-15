package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"

	s "winners.com/recipes/Server"
)

//view.go file prompts and accepts input from the user to pass onto the controller
//Prompts include get all recipes, add recipe, get recipe by name, delete all recipes,
// get a recipe by name, and updating a current recipe.
//The view on load will upload a json file of recipes if a file name was passed
//as a flag in the command line with the exectuion of the executable file.
//command would be ./recipes -file=filename.json
//program quits with input that is not one of the options for input

//switch case constants
const GET_ALL_RECIPES = "1"
const GET_RECIPE_BY_TITLE = "2"
const POST_NEW_RECIPE = "3"
const UPDATE_RECIPE = "4"
const DELETE_ALL_RECIPES = "5"
const DELETE_BY_NAME="6"


var reader = bufio.NewReader(os.Stdin)
var controller = RecipeController{BaseURL: "http://localhost:3000"}

func view_prompt(wait_group *sync.WaitGroup, file *string) {
    //if file is not an empty string, read in the file
	if *file != "" {
		data, err := os.ReadFile(*file)
		if err != nil {
			fmt.Println("File not found!")
		}

		var recipes []s.Recipe
		err = json.Unmarshal(data, &recipes)
        
		if err != nil {
			fmt.Println("File must be json")
		}

        //add all recipes in file to database
		for _, recipe := range recipes {
            recipe.Title = strings.ToTitle(recipe.Title)
			controller.createRecipe(recipe)
			fmt.Println(recipe)
		}
	}
    
    //variable to check for exit
	var run_prompt bool = true

    //while not exit, prompt user
	for {
		fmt.Println("What would you like to do:\n1. Get all recipes\n2. Get Recipe by name\n3. Add Recipe\n4. Update recipe\n5. Delete all recipes.\n6. Delete recipe by name? Press any other key to exit.")
		fmt.Print(">>> ")
		input, _ := reader.ReadString('\n')

		switch strings.TrimSpace(input) {
		case GET_ALL_RECIPES:
			get_recipes()

		case GET_RECIPE_BY_TITLE:
			get_recipes()

			fmt.Println("\nWhich recipe would you like to get?")
			fmt.Print(">>> ")
			input, _ := reader.ReadString('\n')
			title := strings.TrimSpace(input)
            title = strings.ToTitle(title)

			get_recipe(title)

		case POST_NEW_RECIPE:
			addRecipe()
		case UPDATE_RECIPE:
            get_recipes()

			fmt.Println("\nWhich recipe would you like to get?")
			fmt.Print(">>> ")

            input, _ := reader.ReadString('\n')
			title := strings.TrimSpace(input)
            title = strings.ToTitle(title)

			updateRecipe(title)
		case DELETE_ALL_RECIPES:
			controller.DeleteAllRecipes()
		case DELETE_BY_NAME:
            get_recipes()

			fmt.Println("\nWhich recipe would you like to delete?")
			fmt.Print(">>> ")
			input, _ := reader.ReadString('\n')
			title := strings.TrimSpace(input)
            title = strings.ToTitle(title)

			if error:=controller.DeleteRecipeByName(title); error != nil{
                fmt.Println("Could not delete recipe.\n")
            }else{
                fmt.Println("Recipe deleted.\n")
            }

		default:
			run_prompt = false
			fmt.Println("Closing prompt...")
		}

		if !run_prompt {
			break
		}
	}

	fmt.Println("goodbye!")
	wait_group.Done()
}

//get all recipes
func get_recipes() {
	recipes, err := controller.GetRecipes()
	if err != nil {
		fmt.Println("\nResults: Could not find any recipes.")
	} else {
		fmt.Println("\nResults:")
		for i, recipe := range recipes {
			fmt.Printf(" (%v) %s\n", i, strings.ToTitle(recipe.Title))
		}
	}
}

//get recipe and print from controller
func get_recipe(title string) {
	recipe, err := controller.GetRecipeByName(title)
	if err != nil {
		fmt.Println("Could not find recipe by that name")
	} else {
		printRecipe(recipe)
	}
}

//add a recipe
func addRecipe() {
	recipe := s.Recipe{}
	fmt.Print("Title: ")
	//read input and convert to uppercase
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)
	title = strings.ToTitle(title)
	recipe.Title = title
	fmt.Print("Author: ")
	author, _ := reader.ReadString('\n')
	author = strings.TrimSpace(author)
	recipe.Author = author
	recipe.Ingredients = getIngredients()
	recipe.Steps = getSteps()
    fmt.Print("Bake time: ")
	time, _ := reader.ReadString('\n')
	time = strings.TrimSpace(time)
    time_int,_ :=strconv.ParseUint(time, 10, 64)
    recipe.Baketime = uint8(time_int)
    fmt.Print("Rating: ")
	rate, _ := reader.ReadString('\n')
	rate = strings.TrimSpace(rate)
    rate_int, _ := strconv.Atoi(rate)
    recipe.Rating = rate_int 
	controller.createRecipe(recipe)
}

//get list of ingredients
func getIngredients() (ingredients []string){
    fmt.Println("Ingredients: Type the ingredient then press enter.")
	fmt.Println("When finished, press enter key only.")
    var counter uint8 = 1
	for {
		fmt.Printf("Ingredient %v: ", counter)
		counter++
		ingredient, _ := reader.ReadString('\n')
		ingredient = strings.TrimSpace(ingredient)
		ingredient = strings.ToLower(ingredient)
		if ingredient == "" {
			break
		}
		ingredients = append(ingredients, ingredient)
	}
    return ingredients
}

//get list of steps
func getSteps() (steps []string){
    counter := 1
	fmt.Println("Steps: Type the steps then press enter.")
	fmt.Println("When finished, press enter key only.")
	for {
		fmt.Printf("Step %v: ", counter)
		counter++
		step, _ := reader.ReadString('\n')
		step = strings.TrimSpace(step)
		step = strings.ToLower(step)
		if step == "" {
			break
		}
		steps = append(steps, step)
	}
    return steps
}

//updates recipe
func updateRecipe(title string) {
    recipe, _ := controller.GetRecipeByName(title)
	if recipe == nil {
		fmt.Println("Could not find recipe by that name")
	}
	name := recipe.Title
	fmt.Println("What would you like to update\n1.Title\n2.Author\n3.Ingredients\n4.Steps\n5.Baketime\n6.Rating?")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input == "1" {
		fmt.Println("New Title: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
        input = strings.ToTitle(input)
		recipe.Title = input
	} else if input == "2" {
		fmt.Println("New Author: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		recipe.Author = input
	} else if input == "3" {
        ingredients := getIngredients()
		recipe.Ingredients = ingredients
	} else if input == "4" {
		steps := getSteps()
		recipe.Steps = steps
	} else if input == "5" {
		fmt.Println("Enter the bake time: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		baketime_int, _ := strconv.ParseUint(input, 10, 64)
		recipe.Baketime = uint8(baketime_int)
	} else if input == "6" {
		fmt.Println("Enter the rating: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		rate_int, _ := strconv.Atoi(input)
		recipe.Rating = rate_int
	} else {
		fmt.Println("Did not recoginze command")
		return
	}

	controller.UpdateRecipe(name, *recipe)
}

// funciton prints singlular recipe with all information
func printRecipe(r *s.Recipe) {
	outputStr := "Title: " + r.Title + "\nAuthor: " + r.Author + "\nIngredients:\n"

	for _, ingredient := range r.Ingredients {
		outputStr += "\t\u2022" + ingredient + " \n"
	}
	outputStr += "Steps:\n"
	for i, step := range r.Steps {
		outputStr += "\t" + strconv.Itoa(i+1) + ": " + step + " \n"
	}
    outputStr += "Bake time: " + strconv.Itoa(int(r.Baketime)) + "\n"
	outputStr += "Rating: " + strconv.Itoa(r.Rating) +"\n"
	fmt.Println(outputStr)
}

// function takes a recipe slice and outputs in readable format
func output(recipies []s.Recipe) {
	for _, recipe := range recipies {
		fmt.Println(recipe.Title)
        fmt.Println("\n")
	}

}
