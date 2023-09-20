package controller

import (
	"fmt"
	"net/http"
	"qkeruen/dto"
	"qkeruen/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type offerUserController struct {
	OfferUserService service.OfferUserService
}

func NewOfferUserController(offer service.OfferUserService) offerUserController {
	return offerUserController{OfferUserService: offer}
}

func (c *offerUserController) GetByID(ctx *gin.Context) {
	driverId, _ := strconv.ParseInt(ctx.Param("driverId"), 10, 64)

	res, err := c.OfferUserService.GetByID(driverId)

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

func (c *offerUserController) CreateOffer(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var create dto.OfferRequest
	if err := ctx.ShouldBindJSON(&create); err != nil {
		ctx.JSON(
			http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("bad request: %v\n", err),
			},
		)
		return
	}

	if err := c.OfferUserService.Create(id, create); err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusConflict, gin.H{"message": "error in offer create service."})
		return
	}
	ctx.JSON(200, "Saved.")
}

func (c *offerUserController) GetMyOffer(ctx *gin.Context) {
	userId, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)

	data, err := c.OfferUserService.MyOffer(userId)

	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"message": "error in my offer service."})
		return
	}

	ctx.JSON(201, data)
}

func (c *offerUserController) SearchOffers(ctx *gin.Context) {
	var offer dto.OfferRequest

	if err := ctx.ShouldBindJSON(&offer); err != nil {
		ctx.JSON(
			http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("bad request: %v\n", err),
			},
		)
		return
	}
	res, err := c.OfferUserService.Search(offer.From, offer.To, offer.Type)

	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"message": "error in search offer service."})
		return
	}

	ctx.JSON(200, res)
}

func (c *offerUserController) AllOffer(ctx *gin.Context) {
	allOffer, err := c.OfferUserService.FindAllOffers()

	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"message": "error in Get all offer service."})
		return
	}

	ctx.JSON(200, allOffer)
}

func (c *offerUserController) DeleteOffer(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err := c.OfferUserService.DeleteOffer(id); err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"message": "error in delete offer service."})
		return
	}

	ctx.JSON(200, "Deleted.")
}
