package api

// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTgzNjE1NTMsImlhdCI6MTcxODI3NTE1MywiaWQiOiJoZWxsbyIsInVzZXJuYW1lIjoiM2Q0NzFkNWYtZWNlMS00ZDNlLWJjZmUtYWUxOWI3MzE2MzQzIn0.kMTaLtfMN6lNOnshX2FUeD_o9D2mIuDM9pB9psZnZLI

// Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTgzNjYwNjAsImlhdCI6MTcxODI3OTY2MCwiaWQiOiI4MzdmNjlhYy1mMWU0LTRhYjAtYTEwOC03Njc4YWFmYjA4ZGEiLCJ1c2VybmFtZSI6ImFkbWluIn0.-FGuoSliQuWBsOcgEDjKvCWw3UPveMSHIs6Wkta9zeY'
import (
	_ "github.com/Project_Restaurant/api-gateway/api/docs"
	"github.com/Project_Restaurant/api-gateway/api/handlers"
	"github.com/Project_Restaurant/api-gateway/api/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"google.golang.org/grpc"
)

// @title Restaurant reservation system API
// @version 1.0
// @description API for managing Restaurant reservation syste resources
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewGin(payme, reser *grpc.ClientConn) *gin.Engine {
	handler := handlers.NewHandler(payme, reser)

	router := gin.Default()
	router.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Use(middlewares.Auth)
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Adjust for your specific origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	restaurant := router.Group("api/v1/restaurants")
	{
		restaurant.POST("/", middlewares.Role, handler.CreateRestaurantHandler)
		restaurant.GET("/:id", handler.GetRestaurantHandler)
		restaurant.PUT("/:id", middlewares.Role, handler.UpdateRestaurantHandler)
		restaurant.DELETE("/:id", middlewares.Role, handler.DeleteRestaurantHandler)
		restaurant.GET("/", handler.ListRestaurantsHandler)
	}

	reservation := router.Group("api/v1")
	{
		reservation.POST("/reservations", handler.CreateReservationHandler)
		reservation.GET("/reservations", handler.ListReservationsHandler)
		reservation.GET("/reservations/:id", handler.GetReservationHandler)
		reservation.DELETE("/reservations/:id", middlewares.Role, handler.DeleteReservationHandler)
		reservation.PUT("/reservations/:id", handler.UpdateReservationHandler)

		reservation.GET("/reservations/check", handler.CheckAvailabilityHandler)
		reservation.GET("/reservations/:id/foodlist", handler.GetFoodListHandler)
		reservation.POST("/reservations/:id/order", handler.OrderFoodHandler)
		reservation.POST("/reservations/:id/bill", handler.OrderBillHandler)
	}

	menu := router.Group("api/v1/menus")
	{
		menu.POST("/", handler.CreateMenuHandler)
		menu.GET("/:id", handler.GetMenuHandler)
		menu.GET("/", handler.GetAllMenusHandler)

		menu.PUT("/:id", middlewares.Role, handler.UpdateMenuHandler)
		menu.DELETE("/:id", middlewares.Role, handler.DeleteMenuHandler)
	}

	payment := router.Group("api/v1/payments")
	{
		payment.POST("/", handler.CreatePaymentHandler)
		payment.GET("/:id", handler.GetPaymentHandler)
		payment.PUT("/:id", handler.UpdatePaymentHandler)
		payment.DELETE("/:id", handler.DeletePaymentHandler)
		payment.POST("/reservation", handler.PayForReservationHandler)
	}

	return router
}
