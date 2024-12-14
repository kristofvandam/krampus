package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func EmptyBodyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.ContentLength == 0 {
			c.Error(fmt.Errorf("empty request body"))
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
	}
}

func BodyMiddleware[T any](t T) gin.HandlerFunc {
	return func(c *gin.Context) {
		data := t
		if err := c.ShouldBindJSON(&data); err != nil {
			c.Error(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		c.Set("data", data)
		c.Next()
	}
}
