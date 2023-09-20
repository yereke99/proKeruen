package controller

import (
	"fmt"
	"log"
	"net/http"
	"qkeruen/models"
	"qkeruen/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type securityController struct {
	SecurityService service.SecurityService
	JWTService      service.JWTService
}

func NewSecurityController(security service.SecurityService, jwtService service.JWTService) *securityController {
	controller := &securityController{
		SecurityService: security,
		JWTService:      jwtService,
	}

	return controller
}

func (s *securityController) Add(ctx *gin.Context) {
	var data models.Security

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(
			http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("bad request: %v\n", err.Error()),
			},
		)
		return
	}
	log.Println(data)

	if err := s.SecurityService.Create(data); err != nil {
		ctx.JSON(
			http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("bad request: %v\n", err.Error()),
			},
		)
		return
	}

	ctx.JSON(200, "created.")
}

func (s *securityController) GetMyHistory(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)

	res, err := s.SecurityService.GetMyHistory(id)

	if err != nil {
		ctx.JSON(
			http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("bad request: %v\n", err.Error()),
			},
		)
		return
	}

	ctx.JSON(200, res)
}

func (s *securityController) Finish(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)

	res, err := s.SecurityService.Finish(id)

	if err != nil {
		ctx.JSON(
			http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("bad request: %v\n", err.Error()),
			},
		)
		return
	}

	ctx.JSON(200, res)
}
