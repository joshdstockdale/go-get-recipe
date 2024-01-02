package handler

import (
	"get-recipe-inator/model"
	"get-recipe-inator/view/component"
	"get-recipe-inator/view/page"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/labstack/echo/v4"
)

type UrlHandler struct{}

var urls []model.List

func InitUrls() []model.List{
	return []model.List{
		{
			Selector: ".entry-header",
			Title: ".entry-title-link",
			Img: []string{"img"},
			Url: "https://thewoksoflife.com/category/recipes/chinese-take-out/",
			Exclude: []string{"Chicken", "Turkey"},
			Paginator: ".pagination ul li a",
			Category: model.Asian,
			Recipe: model.Recipe{
				Selector: "main",
				Title: ".entry-title",
				Img: []string{".wp-post-image"},
				Ingredients: []string{".wprm-recipe-ingredient"},
				Url: ".entry-title-link",
			},
		},
		{
			Selector: ".box-container .post_box",
			Title: ".headline",
			Img: []string{"img"},
			Url: "https://www.inspiredtaste.net/category/recipes/main-dishes/",
			Include: []string{"Pork", "Beef", "Burger", "Steak", "Ribs", "BBQ", "Chili"},
			Exclude: []string{"Chicken", "Turkey"},
			Category: model.Burger,
			Paginator: "a.page",
			Recipe: model.Recipe{
				Selector: ".post_box",
				Title: ".headline",
				Img: []string{".ytcover", "img"},
				Ingredients: []string{".itr-ingredients p"},
				Url: ".featured_image_link",
			},
		},
	}
}
func (u UrlHandler) HandleHome(cx echo.Context)error{
	categories := []model.Category{model.Asian, model.Burger, model.Soup}
	category := cx.Param("category")
	url := cx.Param("url")
	return render(cx, page.Index(categories, category, url))
}

func (u UrlHandler) HandleRecipes(cx echo.Context) error {
	urls := InitUrls()
	category := cx.QueryParam("category")
	// var allowed []string
	allowed := []string{"thewoksoflife.com", "www.inspiredtaste.net"}
	var matched []model.List
	for _, u := range urls {
		if u.Category == model.Category(category) {
			// allowed = append(allowed,u.Url)
			matched = append(matched, u)
		}
	}

	c := colly.NewCollector(
		colly.AllowedDomains(allowed...),
		colly.CacheDir("./list_cache"),
	)

	var recipes []model.Recipe
	var pages []string
	for _, u := range matched {
		c.OnHTML(u.Selector, func(e *colly.HTMLElement){
			recipes = getRecipes(u, recipes, e)
		})

		c.OnHTML(u.Paginator, func(e *colly.HTMLElement){
			pages = append(pages,e.Request.AbsoluteURL(e.Attr("href")))
		})

		c.Visit(u.Url)
	}
	for _, p := range pages {
			c.Visit(p)
	}
		
	return render(cx, component.List(recipes))	
}

func getRecipes(u model.List,recipes []model.Recipe, e *colly.HTMLElement)[]model.Recipe{
			title := strings.TrimSpace(e.ChildText(u.Title))
			//fmt.Printf("Exclude: %v", isContains(u.Exclude, title))
			if !isContains(u.Exclude, title) {
				if len(u.Include) == 0 || isContains(u.Include, title){
				img := getImg(e, u.Img)
				recipes = append(recipes, 
					model.Recipe{
						Img: []string{img}, 
						Title: title, 
						Url: e.ChildAttr(u.Recipe.Url,"href"),
					})
				}
			}
			return recipes
}

func (u UrlHandler) HandleRecipe(cx echo.Context) error {
	urls := InitUrls()
	url := cx.QueryParam("url")
	// var allowed []string
	allowed := []string{"thewoksoflife.com", "www.inspiredtaste.net"}
	var matched model.List
	for _, u := range urls {
		if strings.Contains(u.Url, getDomain(url))  {
			// allowed = append(allowed,u.Url)
			matched = u
			break
		}
	}

	c := colly.NewCollector(
		colly.AllowedDomains(allowed...),
		colly.CacheDir("./detail_cache"),
	)
	
	var recipe model.Recipe
	var empty model.List
	if empty.Title != matched.Title {

		c.OnHTML(matched.Recipe.Selector, func(e *colly.HTMLElement){
			img := getImg(e, matched.Recipe.Img)
			recipe = model.Recipe{
					Img: []string{img}, 
					Title: strings.TrimSpace(e.ChildText(matched.Recipe.Title)), 
					Url: e.ChildAttr(matched.Recipe.Url,"href"),
					Ingredients: e.ChildTexts(matched.Recipe.Ingredients[0]),
				}
		})
		c.Visit(url)
	}
		
	return render(cx, component.Detail(recipe))	
}
func isContains(ex []string, title string )bool{
	if len(ex) > 0{
		for _, e := range ex{
			lower := strings.ToLower(title)
			if strings.Contains(lower, strings.ToLower(e)) {
				return true
			}	
		}
	}	
	return false
}
func getImg(e *colly.HTMLElement, imgs []string) string{
	var img string
	for _, i := range imgs {
		img = e.ChildAttr(i, "src")
		if strings.HasPrefix(img, "/"){
			//Relative Path
			img = e.Request.AbsoluteURL(img)
		}
		if !strings.HasPrefix(img, "http"){
			img = e.ChildAttr(i, "data-src")
		}
		if !strings.HasPrefix(img, "http"){
			img = e.ChildAttr(i, "data-lazy-src")
		}
		if strings.HasPrefix(img, "http"){
			return img
		}
	}
	return img
}

func getDomain(url string) string{
	domain := strings.TrimPrefix(url, "https://")
	i := strings.Index(domain, "/")
	domain = domain[:i]
	return domain
}