package main

import (
	"log"
	"qkeruen/config"
	"qkeruen/controller"
	"qkeruen/middleware"
	"qkeruen/models"
	"qkeruen/repository"
	"qkeruen/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	//"time"
)

var (
	AppSettings models.Settings
)

var dbPool, err = config.NewDBPool(config.DataBaseConfig{
	Username: "postgres",
	Password: "123456",
	Hostname: "localhost",
	Port:     "5432",
	DBName:   "postgres",
})

var (
	authDB        = repository.NewDatabase(dbPool)
	driverDB      = repository.NewDriverRepository(dbPool)
	userDB        = repository.NewUserRepository(dbPool)
	offerDriverDB = repository.NewOfferDriverRepository(dbPool)
	offerUserDB   = repository.NewOfferUserRepository(dbPool)
	orderDB       = repository.NewOrderRepository(dbPool)
	historyDB     = repository.NewHistoryRepository(dbPool)
	processDb     = repository.NewProcessRepository(dbPool)
	searchDb      = repository.NewSearchRepository(dbPool)
	securityDB    = repository.NewSecurityService(dbPool)
	jwtService    = service.NewJWTService()

	authService        = service.NewAuthService(authDB)
	driverService      = service.NewDriverService(driverDB)
	userService        = service.NewUserService(userDB)
	offerDriverService = service.NewOfferDriverService(offerDriverDB)
	offerUserService   = service.NewOfferuserService(offerUserDB)
	historyService     = service.NewHistoryService(historyDB)
	orderService       = service.NewOrderService(orderDB)
	processService     = service.NewProcessService(processDb)
	searchService      = service.NewSearchService(searchDb)
	securityService    = service.NewSecurityService(securityDB)

	authController        = controller.NewAuthController(authService, jwtService)
	driverController      = controller.NewDriverController(driverService, jwtService)
	userController        = controller.NewUserController(userService, jwtService)
	offerDriverController = controller.NewOfferDriverController(offerDriverService)
	offerUserController   = controller.NewOfferUserController(offerUserService)
	historyController     = controller.NewHistoryController(historyService, jwtService)
	orderController       = controller.NewOrderController(orderService)
	processController     = controller.NewProcessController(processService)
	searchController      = controller.NewSearchController(searchService)
	securityController    = controller.NewSecurityController(securityService, jwtService)
)

func main() {
	defer dbPool.Close()
	e := config.InitTabeles(dbPool)

	if e != nil {
		log.Println(e)
	} else {
		log.Println("Success init.")
	}

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.Use(gin.Recovery())

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Length", "Authorization", "X-CSRF-Token", "Content-Type", "Accept", "X-Requested-With", "Bearer", "Authority"},
		ExposeHeaders:    []string{"Content-Length", "Authorization", "Content-Type", "application/json", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Accept", "Origin", "Cache-Control", "X-Requested-With"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://api.qkeruen.kz"
		},
	}))

	r.GET("/get", func(ctx *gin.Context) {
		ctx.JSON(200, "Hello")
	})

	r.POST("/authorization/sign", authController.Register)
	r.POST("/authorization/check", authController.ValidatorSMS)
	r.POST("/authorization/resend", authController.ResendCode)
	r.POST("/authorization/check-token", authController.CheckToken)

	r.POST("/user", middleware.AuthorizeJWTUser(jwtService), userController.Register)
	r.GET("/user", middleware.AuthorizeJWTUser(jwtService), userController.GetProfile)
	r.PUT("/user", middleware.AuthorizeJWTUser(jwtService), userController.Update)
	r.DELETE("/user/:id", middleware.AuthorizeJWTUser(jwtService), userController.Delete)

	// Secuirity
	r.POST("/user/security", middleware.AuthorizeJWTUser(jwtService), securityController.Add)
	r.GET("/user/security/:id", middleware.AuthorizeJWTUser(jwtService), securityController.GetMyHistory)
	r.PUT("/user/security/finish/:id", middleware.AuthorizeJWTUser(jwtService), securityController.Finish)

	r.GET("/order/user/for-driver/:id", middleware.AuthorizeJWTDriver(jwtService), orderController.GetOrders)
	//orderDriverRouter.GET("/")
	//orderDriverRouter.DELETE("/:orderId")

	r.POST("/order/user", middleware.AuthorizeJWTUser(jwtService), orderController.CreateOrder)
	r.POST("/order/driver/search/:id", middleware.AuthorizeJWTDriver(jwtService), orderController.GetOrders)
	r.GET("/order/user/:id", middleware.AuthorizeJWTUser(jwtService), orderController.GetMyOrders)
	r.DELETE("/order/user/:id", middleware.AuthorizeJWTUser(jwtService), orderController.DeleteOrder)

	r.POST("/driver", middleware.AuthorizeJWTDriver(jwtService), driverController.Register)
	r.GET("/driver", middleware.AuthorizeJWTDriver(jwtService), driverController.GetProfile)
	r.PUT("/driver", middleware.AuthorizeJWTDriver(jwtService), driverController.Update)
	r.DELETE("/driver/:id", middleware.AuthorizeJWTDriver(jwtService), driverController.Delete)

	r.POST("/offer/driver/:userId", middleware.AuthorizeJWTDriver(jwtService), offerDriverController.GetByID)
	r.POST("/offer/driver/create/:id", middleware.AuthorizeJWTDriver(jwtService), offerDriverController.CreateOffer)
	r.GET("/offer/driver/my/:id", middleware.AuthorizeJWTDriver(jwtService), offerDriverController.GetMyOffer)
	r.GET("/offer/driver/all", middleware.AuthorizeJWTDriver(jwtService), offerDriverController.AllOffer)
	r.POST("/offer/driver/search", middleware.AuthorizeJWTDriver(jwtService), offerDriverController.SearchOffers)
	r.DELETE("/offer/driver/:id", middleware.AuthorizeJWTDriver(jwtService), offerDriverController.DeleteOffer)

	r.POST("/offer/user/getByID/:driverId", middleware.AuthorizeJWTUser(jwtService), offerUserController.GetByID)
	r.POST("/offer/user/:id", middleware.AuthorizeJWTUser(jwtService), offerUserController.CreateOffer)
	r.GET("/offer/user/my/:id", middleware.AuthorizeJWTUser(jwtService), offerUserController.GetMyOffer)
	r.GET("/offer/user/all", middleware.AuthorizeJWTUser(jwtService), offerUserController.AllOffer)
	r.POST("/offer/user/search", middleware.AuthorizeJWTUser(jwtService), offerUserController.SearchOffers)
	r.DELETE("/offer/user/:id", middleware.AuthorizeJWTUser(jwtService), offerUserController.DeleteOffer)

	r.GET("/history/driver/:id", middleware.AuthorizeJWTDriver(jwtService), historyController.GetDriverH)
	r.GET("/history/user/:id", middleware.AuthorizeJWTUser(jwtService), historyController.GetUserH)

	processRouter := r.Group("/process")
	{
		processRouter.POST("/start/:driverId/:orderId", processController.AcceptOrder)
		processRouter.POST("/cancel/:orderId", processController.CancellOrder)
		processRouter.POST("/finish/:driverId/:orderId", processController.FinishOrder)
		processRouter.POST("/driver/:driverId", processController.GetOrdersInProcessDriver)
		processRouter.POST("/user/userId", processController.GetOrdersInProcessUser)
	}

	search := r.Group("/search")
	{
		search.POST("/check/:place", searchController.Check)
		search.POST("/checkGeo/:place", searchController.CheckGeo)
	}

	r.Run(":3000")
}
