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
		v1.GET("/static/units", handlers.UnitHandler.AllUnits)
		v1.GET("/static/units/:id", handlers.UnitHandler.UnitByID)
	}
}
