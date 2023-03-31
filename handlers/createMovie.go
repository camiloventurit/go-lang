package handlers

import (
	"fmt"
	"myApi/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateMovie(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var input models.Movie

		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Create movie record

		movie := models.Movie{
			Title:        input.Title,
			ReleasedYear: input.ReleasedYear,
			Raiting:      input.Raiting,
			IDMovie:      input.IDMovie,
			Genre:        input.Genre,
		}

		fmt.Println(movie)
		if err := db.Create(&movie).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create movie record"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": movie})
	}
}
