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
	Url string
	Category
	Selector  string
	Title     string
	Img       []string
	Paginator string
	Recipe
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
	Burger Category = "Burger or BBQ"
	Asian  Category = "Asian"
)
