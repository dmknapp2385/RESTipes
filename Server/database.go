package server

import (
	"errors"
	"sync"
)

type RecipeDatabase struct {
	_CurrentID    uint64
	_Recipes      map[uint64]*Recipe
	_Column_Title map[string]*Recipe
	_Size         uint64
	sync.RWMutex
}

var dbRecipes RecipeDatabase = RecipeDatabase{
	_CurrentID:    0,
	_Recipes:      make(map[uint64]*Recipe),
	_Column_Title: make(map[string]*Recipe),
	_Size:         0,
}

func(db *RecipeDatabase) deleteAll(){
	db._CurrentID = 0;
	db._Recipes =  make(map[uint64]*Recipe);
	db._Column_Title =  make(map[string]*Recipe);
	db._Size = 0;
}

func (db *RecipeDatabase) insert(r *Recipe) {
	db.Lock()
	db._CurrentID++
	r.ID = uint32(db._CurrentID);
	db._Recipes[db._CurrentID] = r
	db._Column_Title[r.Title] = r
	db._Size++
	db.Unlock()
}

func (db *RecipeDatabase) delete(r *Recipe) {
	db.Lock()
	delete(db._Recipes, db._CurrentID)
	delete(db._Column_Title, r.Title)
	db._Size--
	db.Unlock()
}

func (db *RecipeDatabase) query_title(title string) (*Recipe, error) {
	var recipe *Recipe
	var found bool
	db.RLock()
	// recipe, found = db._Column_Title[strings.ToLower(title)]
	recipe, found = db._Column_Title[title]
	db.RUnlock()
	if !found {
		return nil, errors.New("recipe not found")
	}
	return recipe, nil
}

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
