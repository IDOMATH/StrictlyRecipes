package main

import (
	"context"
	"fmt"
	"github.com/IDOMATH/StrictlyRecipes/db"
	"github.com/IDOMATH/StrictlyRecipes/handlers"
	"github.com/IDOMATH/StrictlyRecipes/router"
	"github.com/IDOMATH/StrictlyRecipes/types"
	"github.com/IDOMATH/StrictlyRecipes/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

func main() {
	port := util.EnvOrDefault("PORT", "8080")

	repo := run()
	
	http.HandleFunc("/", repo.Router.Route)

	fmt.Println("Server running on port: ", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

func run() *types.Repository {

	mongoDbUri := util.EnvOrDefault("MONGODBURI", "mongodb://localhost:27017")
	mongoDbName := util.EnvOrDefault("MONGODBNAME", "mongoRecipes")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoDbUri))
	if err != nil {
		log.Fatal(err)
	}

	repository := types.NewRepository()

	recipeHandler := handlers.NewRecipeHandler(db.NewRecipeStore(client, mongoDbName))

	repository.RH = recipeHandler

	repository.Router = router.NewRouter()

	return repository
}
