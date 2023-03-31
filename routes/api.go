package routes

import (
	"myApi/handlers"
	"myApi/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(api *gin.RouterGroup, db *gorm.DB) {
	api.GET("/hello", handlers.HelloHandler)
	api.POST("/new-movie", handlers.CreateMovie(db))
	api.GET("/all-movies", handlers.AllMovies(db))
	api.GET("/search-by-title", middlewares.SearchMiddleware(db), handlers.Searcher())
}
