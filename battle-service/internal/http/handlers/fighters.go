package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrcampbell/gacha-game/battle-service/internal/app"
)

// todo: dev
var units = map[int]app.FighterUnit{
	1: app.FighterUnit{Element: app.FIRE, Name: "Flintly", Type: app.ATTACKER},
	2: app.FighterUnit{Element: app.WATER, Name: "Splishy", Type: app.ATTACKER},
	3: app.FighterUnit{Element: app.GRASS, Name: "Sprouty", Type: app.ATTACKER},
}

type myTeamRequest struct {
	UserID string `json:"user_id"`
}

// My Team godoc
// @Summary List all members of my team
// @Description Is User Specific
// @Accept  json
// @Produce  json
// @Param req body myTeamRequest false "the user id who owns the team"
// @Success 200 {object} pong
// // @Header 200 {string} Token "qwerty"
// // @Failure 400,404 {object} httputil.HTTPError
// // @Failure 500 {object} httputil.HTTPError
// // @Failure default {object} httputil.DefaultError
// @Router /my/team [post]
func MyTeam(c *gin.Context) {
	req := myTeamRequest{}

	if err := c.ShouldBind(&req); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("failed to parse request body: %s", err))
		return
	}

	if req.UserID == "" {
		c.String(http.StatusBadRequest, "missing user id")
		return
	}

	c.JSON(200, []app.Fighter{
		app.Fighter{
			Level: 7,
			Unit:  units[1],
		},
		app.Fighter{
			Level: 14,
			Unit:  units[0],
		},
		app.Fighter{
			Level: 14,
			Unit:  units[2],
		},
	})
}

// ./todo
