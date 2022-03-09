package main

import "github.com/gin-gonic/gin"

// User demo
type User struct {
	UserName  string
	FirstName string
	LastName  string
}

func main() {
	router := gin.Default()
	router.GET("/randomCocktail", RandomCocktail)
	router.GET("/randomDinner", RandomDinner)
	router.GET("/getCocktails", GetCocktails)
	router.GET("/getDinner", GetDinner)
	router.GET("/randomDessert", RandomDessert)
	router.GET("/getDesserts", GetDessert)
	router.Run()
	// webscraper()
}
