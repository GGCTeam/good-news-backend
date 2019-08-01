package server

import (
	"github.com/kanzitelli/good-news-backend/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// NewRouter <function>
// is used to create a GIN engine instance where all controller and routes will be placed
func NewRouter() *gin.Engine {
	// var envVars = utils.GetEnvVars()

	// if envVars.DebugMode {
	// 	gin.SetMode(gin.DebugMode)
	// } else {
	// 	gin.SetMode(gin.ReleaseMode)
	// }

	gin.SetMode(gin.DebugMode)

	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	router.Use(cors.Default())

	v1 := router.Group("v1")
	{
		news := v1.Group("news")
		{
			nc := controllers.NewsController{}

			news.GET("/", nc.Get)
		}
	}

	return router
}
