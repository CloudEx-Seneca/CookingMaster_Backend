package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Recipe struct defines the structure of a recipe
type Recipe struct {
	ID           int      `json:"id"`
	Title        string   `json:"title"`
	Ingredients  []string `json:"ingredients"`
	Instructions string   `json:"instructions"`
}

var recipes []Recipe
var nextRecipeID = 1

// Create a new recipe
func createRecipe(w http.ResponseWriter, r *http.Request) {
	var recipe Recipe
	err := json.NewDecoder(r.Body).Decode(&recipe)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	recipe.ID = nextRecipeID
	nextRecipeID++
	recipes = append(recipes, recipe)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(recipe)
}

// Get all recipes
func getRecipes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recipes)
}

// Get a single recipe by ID
func getRecipe(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid recipe ID", http.StatusBadRequest)
		return
	}

	for _, recipe := range recipes {
		if recipe.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(recipe)
			return
		}
	}

	http.Error(w, "Recipe not found", http.StatusNotFound)
}

// Main function to start the server
func main() {
	router := mux.NewRouter()

	// Define API routes
	router.HandleFunc("/recipes", createRecipe).Methods("POST")
	router.HandleFunc("/recipes", getRecipes).Methods("GET")
	router.HandleFunc("/recipes/{id}", getRecipe).Methods("GET")

	// Start the server
	fmt.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", router)
}
