package competitor

import (
	"net/http"
	"tournament-tracker/backend/models"

	"tournament-tracker/backend/database"

	"context"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var competitorCollection *mongo.Collection = database.GetCollection(database.DB, "competitors")

func AddCompetitor (c *gin.Context) {
	var newCompetitor models.Competitor
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
	if err := c.BindJSON(&newCompetitor); err != nil {
        return
    }
	result, err := competitorCollection.InsertOne(ctx, newCompetitor)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusCreated, result.InsertedID)
}

func GetCompetitors (c *gin.Context){
	
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        var users []models.Competitor
        defer cancel()

        results, err := competitorCollection.Find(ctx, bson.M{})

        if err != nil {
            c.IndentedJSON(http.StatusInternalServerError, "error")
            return
        }

        //reading from the db in an optimal way
        defer results.Close(ctx)
        for results.Next(ctx) {
            var singleUser models.Competitor
            if err = results.Decode(&singleUser); err != nil {
                c.IndentedJSON(http.StatusInternalServerError, "error")
            }

            users = append(users, singleUser)
        }

        c.IndentedJSON(http.StatusOK, map[string]interface{}{"data": users})
    }



func AddCompetitorRoutes (r *gin.Engine){
	r.POST("/competitors", AddCompetitor)
    r.GET("/competitors", GetCompetitors)
}