package api

import (
	"api-getway/api/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"google.golang.org/grpc"
	_ "api-getway/docs"
)

func NewGin(payme, reser, use *grpc.ClientConn) *gin.Engine {
	handler := handlers.NewHandler(payme, reser, use)

	router := gin.Default()
	router.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Adjust for your specific origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	restaurant := router.Group("api/v1")
	{
		restaurant.POST("/restaurants", handler.CreateRestaurantHandler)
		restaurant.GET("/restaurants", handler.GetRestaurantAllHandler)
        restaurant.GET("/restaurants/:id", handler.GetRestaurantByIdHandler)
        restaurant.PUT("/restaurants/:id", handler.UpdateRestaurantHandler)
        restaurant.DELETE("/restaurants/:id", handler.DeleteRestaurantHandler)
	}

	reservation := router.Group("api/v1")
	{
        reservation.POST("/reservations", handler.CreateReservationHandler)
		reservation.GET("/reservations", handler.GetReservationAllHandler)
		reservation.GET("/reservations/:id", handler.GetReservationByIdHandler)
        reservation.PUT("/reservations/:id", handler.UpdateReservationHandler)
        reservation.DELETE("/reservations/:id", handler.DeleteReservationHandler)
		reservation.POST("/reservations/check", handler.ReservationCheckHandler)
		reservation.POST("/reservations/:id/order", handler.ReservationOrderIdHandler)
		reservation.POST("/reservations/:id/payment", handler.ReservationPaymentHandler)
	}


	menu := router.Group("api/v1")
	{
		menu.POST("/menu", handler.CreateMenuHandler)
		menu.GET("/menu", handler.GetAllMenuHandler)
		menu.GET("/menu/:id", handler.GetMenuByIdHandler)
		menu.PUT("/menu/:id", handler.UpdateMenuHandler)
		menu.DELETE("/menu/:id", handler.DeleteMenuHandler)
	}


	payment := router.Group("api/v1")
	{
        payment.POST("/payments", handler.CreatePaymentHandler)
        payment.GET("/payments/:id", handler.GetByIdPaymentHandler)
        payment.PUT("/payments/:id", handler.UpdatePaymentHandler)
	}

	return router
}
