package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getValueFromContext(ctx *gin.Context, keys ...string) (interface{}, bool) {
	for _, key := range keys {
		value, exists := ctx.Get(key)
		if exists {
			return value, true
		}
	}
	return nil, false
}

func Searcher() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// get the OMDb API response from the custom header
		headerData, exists := getValueFromContext(ctx, "X-Inserted-value", "X-Inserted-movie", "X-Omdb-Response")
		if !exists {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "X-Response header not found",
			})
			return
		}

		headerStr, ok := headerData.(string)
		if !ok {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "X-Response header is not a string",
			})
			return
		}

		var jsonData interface{}
		if err := json.Unmarshal([]byte(headerStr), &jsonData); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to decode JSON data from header",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"response": jsonData,
		})

	}
}
