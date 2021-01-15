package http

import (
	"github.com/gin-gonic/gin"
	"github.com/mrcampbell/gacha-game/monolith/internal/http/handlers"
)

func RegisterHandlers(r *gin.Engine) {
	registerSwaggerHandlers(r)
	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", handlers.Ping)
		v1.GET("/units", handlers.UnitHandler.AllUnits)
		v1.GET("/units/:id", handlers.UnitHandler.UnitByID)
		v1.POST("/units", handlers.UnitHandler.CreateUnit)

		v1.GET("/fighters/:id", handlers.FighterHandler.FighterByID)

		// my
		v1.GET("/my/fighters", handlers.FighterHandler.FightersByUserID)
	}
}
