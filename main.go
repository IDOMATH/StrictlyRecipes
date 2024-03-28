package main

import (
	"context"
	"fmt"
	"github.com/IDOMATH/StrictlyRecipes/db"
	"github.com/IDOMATH/StrictlyRecipes/handlers"
	"github.com/IDOMATH/StrictlyRecipes/repository"
	"github.com/IDOMATH/StrictlyRecipes/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

func main() {
	port := util.EnvOrDefault("PORT", "8080")

	router := http.NewServeMux()

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: router,
	}

	repo := run()

	router.HandleFunc("GET /", repo.HandleHome)
	//http.HandleFunc("/", repo.Route)

	fmt.Println("Server running on port: ", port)
	log.Fatal(server.ListenAndServe())
	//http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

func run() *repository.Repository {

	mongoDbUri := util.EnvOrDefault("MONGODBURI", "mongodb://localhost:27017")
	mongoDbName := util.EnvOrDefault("MONGODBNAME", "mongoRecipes")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoDbUri))
	if err != nil {
		log.Fatal(err)
	}

	repository := repository.NewRepository()

	recipeHandler := handlers.NewRecipeHandler(db.NewRecipeStore(client, mongoDbName))

	repository.RH = recipeHandler

	return repository
}
