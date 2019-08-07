package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kanzitelli/good-news-backend/db"
)

// NewsController <controller>
// is used for describing controller actions for news.
type NewsController struct{}

// Get <function>
// is used to handle get action of news controller which will return <count> number of news.
// url: /v1/news?count=80 , by default <count> = 50
func (nc NewsController) Get(c *gin.Context) {
	countStr := c.DefaultQuery("count", "50")
	count, err := strconv.Atoi(countStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Please, enter corrent number of count parameter.",
		})
		return
	}

	dbClient := db.GetClient()
	news, err := dbClient.NewsGet(int64(count))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":       "Problem with fetching news from DB",
			"description": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, news)
}
