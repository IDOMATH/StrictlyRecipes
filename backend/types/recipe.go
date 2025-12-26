package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type Recipe struct {
	Id           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title        string             `bson:"title" json:"title"`
	Ingredients  []Ingredient       `bson:"ingredients" json:"ingredients"`
	Instructions []Instruction      `bson:"instructions" json:"instructions"`
	Author       string             `bson:"author" json:"author"`
	Thumbnail    string             `bson:"thumbnail" json:"thumbnail"`
}

type Ingredient struct {
	Name     string
	Quantity string // Want this to be a string for thing like fractions
	Unit     string
}

type Instruction struct {
	ImageLocation string
	Text          string
}
