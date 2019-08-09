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

// GetSources <function>
// is used to handle get action of news controller which will return all news sources.
// url: /v1/news/sources
func (nc NewsController) GetSources(c *gin.Context) {
	dbClient := db.GetClient()
	newsSources, err := dbClient.NewsSourcesGet()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":       "Problem with fetching news sources from DB",
			"description": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, newsSources)
}

// GetTypes <function>
// is used to handle get action of news controller which will return all news types.
// url: /v1/news/sources
func (nc NewsController) GetTypes(c *gin.Context) {
	dbClient := db.GetClient()
	newsTypes, err := dbClient.NewsTypesGet()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":       "Problem with fetching news types from DB",
			"description": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, newsTypes)
}
