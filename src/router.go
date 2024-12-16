package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.hagfi.sh/krampus/docs"
	"go.hagfi.sh/krampus/draw"
	"go.hagfi.sh/krampus/member"
)

func loadRouter() {

	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"http://localhost:5173", "http://localhost:4173", "http://localhost:8081"},
		AllowMethods:  []string{"OPTIONS", "GET", "POST", "PATCH", "DELETE"},
		AllowHeaders:  []string{"Origin", "Content-Length", "Content-Type"},
		ExposeHeaders: []string{"Content-Length"},
		MaxAge:        12 * time.Hour,
	}))
	r.Use(ErrorHandler())

	v1 := r.Group("/api/v1")
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1.GET("/health", healthCheckHandler())
	v1.GET("/docs/*any", docsHandler())

	drawRoute := v1.Group("/draw")
	drawRoute.POST("", drawCtrl.CreateHandler())

	drawUuidRoute := drawRoute.Group("/:uuid")
	drawUuidRoute.Use(drawCtrl.UuidMiddleware())

	drawUuidRoute.GET("", drawCtrl.ShowHandler())
	drawUuidRoute.PATCH("",
		EmptyBodyMiddleware(),
		BodyMiddleware[draw.Drawing](draw.Drawing{}),
		drawCtrl.UpdateHandler())
	drawUuidRoute.POST("/member",
		EmptyBodyMiddleware(),
		BodyMiddleware[member.Member](member.Member{}),
		drawCtrl.CreateMemberHandler())
	drawUuidRoute.DELETE("/member/:memberUuid", drawCtrl.DeleteMemberHandler())

	memberRoute := v1.Group("/member")
	memberUuidRoute := memberRoute.Group("/:uuid")
	memberUuidRoute.Use(memberCtrl.UuidMiddleware())
	memberUuidRoute.GET("", memberCtrl.ShowHandler())
	memberUuidRoute.DELETE("", memberCtrl.DeleteHandler())

	log.Info("Starting server on :8080")
	r.Run(":8080")
}
