package main

import (
	"fmt"
)

func getRecipesTest(controller *RecipeController) {
	getAllRecipesTest(controller)
	getRecipeByNameTest(controller, "Meat Loaf")
	getNonexistentRecipeTest(controller)
}

func getAllRecipesTest(controller *RecipeController) {
	recipes, err := controller.GetRecipes()
	if err != nil {
		fmt.Println("Error getting recipes:", err)
	} else {
		fmt.Println("Recipes:", recipes)
	}
}

func getRecipeByNameTest(controller *RecipeController, name string) {
	recipes, err := controller.GetRecipeByName(name)
	if err != nil {
		fmt.Println("Error getting recipes:", err)
	} else {
		fmt.Println("Recipes:", recipes)
	}
}

func getNonexistentRecipeTest(controller *RecipeController) {
	recipes, err := controller.GetRecipeByName("bad name")
	if err != nil {
		fmt.Println("Error getting recipes:", err)
	} else {
		fmt.Println("Recipes:", recipes)
	}
}