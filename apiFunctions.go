package main

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

type recipe struct {
	Id                 int        `json:"Id"`
	Name               string     `json:"Name"`
	Image              string     `json:"Image"`
	Stars              float64    `json:"Stars"`
	Ratings            int        `json:"Ratings"`
	Times              []string   `json:"Times"`
	Difficulty         string     `json:"Difficulty"`
	Serves             string     `json:"Serves"`
	Description        string     `json:"Description"`
	Ingredients        [][]string `json:"Ingredients"`
	IngredientHeadings []string   `json:"IngredientHeadings"`
	Method             []string   `json:"Method"`
}

// type recipeSummary struct {
// 	Id          int     `json:"Id"`
// 	Name        string  `json:"Name"`
// 	Image       string  `json:"Image"`
// 	Stars       float64 `json:"Stars"`
// 	Ratings     int     `json:"Ratings"`
// 	Description string  `json:"Description"`
// }

func RandomCocktail(ginReturn *gin.Context) {
	ginReturn.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	cocktails := []recipe{}
	file, _ := ioutil.ReadFile("recipes/cocktails.json")

	_ = json.Unmarshal([]byte(file), &cocktails)
	min := 1
	max := 149
	cocktailId := rand.Intn(max-min) + min
	for _, cocktail := range cocktails {
		if cocktail.Id == cocktailId {
			ginReturn.IndentedJSON(http.StatusOK, cocktail)
		}
	}

}

func RandomDinner(ginReturn *gin.Context) {
	ginReturn.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	dinnerRecipes := []recipe{}
	file, _ := ioutil.ReadFile("recipes/dinner.json")

	_ = json.Unmarshal([]byte(file), &dinnerRecipes)
	min := 1
	max := 487
	dinnerId := rand.Intn(max-min) + min
	for _, dinnerRecipe := range dinnerRecipes {
		if dinnerRecipe.Id == dinnerId {
			ginReturn.IndentedJSON(http.StatusOK, dinnerRecipe)
		}
	}

}
