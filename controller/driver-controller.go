package controller

import (
	"fmt"
	"net/http"
	"qkeruen/models"
	"qkeruen/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type driverController struct {
	DriverService service.DriverService
	JWTService    service.JWTService
}

func NewDriverController(driverService service.DriverService, jwtService service.JWTService) driverController {
	return driverController{
		DriverService: driverService,
		JWTService:    jwtService,
	}
}

func (d *driverController) Register(ctx *gin.Context) {
	var driver models.DriverRegister

	if err := ctx.ShouldBindJSON(&driver); err != nil {
		ctx.JSON(
			http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("bad request: %v\n", err),
			},
		)
		return
	}

	driver.Token = ctx.GetHeader("Authorization")
	data, err := d.DriverService.CreateDriver(driver)
	//fmt.Println(err.Error())
	if err != nil {
		ctx.JSON(http.StatusConflict, "error in driver create service.")
		return
	}

	ctx.JSON(201, data)
}

func (d *driverController) GetProfile(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")

	if token == "" {
		ctx.JSON(
			http.StatusBadRequest, gin.H{
				"error": "Empty token",
			},
		)
		// exit process
		return
	}
	data, err := d.DriverService.GetProfile(token)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest, gin.H{
				"error": "error in driver check token service.",
			},
		)
		// exit process
		return
	}
	ctx.JSON(200, data)
}

func (c *driverController) Update(ctx *gin.Context) {
	var model models.DriverModel
	token := ctx.GetHeader("Authorization")
	if err := ctx.ShouldBindJSON(&model); err != nil {
		ctx.JSON(
			http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("bad request: %v\n", err),
			},
		)
		// exit process
		return
	}

	model.Token = token
	res, err := c.DriverService.UpdateService(model)

	if err != nil {
		ctx.JSON(http.StatusConflict, "error in update service.")
		return
	}

	ctx.JSON(200, res)
}

func (c *driverController) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}

	if err := c.DriverService.Delete(id); err != nil {
		ctx.JSON(http.StatusConflict, "error in delete service.")
		return
	}

	ctx.JSON(200, "deleted")
}
