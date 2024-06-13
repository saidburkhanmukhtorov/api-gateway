package api

import (
	"net/http"

	_ "github.com/Project_Restaurant/api-gateway/api/docs"
	"github.com/Project_Restaurant/api-gateway/api/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Restaurant Reservation System
// @version 1.0
// @description API for managing restaurant reservation system resources
// @host localhost:8080
// @BasePath /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewGin() *gin.Engine {

	// handler support database connection
	// handler := handler.NewHandler()

	// Connecting to gin router
	router := gin.Default()

	router.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Adjust for your specific origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.GET("/hello", hello)
	router.Use(middleware.Auth)
	return router
}

// @Summary Create a new party
// @Description Creates a new party.
// @Tags Parties
// @Accept json
// @Produce json
// @security BearerAuth
// @Success 200 {object} string "hello"
// @Failure 400 {object} string "Invalid request body"
// @Failure 500 {object} string "Internal server error"
// @Router /hello [get]
func hello(ctx *gin.Context) {
	ctx.JSON(http.StatusUnauthorized, gin.H{"welcome": "Hello"})
}
