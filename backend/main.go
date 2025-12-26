package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/IDOMATH/StrictlyRecipes/db"
	"github.com/IDOMATH/StrictlyRecipes/handlers"
	"github.com/IDOMATH/StrictlyRecipes/middleware"
	"github.com/IDOMATH/StrictlyRecipes/repository"
	"github.com/IDOMATH/StrictlyRecipes/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	port := util.EnvOrDefault("PORT", "8080")

	router := http.NewServeMux()

	stack := middleware.CreateStack(
		middleware.Logger,
		middleware.Authenticate)

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: stack(router),
	}

	repo := run()

	router.HandleFunc("GET /", repo.HandleHome)
	router.HandleFunc("GET /recipes", repo.RH.HandleGetAllRecipes)
	router.HandleFunc("GET /recipes/{id}", repo.RH.HandleGetRecipeById)

	router.HandleFunc("GET /new-recipe", repo.RH.HandleNewRecipeForm)
	router.HandleFunc("POST /new-recipe", repo.RH.HandlePostRecipe)

	router.HandleFunc("GET /authors", repo.RH.HandleGetAuthors)
	router.HandleFunc("GET /authors/{id}", repo.RH.HandleGetAuthorById)

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

	dbHost := util.EnvOrDefault("DBHOST", "localhost")
	dbPort := util.EnvOrDefault("DBPORT", "5432")
	dbName := util.EnvOrDefault("DBNAME", "portfolio")
	dbUser := util.EnvOrDefault("DBUSER", "postgres")
	dbPass := util.EnvOrDefault("DBPASS", "postgres")
	dbSsl := util.EnvOrDefault("DBSSL", "disable")

	connectionString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s", dbHost, dbPort, dbName, dbUser, dbPass, dbSsl)
	fmt.Println("Connecting to Postgres")
	postgresDb, err := db.ConnectSQL(connectionString)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to Postgres")

	repository := repository.NewRepository()

	recipeHandler := handlers.NewRecipeHandler(db.NewRecipeStore(client, mongoDbName))

	repository.RH = recipeHandler

	return repository
}
