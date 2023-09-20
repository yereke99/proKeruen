package controller

import (
	"net/http"
	"qkeruen/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type historyController struct {
	HistoryService service.HistoryService
	JWTService     service.JWTService
}

func NewHistoryController(hService service.HistoryService, jwtService service.JWTService) historyController {
	return historyController{
		HistoryService: hService,
		JWTService:     jwtService,
	}
}

func (c *historyController) GetUserH(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)

	hData, err := c.HistoryService.GetDriverHistory(id)

	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"message": "error in user history service."})
		return
	}

	ctx.JSON(200, hData)
}

func (c *historyController) GetDriverH(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	
	hData, err := c.HistoryService.GetDriverHistory(id)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"message": "error in driver history service."})
		return
	}
	ctx.JSON(200, hData)
}
