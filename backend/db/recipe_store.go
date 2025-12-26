package db

import (
	"context"
	"fmt"

	"github.com/IDOMATH/StrictlyRecipes/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const recipeCollection = "recipes"

type RecipeStore interface {
	Dropper

	InsertRecipe(context.Context, *types.Recipe) (*types.Recipe, error)
	GetAllRecipes(context.Context) ([]*types.Recipe, error)
	GetRecipeById(context.Context, string) (*types.Recipe, error)
	GetRecipesByAuthor(context.Context, string) ([]*types.Recipe, error)
}

type MongoRecipeStore struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func (s *MongoRecipeStore) Drop(ctx context.Context) error {
	fmt.Println("--- Dropping recipe collection")
	return s.collection.Drop(ctx)
}

func NewRecipeStore(client *mongo.Client, dbName string) *MongoRecipeStore {
	return &MongoRecipeStore{
		client:     client,
		collection: client.Database(dbName).Collection(recipeCollection),
	}
}

func (s *MongoRecipeStore) InsertRecipe(ctx context.Context, recipe *types.Recipe) (*types.Recipe, error) {
	res, err := s.collection.InsertOne(ctx, recipe)
	if err != nil {
		return nil, err
	}
	recipe.Id = res.InsertedID.(primitive.ObjectID)
	return recipe, nil
}

func (s *MongoRecipeStore) GetAllRecipes(ctx context.Context) ([]*types.Recipe, error) {
	res, err := s.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var recipes []*types.Recipe
	if err := res.All(ctx, &recipes); err != nil {
		return nil, err
	}
	return recipes, nil
}

func (s *MongoRecipeStore) GetRecipeById(ctx context.Context, id string) (*types.Recipe, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var recipe types.Recipe
	if err := s.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&recipe); err != nil {
		return nil, err
	}
	return &recipe, nil
}

func (s *MongoRecipeStore) GetRecipesByAuthor(ctx context.Context, author string) ([]*types.Recipe, error) {
	res, err := s.collection.Find(ctx, bson.M{"author": author})
	if err != nil {
		return nil, err
	}
	var recipes []*types.Recipe
	if err := res.All(ctx, &recipes); err != nil {
		return nil, err
	}

	return recipes, nil
}

func (s *MongoRecipeStore) GetAllAuthors(ctx context.Context) ([]string, error) {
	res, err := s.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var authors []string
	if err := res.All(ctx, &authors); err != nil {
		return nil, err
	}
	return authors, nil
}
