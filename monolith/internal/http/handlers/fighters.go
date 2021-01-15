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

func InitializeFighterHandler(fs app.FighterService) {
	FighterHandler = fighterHandler{fighterService: fs}
}

// Show Fighter godoc
// @Summary Gets Fighter by ID (Public/Static).
// @Description responds with single fighter
// @Header 200 {string} Token "qwerty"
// @Accept  json
// @Produce  json
// @Param id path string false "get fighter by id"
// @Success 200 {object} app.Fighter
// // @Header 200 {string} Token "qwerty"
// // @Failure 400,404 {object} httputil.HTTPError
// // @Failure 500 {object} httputil.HTTPError
// // @Failure default {object} httputil.DefaultError
// @Router /fighters/{id} [get]
func (h fighterHandler) FighterByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		_ = c.Error(errors.New("missing required parameter: id"))
		return
	}

	fighter, err := h.fighterService.FighterByID(c.Request.Context(), "mike", id)
	if err != nil {
		_ = c.AbortWithError(http.StatusNotFound, fmt.Errorf("no unit with id: %s", id))
		return
	}

	c.JSON(http.StatusOK, fighter)
}

// // List Units godoc
// // @Summary Gets Unit by ID (Public/Static).
// // @Description responds with all available single unit
// // @Accept  json
// // @Produce  json
// // @Success 200 {object} []app.Unit
// // // @Header 200 {string} Token "qwerty"
// // // @Failure 400,404 {object} httputil.HTTPError
// // // @Failure 500 {object} httputil.HTTPError
// // // @Failure default {object} httputil.DefaultError
// // @Router /units [get]
// func (h fighterHandler) FightersByUserID(c *gin.Context) {
// 	units, err := h.unitService.ListAllUnits(c.Request.Context())
// 	if err != nil {
// 		_ = c.Error(fmt.Errorf("error listing units"))
// 		return
// 	}

// 	c.JSON(http.StatusOK, units)
// }

// // Create Unit godoc
// // @Summary Create Unit (Administrative)
// // @Description responds with the newly created unit
// // @Accept json
// // @Produce json
// // @Param unit body app.Unit true "new unit"
// // @Success 200 {object} app.Unit "newly created unit"
// // // @Header 200 {string} Token "qwerty"
// // // @Failure 400,404 {object} httputil.HTTPError
// // // @Failure 500 {object} httputil.HTTPError
// // // @Failure default {object} httputil.DefaultError
// // @Router /units [post]
// func (h unitHandler) CreateFighter(c *gin.Context) {
// 	var newUnit app.Unit
// 	err := c.ShouldBind(&newUnit)

// 	if err != nil {
// 		_ = c.AbortWithError(http.StatusBadRequest, fmt.Errorf("invalid create unit request: %v", newUnit))
// 		return
// 	}

// 	createdUnit, err := h.unitService.CreateUnit(c.Request.Context(), newUnit)
// 	if err != nil {
// 		_ = c.AbortWithError(http.StatusNotFound, fmt.Errorf("failed to create unit: %s", err))
// 		return
// 	}

// 	c.JSON(http.StatusCreated, createdUnit)
// }
