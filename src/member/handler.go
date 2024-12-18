package member

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type Controller struct {
	DB  *bun.DB
	Log *slog.Logger
}

type DefaultResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// @Tags Member
// @Summary Show a member
// @Description Get a member by name
// @Produce json
// @Param uuid path string true "Drawing UUID"
// @Param memberUuid path string true "Member UUID"
// @Success 200 {object} DrawingMember
// @Router /draw/{uuid}/member/{memberUuid} [get]
func (ctrl *Controller) ShowHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		member := c.MustGet("member").(Member)

		c.JSON(http.StatusOK, member)
	}
}

func (ctrl *Controller) DeleteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		member := c.MustGet("member").(Member)

		ctrl.Log.Info("deleting member",
			"drawUuid", member.DrawingUUID,
			"memberUuid", member.UUID,
			"name", member.Name,
		)

		if _, err := ctrl.DB.NewDelete().
			Model(&member).
			Where("uuid = ?", member.UUID).
			Exec(c.Request.Context()); err != nil {
			ctrl.Log.Error(err.Error())
			c.JSON(http.StatusInternalServerError, DefaultResponse{
				Status:  http.StatusInternalServerError,
				Message: "Error deleting member",
			})
			return
		}

		ctrl.Log.Info("deleted member",
			"drawUuid", member.DrawingUUID,
			"memberUuid", member.UUID,
			"name", member.Name,
		)

		c.JSON(http.StatusOK, DefaultResponse{
			Status:  http.StatusOK,
			Message: "Member deleted",
		})
	}
}
