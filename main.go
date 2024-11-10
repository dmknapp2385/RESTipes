package main

import s "winners.com/recipes/Server"
	


func main(){
	//This will be where the view is
	go s.StartServer() // go routine so tests can be done concurrently
	// s.StartServer()

	// temporary for testing controller
	controller := RecipeController{BaseURL: "http://localhost:3000"}
	getRecipesTest(&controller)

	select{} // blocks to goroutine from exiting so the server can continue running
}