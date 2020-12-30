package http

import (
	"github.com/gin-gonic/gin"
	"github.com/mrcampbell/gacha-game/battle-service/internal/http/handlers"
)

func RegisterHandlers(r *gin.Engine) {
	registerSwaggerHandlers(r)
	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", handlers.Ping)
		v1.POST("/my/team", handlers.MyTeam)
	}
}
