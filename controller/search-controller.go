package controller

import (
	"net/http"
	"qkeruen/service"

	"github.com/gin-gonic/gin"
)

type searchController struct {
	SearchService service.SearchService
}

func NewSearchController(service service.SearchService) searchController {
	return searchController{
		SearchService: service,
	}
}

func (c *searchController) Check(ctx *gin.Context) {
	place := ctx.Param("place")

	res, err := c.SearchService.Check(place)

	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"message": "error in Search service."})
		return
	}

	ctx.JSON(200, res)
}

func (c *searchController) Create(ctx *gin.Context) {
	place := ctx.Param("place")

	if err := c.SearchService.Create(place); err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"message": "error in Search create service."})
		return
	}

	ctx.JSON(200, "created new places")
}

func (c *searchController) CheckGeo(ctx *gin.Context) {
	place := ctx.Param("place")

	res, err := c.SearchService.CheckGeo(place)

	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"message": "error in Search geo service."})
		return
	}

	ctx.JSON(200, res)
}

func (c *searchController) CreateGeo(ctx *gin.Context) {
	place := ctx.Param("place")

	if err := c.SearchService.CreateGeo(place); err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"message": "error in Search create geo service."})
		return
	}

	ctx.JSON(200, "created new places")
}
