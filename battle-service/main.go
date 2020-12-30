package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mrcampbell/gacha-game/battle-service/internal/http"
)

func main() {
	r := gin.Default()

	http.RegisterHandlers(r)

	if err := r.Run(); err != nil { // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
		println(err)
	}
}
