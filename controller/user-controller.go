package controller

import (
	"fmt"
	"net/http"
	"qkeruen/models"
	"qkeruen/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userController struct {
	UserService service.UserService
	JWTService  service.JWTService
}

func NewUserController(userService service.UserService, jwtService service.JWTService) userController {
	return userController{
		UserService: userService,
		JWTService:  jwtService,
	}
}

func (c *userController) Register(ctx *gin.Context) {
	var user models.UserRegister
	fmt.Println(user)
	if err := ctx.ShouldBindJSON(&user); err != nil {
		//fmt.Println(err)
		ctx.JSON(
			http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("bad request: %v\n", err),
			},
		)
		return
	}

	user.Token = ctx.GetHeader("Authorization")
	data, err := c.UserService.Create(user)

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusConflict, gin.H{"message": "error in user create service."})
		return
	}

	ctx.JSON(201, data)
}

func (c *userController) GetProfile(ctx *gin.Context) {
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

	dataUser, err := c.UserService.CheckTokenUser(token)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest, gin.H{
				"error": "error in user check token service.",
			},
		)
		// exit process
		return
	}
	ctx.JSON(200, dataUser)
}

func (c *userController) Update(ctx *gin.Context) {
	var model models.UserModel
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
	res, err := c.UserService.Update(model)

	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"message": "error in update service."})
		return
	}

	ctx.JSON(200, res)
}

func (c *userController) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	if err := c.UserService.Delete(id); err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"message": "error in delete service."})
		return
	}

	ctx.JSON(200, "Deleted.")
}
