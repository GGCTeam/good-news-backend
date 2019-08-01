package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// NewsController <controller>
// is used for describing controller actions for news.
type NewsController struct{}

// Get <function>
// is used to handle get action of news controller which will return <count> number of news.
// url: /v1/news?count=80 , by default <count> = 50
func (nc NewsController) Get(c *gin.Context) {
	count := c.DefaultQuery("count", "50")

	c.JSON(200, gin.H{
		"news": fmt.Sprintf("with count : %s", count),
	})
}
