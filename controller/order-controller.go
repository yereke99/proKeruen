package controller

import (
	"fmt"
	"net/http"
	"qkeruen/dto"
	"qkeruen/help"
	"qkeruen/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type orderController struct {
	OrderService service.OrderService
}

func NewOrderController(service service.OrderService) orderController {
	return orderController{OrderService: service}
}

func (c *orderController) CreateOrder(ctx *gin.Context) {
	var order dto.OrderRequest

	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(
			http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("bad request: %v\n", err),
			},
		)
		return
	}

	if err := c.OrderService.CreateOrder(order); err != nil {
		ctx.JSON(http.StatusConflict, "error in create order service.")
		return
	}

	ctx.JSON(200, "Order created")
}

func (c *orderController) GetOrders(ctx *gin.Context) {
	var location help.Location
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err := ctx.ShouldBindJSON(&location); err != nil {
		ctx.JSON(
			http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("bad request: %v\n", err),
			},
		)
		return
	}
	res, err := c.OrderService.GetOrders(id, location)
	if err != nil {
		ctx.JSON(http.StatusConflict, "error in search order for driver service.")
		return
	}

	ctx.JSON(200, res)
}

func (c *orderController) GetMyOrders(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)

	res, err := c.OrderService.GetMyOrders(id)

	if err != nil {
		ctx.JSON(http.StatusConflict, "error in get my order service.")
		return
	}

	ctx.JSON(200, res)

}

func (c *orderController) DeleteOrder(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err := c.OrderService.DeleteOrder(id); err != nil {
		ctx.JSON(http.StatusConflict, "error in delete order service.")
		return
	}

	ctx.JSON(200, "Deleted order.")
}
