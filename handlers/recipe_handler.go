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

	render.Template(w, r, "all-recipes.go.html", &types.TemplateData{PageTitle: "All Recipes"})
}
