package types

type Recipe struct {
	Ingredients  []Ingredient
	Instructions []string
}

type Ingredient struct {
	Name     string
	Quantity string // Want this to be a string for thing like fractions
	Unit     string
}
