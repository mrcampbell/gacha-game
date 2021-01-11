package handlers

import (
	"github.com/gin-gonic/gin"
)

type pong struct {
	Message string `json:"message"`
}

// Ping godoc
// @Summary Ensure the API is working/listening
// @Description responds with `pong`
// @Accept  json
// @Produce  json
// // @Param q query string false "name search by q"
// @Success 200 {object} pong
// // @Header 200 {string} Token "qwerty"
// // @Failure 400,404 {object} httputil.HTTPError
// // @Failure 500 {object} httputil.HTTPError
// // @Failure default {object} httputil.DefaultError
// @Router /ping [get]
func Ping(c *gin.Context) {
	c.JSON(200, pong{Message: "pong"})
}
