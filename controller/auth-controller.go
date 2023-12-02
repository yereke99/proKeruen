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

type authController struct {
	AuthService service.AuthService
	JWTService  service.JWTService
}

func NewAuthController(authservice service.AuthService, jwtService service.JWTService) authController {
	return authController{
		AuthService: authservice,
		JWTService:  jwtService,
	}
}

func (c *authController) Register(ctx *gin.Context) {
	var registerDTO dto.RequestRegisterDTO

	if err := ctx.ShouldBindJSON(&registerDTO); err != nil {
		ctx.JSON(
			http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("bad request: %v\n", err),
			},
		)
		return
	}

	code := help.GenerateRandomId(4)
	if smsServiceErr := c.AuthService.Create(registerDTO.PhoneNumber, code); smsServiceErr != nil {
		ctx.JSON(http.StatusConflict, smsServiceErr.Error())
		return
	}

	ctx.JSON(200, "Sent sms code")
}

func (c *authController) ValidatorSMS(ctx *gin.Context) {
	var checkCode dto.CheckCodeRequest
	var responseDTO dto.ResponseDTO
	if err := ctx.ShouldBindJSON(&checkCode); err != nil {
		ctx.JSON(
			http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("bad request: %v\n", err),
			},
		)
		return
	}

	code, _ := strconv.Atoi(checkCode.Code)
	ok, err := c.AuthService.ValidateSMS(checkCode.PhoneNumber, code)

	if err != nil {
		ctx.JSON(
			http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("bad request: %v\n", err),
			},
		)
		return
	}

	if !ok {
		ctx.JSON(http.StatusConflict, "wrong sms code.")
		return
	}

	_ok, err := c.AuthService.Check(checkCode.PhoneNumber, checkCode.Role)
	if err != nil {
		ctx.JSON(http.StatusConflict, "error in check sms service.")
		return
	}
	if !_ok {
		token_, err := c.AuthService.GiveTokenService(checkCode.PhoneNumber, checkCode.Role)

		if err != nil {
			ctx.JSON(http.StatusConflict, "can not take token.")
			return
		}
		responseDTO.Token = token_
		responseDTO.IsAuthorized = true
		ctx.JSON(http.StatusAccepted, responseDTO)
		return
	}

	token, err := c.JWTService.GenerateToken(checkCode.PhoneNumber, checkCode.Role)
	if err != nil {
		ctx.JSON(http.StatusConflict, "can not generate token")
		return
	}

	responseDTO.Token = token
	responseDTO.IsAuthorized = false
	ctx.JSON(200, responseDTO)
}

func (c *authController) ResendCode(ctx *gin.Context) {
	var registerDTO dto.RequestRegisterDTO

	if err := ctx.ShouldBindJSON(&registerDTO); err != nil {
		ctx.JSON(
			http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("bad request: %v\n", err),
			},
		)
		// exit process
		return
	}
	code := help.GenerateRandomID(4)
	if smsServiceErr := c.AuthService.Create(registerDTO.PhoneNumber, code); smsServiceErr != nil {
		ctx.JSON(http.StatusConflict, "error in sms service.")
		return
	}

	ctx.JSON(200, "Resent sms code")
}

func (c *authController) CheckToken(ctx *gin.Context) {
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

	which, _ := c.JWTService.Definition(token)
	fmt.Println(which)
	switch which {
	case "driver":
		data, err := c.AuthService.CheckTokenDriver(token)
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

	case "user":
		datauser, err := c.AuthService.CheckTokenUser(token)
		if err != nil {
			ctx.JSON(
				http.StatusBadRequest, gin.H{
					"error": "error in user check token service.",
				},
			)
			// exit process
			return
		}
		ctx.JSON(200, datauser)

	}

}
