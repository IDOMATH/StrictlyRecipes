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
	"github.com/IDOMATH/session/memorystore"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	port := util.EnvOrDefault("PORT", "8080")

	router := http.NewServeMux()

	repo := run()

	stack := middleware.CreateStack(
		middleware.Logger,
		middleware.Authenticate(repo))

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: router,
	}

	router.HandleFunc("GET /", middleware.Use(repo.HandleHome, stack))
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

	repository.Session = memorystore.New[string]()

	recipeHandler := handlers.NewRecipeHandler(db.NewRecipeStore(client, mongoDbName))
	userHandler := handlers.NewUserHandler(db.NewUserStore(postgresDb.SQL))

	repository.RH = recipeHandler
	repository.UH = userHandler

	return repository
}
