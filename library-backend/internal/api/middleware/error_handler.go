package middleware

import (
	"library-backend/internal/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func ErrorHandler(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Check if there were any errors
		if len(c.Errors) > 0 {
			err := c.Errors.Last()

			logger.WithFields(logrus.Fields{
				"method": c.Request.Method,
				"path":   c.Request.URL.Path,
				"error":  err.Error(),
			}).Error("Request error")

			errorResp := &models.ErrorResponse{
				Success:   false,
				Error:     "An error occurred",
				Code:      "REQUEST_ERROR",
				Details:   err.Error(),
				Timestamp: time.Now().Format(time.RFC3339),
			}

			c.JSON(http.StatusInternalServerError, errorResp)
		}
	}
}
