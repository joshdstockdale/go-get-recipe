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
			Url: "https://thewoksoflife.com/category/recipes/chinese-take-out/",
			Category: model.Asian,
			Selector: ".entry-header",
			Title: ".entry-title-link",
			Img: []string{"img"},
			Paginator: ".pagination ul li a",
			Recipe: model.Recipe{
				Selector: "main",
				Title: ".entry-title",
				Img: []string{".wp-post-image"},
				Ingredients: []string{".wprm-recipe-ingredient"},
				Url: ".entry-title-link",
			},
		},
		{
			Url: "https://www.inspiredtaste.net/category/recipes/main-dishes/",
			Category: model.Burger,
			Selector: ".box-container .post_box",
			Title: ".headline",
			Img: []string{"img"},
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
	categories := []model.Category{model.Asian, model.Burger}
	return render(cx, page.Index(categories))
}

func (u UrlHandler) HandleList(cx echo.Context) error {
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
	for _, u := range matched {

		c.OnHTML(u.Selector, func(e *colly.HTMLElement){
			
			// fmt.Printf("Img %v\n Title %v\n Url %v\n", e.ChildAttr(u.Img, "data-src"), e.ChildText("a[href]"), e.ChildAttr("a", "href"))
			img := getImg(e, u.Img)
			recipes = append(recipes, 
				model.Recipe{
					Img: []string{img}, 
					Title: strings.TrimSpace(e.ChildText(u.Title)), 
					Url: e.ChildAttr(u.Recipe.Url,"href"),

				})
		})
		c.Visit(u.Url)
	}
		
	return render(cx, component.List(recipes))	
}

func (u UrlHandler) HandleDetail(cx echo.Context) error {
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
		colly.CacheDir("./recipe_cache"),
	)
	
	var recipe model.Recipe
	var empty model.List
	if empty.Title != matched.Title {

		c.OnHTML(matched.Recipe.Selector, func(e *colly.HTMLElement){
			
			// fmt.Printf("Img %v\n Title %v\n Url %v\n", e.ChildAttr(u.Img, "data-src"), e.ChildText("a[href]"), e.ChildAttr("a", "href"))
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