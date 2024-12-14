package draw

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) UuidMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := c.Param("uuid")
		if uuid == "" {
			c.JSON(http.StatusBadRequest, DefaultResponse{
				Status:  http.StatusBadRequest,
				Message: "UUID is required",
			})
			c.Abort()
			return
		}

		drawing := Drawing{}
		err := ctrl.DB.NewSelect().
			Model(&drawing).
			Relation("Members").
			Relation("Groups").
			Where("uuid = ?", uuid).
			Scan(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusNotFound, DefaultResponse{
				Status:  http.StatusNotFound,
				Message: "Drawing not found",
			})
			c.Abort()
			return
		}

		ctrl.DB.NewUpdate().
			Model(&drawing).
			Set("visited_at = ?", time.Now()).
			Where("uuid = ?", uuid).
			Exec(c.Request.Context())

		c.Set("drawing", drawing)
		c.Next()
	}
}
