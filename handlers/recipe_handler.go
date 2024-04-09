package handlers

import (
	"context"
	"fmt"
	"github.com/IDOMATH/StrictlyRecipes/db"
	"github.com/IDOMATH/StrictlyRecipes/render"
	"github.com/IDOMATH/StrictlyRecipes/types"
	"net/http"
)

type RecipeHandler struct {
	recipeStore db.RecipeStore
}

func NewRecipeHandler(recipeStore db.RecipeStore) *RecipeHandler {
	return &RecipeHandler{recipeStore: recipeStore}
}

func (h *RecipeHandler) HandleGetAllRecipes(w http.ResponseWriter, r *http.Request) {
	c, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	recipes, err := h.recipeStore.GetAllRecipes(c)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	objects := make(map[string]interface{})
	objects["recipes"] = recipes

	render.Template(w, r, "all-recipes.go.html", &types.TemplateData{PageTitle: "All Recipes", ObjectMap: objects})
}

func (h *RecipeHandler) HandleGetRecipeById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	c, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	recipe, err := h.recipeStore.GetRecipeById(c, id)
	if err != nil {
		fmt.Println(err)
	}

	objects := make(map[string]interface{})
	objects["recipe"] = recipe

	render.Template(w, r, "recipe.go.html", &types.TemplateData{PageTitle: recipe.Title, ObjectMap: objects})
}

func (h *RecipeHandler) HandleNewRecipeForm(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "new-recipe-form.go.html", &types.TemplateData{PageTitle: "New Ricipe"})
}

func (h *RecipeHandler) HandlePostRecipe(w http.ResponseWriter, r *http.Request) {
	//TODO:
}

func (h *RecipeHandler) HandleGetAuthors(w http.ResponseWriter, r *http.Request) {

}

func (h *RecipeHandler) HandleGetAuthorById(w http.ResponseWriter, r *http.Request) {
	
}
