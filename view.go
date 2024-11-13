package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"

	s "winners.com/recipes/Server"
)

const GET_ALL_RECIPES = "1"
const GET_RECIPE_BY_TITLE = "2"
const POST_NEW_RECIPE = "3"
const UPDATE_RECIPE = "4"
const DELETE_ALL_RECIPES = "5"

var reader = bufio.NewReader(os.Stdin)
var controller = RecipeController{BaseURL: "http://localhost:3000"}

func view_prompt(wait_group *sync.WaitGroup, file *string) {

	if *file !=""{
		fmt.Println(*file)
	}
	var run_prompt bool = true
	for {
		fmt.Println("What would you like to do:\n1. Get all recipes\n2. Get Recipe by name\n3. Add Recipe\n4. Update recipe\n5. Delete all recipes? Press any other key to exit.")
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
			get_recipe(title)

		case POST_NEW_RECIPE:
			addRecipe()
		case UPDATE_RECIPE:
			updateRecipe()
		case DELETE_ALL_RECIPES:
			controller.DeleteAllRecipes()

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

func get_recipe(title string) {
	recipe, err := controller.GetRecipeByName(title)
	if err != nil {
		fmt.Println("Could not find recipe by that name")
	} else {
		fmt.Printf("'%s', by %s\n", recipe.Title, recipe.Author)
	}
}

func getRecipe() *s.Recipe {
	fmt.Println("Which recipe would you like to get?")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	recipe, err := controller.GetRecipeByName(input)

	if err != nil {
		fmt.Println("Could not find recipe by that name")
		return nil
	}
	return recipe
}

func addRecipe() {
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
	ingredList := strings.Split(ingredients, ",")
	recipe.Ingredients = ingredList
	fmt.Println("Steps (endter a line separated list): ")
	steps, _ := reader.ReadString('\n')
	steps = strings.TrimSpace(steps)
	stepList := strings.Split(steps, ",")
	recipe.Steps = stepList
}

func updateRecipe() {
	recipe := getRecipe()
	if recipe == nil {
		return
	}
	name := recipe.Title
	fmt.Println("What would you like to update\n1.Title\n2.Author\n3.Ingredients\n4.Steps?")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input == "1" {
		fmt.Println("New Title: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		recipe.Title = input
	} else if input == "2" {
		fmt.Println("New Author: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		recipe.Author = input
	} else if input == "3" {
		fmt.Println("Enter comma separated list of ingredients: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		ingredients := strings.Split(input, ",")
		recipe.Ingredients = ingredients
	} else if input == "4" {
		fmt.Println("Enter comma separated list of steps: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		steps := strings.Split(input, ",")
		recipe.Ingredients = steps
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
	outputStr += "Rating: " + string(r.Rating)
	fmt.Println(outputStr)
}

// function takes a recipe slice and outputs in readable format
func output(recipies []s.Recipe) {
	for _, recipe := range recipies {
		fmt.Println(recipe.Title)
	}

}
