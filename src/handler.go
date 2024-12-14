package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type defaultResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			c.JSON(http.StatusInternalServerError, c.Errors)
		}
	}
}

func docsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Param("any") == "/" {
			c.Redirect(http.StatusMovedPermanently, "index.html")
			return
		}
		ginSwagger.WrapHandler(swaggerFiles.Handler)(c)
	}
}

// @Tags Health
// @Summary Health check
// @Description Check if the service is running
// @Produce json
// @Success 200 {object} defaultResponse
// @Router /health [get]
func healthCheckHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, defaultResponse{
			Status:  http.StatusOK,
			Message: "Service is running",
		})
	}
}
