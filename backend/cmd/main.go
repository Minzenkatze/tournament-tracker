package main

import (
	"tournament-tracker/backend/database"
	"tournament-tracker/backend/rest/competitor"

	"fmt"
	"tournament-tracker/backend/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main (){
	mymodel := models.Competitor{Weight: 73.3}
	fmt.Println(mymodel.Weight)
	
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/",  func(c *gin.Context) {
		c.JSON(200, gin.H{
				"data": "Hello from Gin-gonic & mongoDB",
		})
	})
	competitor.AddCompetitorRoutes(router)
	database.ConnectDB()
	router.Run(":8080")
}