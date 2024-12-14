package draw

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	"go.hagfi.sh/krampus/member"
)

type Controller struct {
	DB  *bun.DB
	Log *slog.Logger
}

type DefaultResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// @Tags Drawing
// @Summary Create a drawing
// @Description Create a drawing
// @Accept json
// @Produce json
// @Param data body Drawing true "Drawing data"
// @Success 201 {object} Drawing
// @Router /draw [post]
func (ctrl *Controller) CreateHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		drawing := Drawing{}
		drawing.VisitedAt = time.Now()

		_, err := ctrl.DB.NewInsert().Model(&drawing).Exec(c.Request.Context())
		if err != nil {
			ctrl.Log.Error(err.Error())
			c.JSON(http.StatusInternalServerError, DefaultResponse{
				Status:  http.StatusInternalServerError,
				Message: "Failed to create drawing",
			})
			return
		}

		c.JSON(http.StatusCreated, drawing)
	}
}

// @Tags Drawing
// @Summary Show a drawing
// @Description Get a drawing by name
// @Produce json
// @Param uuid path string true "Drawing UUID"
// @Success 200 {object} Drawing
// @Router /draw/{uuid} [get]
func (ctrl *Controller) ShowHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		drawing := c.MustGet("drawing").(Drawing)

		c.JSON(http.StatusOK, drawing)
	}
}

// @Tags Drawing
// @Summary Update a drawing
// @Description Update a drawing by name
// @Accept json
// @Produce json
// @Param uuid path string true "Drawing UUID"
// @Param data body Drawing true "Drawing data"
// @Success 200 {object} Drawing
// @Router /draw/{uuid} [patch]
func (ctrl *Controller) UpdateHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		drawing := c.MustGet("drawing").(Drawing)
		data := c.MustGet("data").(Drawing)

		_, err := ctrl.DB.NewUpdate().
			Model(&data).
			ExcludeColumn("uuid").
			ExcludeColumn("created_at").
			ExcludeColumn("updated_at").
			ExcludeColumn("visited_at").
			ExcludeColumn("deleted_at").
			ExcludeColumn("slug").
			Where("uuid = ?", drawing.UUID).
			OmitZero().
			Exec(c.Request.Context())

		if err != nil {
			ctrl.Log.Error(err.Error())
			c.JSON(http.StatusInternalServerError, DefaultResponse{
				Status:  http.StatusInternalServerError,
				Message: "Failed to update drawing",
			})
			return
		}

		newDrawing := Drawing{}
		ctrl.DB.NewSelect().
			Model(&newDrawing).
			Where("uuid = ?", drawing.UUID).
			Scan(c.Request.Context())

		c.JSON(http.StatusOK, newDrawing)
	}
}

// @Tags Drawing
// @Summary Member create
// @Description Create a member for a drawing
// @Accept json
// @Produce json
// @Param uuid path string true "Drawing UUID"
// @Param data body DrawingMember true "Member data"
// @Success 201 {object} DrawingMember
// @Router /draw/{uuid}/member [post]
func (ctrl *Controller) CreateMemberHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		drawing := c.MustGet("drawing").(Drawing)
		member := c.MustGet("data").(member.Member)

		member.DrawingUUID = drawing.UUID

		_, err := ctrl.DB.NewInsert().Model(&member).Exec(c.Request.Context())
		if err != nil {
			ctrl.Log.Error(err.Error())
			c.JSON(http.StatusInternalServerError, DefaultResponse{
				Status:  http.StatusInternalServerError,
				Message: "Failed to create member",
			})
			return
		}

		c.JSON(http.StatusCreated, member)
	}
}

// @Tags Drawing
// @Summary Member delete
// @Description Delete a member from a drawing
// @Produce json
// @Param uuid path string true "Drawing UUID"
// @Param memberUuid path string true "Member UUID"
// @Success 204
// @Router /draw/{uuid}/member/{memberUuid} [delete]
func (ctrl *Controller) DeleteMemberHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		memberUUID := c.Param("memberUuid")

		_, err := ctrl.DB.NewDelete().
			Model(&member.Member{}).
			Where("uuid = ?", memberUUID).
			Exec(c.Request.Context())
		if err != nil {
			ctrl.Log.Error(err.Error())
			c.JSON(http.StatusInternalServerError, DefaultResponse{
				Status:  http.StatusInternalServerError,
				Message: "Failed to delete member",
			})
			return
		}

		c.Status(http.StatusNoContent)
	}
}
