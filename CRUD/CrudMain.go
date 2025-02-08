package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
)

type Recipe struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Ingredients string `json:"ingredients"`
	Instructions string `json:"instructions"`
}

var recipes []Recipe
var nextID = 1

func getRecipes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recipes)
}

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

func createRecipe(w http.ResponseWriter, r *http.Request) {
	var recipe Recipe
	err := json.NewDecoder(r.Body).Decode(&recipe)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	recipe.ID = nextID
	nextID++
	recipes = append(recipes, recipe)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(recipe)
}

func updateRecipe(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid recipe ID", http.StatusBadRequest)
		return
	}

	for i, recipe := range recipes {
		if recipe.ID == id {
			var updatedRecipe Recipe
			err := json.NewDecoder(r.Body).Decode(&updatedRecipe)
			if err != nil {
				http.Error(w, "Invalid request payload", http.StatusBadRequest)
				return
			}

			updatedRecipe.ID = id
			recipes[i] = updatedRecipe

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedRecipe)
			return
		}
	}

	http.Error(w, "Recipe not found", http.StatusNotFound)
}

func deleteRecipe(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid recipe ID", http.StatusBadRequest)
		return
	}

	for i, recipe := range recipes {
		if recipe.ID == id {
			recipes = append(recipes[:i], recipes[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Recipe not found", http.StatusNotFound)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/recipes", getRecipes).Methods("GET")
	router.HandleFunc("/recipes/{id}", getRecipe).Methods("GET")
	router.HandleFunc("/recipes", createRecipe).Methods("POST")
	router.HandleFunc("/recipes/{id}", updateRecipe).Methods("PUT")
	router.HandleFunc("/recipes/{id}", deleteRecipe).Methods("DELETE")

	fmt.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", router)
}
