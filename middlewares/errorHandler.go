package middlewares

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Check if there's an error
		err := c.Errors.Last()

		if err != nil {
			// Log the error
			log.Printf("Error: %v", err.Error)

			// Respond with JSON error
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})

			// Abort the request, no need to continue processing
			c.Abort()
		}
	}
}
