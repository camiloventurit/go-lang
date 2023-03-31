package middlewares

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"myApi/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Movie struct {
	Title    string `json:"title"`
	Genre    string `json:"genre"`
	Released string `json:"released"`
	Raiting  string `json:"raiting"`
	ImdbID   string `json:"imdbID"`
}

func SearchMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		title := ctx.Query("title")
		queryType := ctx.Query("queryType")

		//Petition
		resp, err := http.Get("https://www.omdbapi.com/?" + queryType + "=" + title + "&apikey=936c6523")

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}

		defer resp.Body.Close()

		// parse the response body as a string first
		respBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		omdbResponseStr := string(respBytes)

		// parse the omdbResponseStr value as JSON
		var omdbResponseData map[string]interface{}
		if err := json.Unmarshal(respBytes, &omdbResponseData); err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		// set the omdbResponseData value in the custom header
		ctx.Set("X-Omdb-Response", omdbResponseStr)

		var myMovie Movie

		err = json.Unmarshal(respBytes, &myMovie)
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
			return
		}

		// search for the title in the database
		var movie []models.Movie
		if err := db.Where("title = ?", myMovie.Title).First(&movie).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				insertMovie := models.Movie{
					Title:        myMovie.Title,
					ReleasedYear: myMovie.Released,
					Raiting:      myMovie.Raiting,
					IDMovie:      myMovie.ImdbID,
					Genre:        myMovie.Genre,
				}

				if err := db.Create(&insertMovie).Error; err != nil {
					ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create movie record"})
					return
				}

				// Set the movie's information in the custom header
				ctx.Writer.Header().Set("X-Inserted-Movie", fmt.Sprintf("Title: %s, Released Year: %s, Rating: %s, IDMovie: %s, Genre: %s", insertMovie.Title, insertMovie.ReleasedYear, insertMovie.Raiting, insertMovie.IDMovie, insertMovie.Genre))
			} else {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				ctx.Abort()
				return
			}
		} else {
			movieJSON, err := json.Marshal(movie)
			if err != nil {
				fmt.Println("Error", err)
			}

			// set the omdbResponseData value in the custom header
			ctx.Set("X-Inserted-value", string(movieJSON))
		}

		ctx.Next()
	}
}
