package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

func webscraper() {
	c := colly.NewCollector()
	detailCollector := c.Clone()
	recipes := []recipe{}
	c.OnError(func(_ *colly.Response, e error) {
		fmt.Println(e)
	})
	id := 0
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting ", r.URL.String())
	})
	c.OnHTML(".row", func(h *colly.HTMLElement) {
		h.ForEach("div.template-search-universal__card", func(i int, d *colly.HTMLElement) {
			url := "https://www.bbcgoodfood.com" + d.ChildAttr(".standard-card-new__article-title", "href")
			URL := h.Request.AbsoluteURL(url)
			detailCollector.Visit(URL)
		})
	})
	detailCollector.OnHTML(".post", func(f *colly.HTMLElement) {
		steps := []string{}
		ingredientsList := [][]string{}
		headings := []string{}

		ingredients := f.DOM.Find(".recipe__ingredients")
		ingredients.Find("section").Each(func(sectionIndex int, section *goquery.Selection) {
			if sectionIndex == 0 {
				headings = append(headings, "Ingredients")
			} else {
				headings = append(headings, section.Find("h3").Text())
			}
			tempList := []string{}
			section.Find("li").Each(func(itemIndex int, s *goquery.Selection) {
				tempList = append(tempList, s.Text())
			})
			ingredientsList = append(ingredientsList, tempList)
		})

		method := f.DOM.Find(".recipe__method-steps")
		method.Find("li").Each(func(listIndex int, stepDetails *goquery.Selection) {
			steps = append(steps, stepDetails.Find("p").Text())
		})

		ratingString := f.ChildText(".mr-lg > div:nth-child(1) > span:nth-child(2)")

		ratingValue := 0.0

		if ratingString == "" {
			ratingValue = 0
		} else {
			ratingValue, _ = strconv.ParseFloat(strings.Split(ratingString, " ")[4], 32)
		}

		rateCountString := f.ChildText("div.mr-lg:nth-child(1) > div:nth-child(1) > span:nth-child(3)")
		rateCountValue, err := strconv.Atoi(strings.Split(rateCountString, " ")[0])

		if err != nil {
			fmt.Println(err)
		}

		times := []string{}

		timeList := f.DOM.Find(".time-range-list > div:nth-child(2) > ul:nth-child(1)")
		timeList.Find("li").Each(func(i int, s *goquery.Selection) {
			times = append(times, s.Text())
		})
		difficulty := f.DOM.Find(".post-header__skill-level > div:nth-child(2)")
		serves := f.DOM.Find(".post-header__servings > div:nth-child(2)")
		id += 1
		temp := recipe{
			Id:                 id,
			Name:               f.ChildText(".headline"),
			Image:              f.ChildAttr(".image__img", "src"),
			Times:              times,
			Stars:              ratingValue,
			Ratings:            rateCountValue,
			Difficulty:         difficulty.Text(),
			Serves:             serves.Text(),
			Description:        f.ChildText("div.editor-content:nth-child(1) > p:nth-child(1)"),
			Ingredients:        ingredientsList,
			IngredientHeadings: headings,
			Method:             steps,
		}
		recipes = append(recipes, temp)
	})

	for i := 1; i <= 37; i++ {
		var visitUrl string
		if i == 1 {
			visitUrl = "https://www.bbcgoodfood.com/search?q=dessert"
		} else {
			visitUrl = fmt.Sprintf("https://www.bbcgoodfood.com/search/recipes/page/%s/?q=dessert&sort=-relevance", strconv.Itoa(i))
		}
		c.Visit(visitUrl)
	}
	// c.Visit("https://www.bbcgoodfood.com/search/recipes?q=Easy+dinner+recipes")
	file, _ := json.MarshalIndent(recipes, "", " ")
	_ = ioutil.WriteFile("recipes/dessert.json", file, 0644)
}
