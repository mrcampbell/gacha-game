package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mrcampbell/gacha-game/monolith/internal/http"
	"github.com/mrcampbell/gacha-game/monolith/internal/http/handlers"
	"github.com/mrcampbell/gacha-game/monolith/internal/inmemory"
)

func main() {
	// none of these are thread safe.
	moveService := inmemory.NewMoveService()
	unitService, err := inmemory.NewUnitService()
	if err != nil {
		log.Fatal(err)
	}

	unitMoveService := inmemory.NewUnitMoveService(moveService)
	fighterService := inmemory.NewFighterService(unitService, unitMoveService)

	handlers.InitializeUnitHandlers(unitService)
	handlers.InitializeFighterHandler(fighterService)

	r := gin.Default()
	http.RegisterHandlers(r)

	if err := r.Run(); err != nil { // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
		println(err)
	}
}
