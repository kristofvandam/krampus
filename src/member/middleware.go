package member

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

		member := Member{}
		err := ctrl.DB.NewSelect().
			Model(&member).
			Where("uuid = ?", uuid).
			Scan(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusNotFound, DefaultResponse{
				Status:  http.StatusNotFound,
				Message: "Member not found",
			})
			c.Abort()
			return
		}

		ctrl.DB.NewUpdate().
			Model(&member).
			Set("visited_at = ?", time.Now()).
			Where("uuid = ?", uuid).
			Exec(c.Request.Context())

		c.Set("member", member)
		c.Next()
	}
}
