package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrcampbell/gacha-game/monolith/internal/app"
)

type fighterHandler struct {
	fighterService app.FighterService
}

var FighterHandler fighterHandler

func InitializeFighterHandlers(fs app.FighterService) {
	FighterHandler = fighterHandler{fighterService: fs}
}

// UnitByID godoc
// @Summary Gets Unit by ID (Public/Static).
// @Description responds with single unit
// @Accept  json
// @Produce  json
// @Param id path string false "get unit by id"
// @Success 200 {object} app.Unit
// // @Header 200 {string} Token "qwerty"
// // @Failure 400,404 {object} httputil.HTTPError
// // @Failure 500 {object} httputil.HTTPError
// // @Failure default {object} httputil.DefaultError
// @Router /static/units/{id} [get]
func (h fighterHandler) UnitByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		_ = c.Error(errors.New("missing required parameter: id"))
		return
	}

	unit, err := h.fighterService.UnitByID(c.Request.Context(), id)
	if err != nil {
		_ = c.Error(fmt.Errorf("no unit with id: %s", id))
		return
	}

	c.JSON(http.StatusOK, unit)
}

// AllUnits godoc
// @Summary Gets Unit by ID (Public/Static).
// @Description responds with all available single unit
// @Accept  json
// @Produce  json
// @Success 200 {object} []app.Unit
// // @Header 200 {string} Token "qwerty"
// // @Failure 400,404 {object} httputil.HTTPError
// // @Failure 500 {object} httputil.HTTPError
// // @Failure default {object} httputil.DefaultError
// @Router /static/units [get]
func (h fighterHandler) AllUnits(c *gin.Context) {
	units, err := h.fighterService.ListAllUnits(c.Request.Context())
	if err != nil {
		_ = c.Error(fmt.Errorf("error listing units"))
		return
	}

	c.JSON(http.StatusOK, units)
}
