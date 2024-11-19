package server

// database.go
// A GOlang database implementation.

import (
	"errors"
	"strings"
	"sync"
)

// database table implementation
type RecipeDatabase struct {
	_CurrentID    uint64             // uid counter
	_Recipes      map[uint64]*Recipe // primary data structure (primary key)
	_Column_Title map[string]*Recipe // secondary data structure (secondary key)
	_Size         uint64             // number of rows in table
	sync.RWMutex
}

// database schema for recipe database
var dbRecipes RecipeDatabase = RecipeDatabase{
	_CurrentID:    0,
	_Recipes:      make(map[uint64]*Recipe),
	_Column_Title: make(map[string]*Recipe),
	_Size:         0,
}

// deletes all recipes from database
func (db *RecipeDatabase) deleteAll() {
	db._CurrentID = 0
	db._Recipes = make(map[uint64]*Recipe)
	db._Column_Title = make(map[string]*Recipe)
	db._Size = 0
}

// inserts a recipe into the database
func (db *RecipeDatabase) insert(r *Recipe) {
	var title string = r.Title
	r.Title = strings.ToUpper(title)
	db.Lock()
	db._CurrentID++
	r.ID = uint32(db._CurrentID)
	db._Recipes[db._CurrentID] = r
	db._Column_Title[r.Title] = r
	db._Size++
	db.Unlock()
}

// removes a recipe from the database
func (db *RecipeDatabase) delete(r *Recipe) {
	db.Lock()
	delete(db._Recipes, uint64(r.ID))
	delete(db._Column_Title, r.Title)
	db._Size--
	db.Unlock()
}

// selects a recipe with the provided title from the database
func (db *RecipeDatabase) query_title(title string) (*Recipe, error) {
	var recipe *Recipe
	var found bool
	db.RLock()
	recipe, found = db._Column_Title[strings.ToUpper(title)]
	db.RUnlock()
	if !found {
		return nil, errors.New("recipe not found")
	}
	return recipe, nil
}

// selects a recipe with the provided id from the database
func (db *RecipeDatabase) query_id(id uint64) (*Recipe, error) {
	var recipe *Recipe
	var found bool
	db.RLock()
	recipe, found = db._Recipes[id]
	db.RUnlock()
	if !found {
		return nil, errors.New("recipe not found")
	}
	return recipe, nil
}
