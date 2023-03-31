package handlers

import (
	"myApi/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AllMovies(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var movies []models.Movie
		result := db.Find(&movies)
		if result.Error != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
		ctx.JSON(http.StatusOK, movies)
	}
}
