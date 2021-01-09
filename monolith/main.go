package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mrcampbell/gacha-game/monolith/internal/http"
	"github.com/mrcampbell/gacha-game/monolith/internal/http/handlers"
	"github.com/mrcampbell/gacha-game/monolith/internal/inmemory"
)

func main() {
	fighterService := inmemory.NewFighterService()
	handlers.InitializeFighterHandlers(fighterService)

	r := gin.Default()
	http.RegisterHandlers(r)

	if err := r.Run(); err != nil { // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
		println(err)
	}
}
