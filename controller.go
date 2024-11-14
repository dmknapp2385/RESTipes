package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	s "winners.com/recipes/Server"
)

type RecipeController struct {
	BaseURL string
}

// Get all recipes from the backend
func (rc *RecipeController) GetRecipes() ([]s.Recipe, error) {
	// Make an HTTP GET request to the endpoint
	resp, err := http.Get(rc.BaseURL + "/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check if the request was successful
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("recipe not found: %s", resp.Status)
	}

	// Decode the response into the Recipe slice
	var recipes []s.Recipe
	if err := json.NewDecoder(resp.Body).Decode(&recipes); err != nil {
		return nil, err
	}

	return recipes, nil
}

// Get a single recipe by name
func (rc *RecipeController) GetRecipeByName(name string) (*s.Recipe, error) {
	// ensure the name is properly encoded and replace spaces with +
	name = url.QueryEscape(strings.ToLower(name))

	// Make an HTTP GET request with the name as a query parameter
	resp, err := http.Get(rc.BaseURL + "/recipe?title=" + name)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check if we got a valid recipe object or an error message
	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("recipe not found")
	}

	// Read the raw response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Try unmarshalling into Recipe struct
	var recipe s.Recipe
	if err := json.Unmarshal(body, &recipe); err != nil {
		return nil, fmt.Errorf("failed to unmarshal recipe: %v", err)
	}

	return &recipe, nil
}

// Get a single recipe by name
func (rc *RecipeController) GetRecipeByID(id uint64) (*s.Recipe, error) {
	// Make an HTTP GET request with the name as a query parameter
	var url strings.Builder
	url.WriteString(rc.BaseURL)
	url.WriteString("/recipe?id=")
	url.WriteString(fmt.Sprintf("%d", id))
	resp, err := http.Get(url.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check if we got a valid recipe object or an error message
	if resp.StatusCode == http.StatusNotFound {
		return nil, errors.New("recipe not found")
	}

	// Read the raw response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Try unmarshalling into Recipe struct
	var recipe s.Recipe
	if err := json.Unmarshal(body, &recipe); err != nil {
		return nil, fmt.Errorf("failed to unmarshal recipe: %v", err)
	}

	return &recipe, nil
}

// Update an existing recipe by name using a PUT request
func (rc *RecipeController) UpdateRecipe(name string, updatedRecipe s.Recipe) error {
	// ensure the name is properly encoded and replace spaces with +
	name = url.QueryEscape(name)

	// Convert the updated Recipe struct to JSON
	jsonData, err := json.Marshal(updatedRecipe)
	if err != nil {
		return err
	}

	// Create a PUT request
	req, err := http.NewRequest(http.MethodPut, rc.BaseURL+"/recipe?title="+name, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	// Send the PUT request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check if the update was successful
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("failed to update recipe: %s", string(body))
	}

	return nil
}

// Delete all recipes from the backend
func (rc *RecipeController) DeleteAllRecipes() error {
	// Create an HTTP DELETE request
	req, err := http.NewRequest(http.MethodDelete, rc.BaseURL+"/", nil)
	if err != nil {
		return err
	}

	// Send the DELETE request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check if the deletion was successful
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("failed to delete recipes: %s", string(body))
	}

	return nil
}

// delete given recipe
func (rc *RecipeController) DeleteRecipeByName(name string) error {
	// ensure the name is properly encoded and replace spaces with +
	name = url.QueryEscape(name)

	// Create an HTTP DELETE request
	req, err := http.NewRequest(http.MethodDelete, rc.BaseURL+"/recipe?title="+name, nil)
	if err != nil {
		return err
	}

	// Send the DELETE request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check if the deletion was successful
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("failed to delete recipes: %s", string(body))
	}

	return nil
}

// adds new recipe using POST
func (rc *RecipeController) createRecipe(newRecipe s.Recipe) error {
	// Convert the updated Recipe struct to JSON
	jsonData, err := json.Marshal(newRecipe)
	if err != nil {
		return err
	}

	// Create a POST request
	req, err := http.NewRequest(http.MethodPost, rc.BaseURL+"/recipe", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	// Send the POST request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check if the create was successful
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("failed to create recipe: %s", string(body))
	}

	return nil
}