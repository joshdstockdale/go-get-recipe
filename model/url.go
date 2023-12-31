package model

type Selector struct {
	title       string
	img         string
	ingredients []string
}

type Child struct {
	url       string
	selectors []Selector
}

type List struct {
	Selector  string
	Title     string
	Img       []string
	Url       string
	Paginator string
	Include   []string
	Exclude   []string
	Recipe
	Category
	//children  []Child
}

type Recipe struct {
	Selector    string
	Title       string
	Img         []string
	Url         string
	Ingredients []string
}

type Category string

const (
	Asian  Category = "Asian"
	Taco   Category = "Taco or Burrito"
	Pasta  Category = "Pasta"
	Fish   Category = "Fish"
	Burger Category = "Burger or BBQ"
	Soup   Category = "Soup"
)
