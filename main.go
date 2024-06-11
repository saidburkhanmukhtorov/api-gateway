package main

import (
	"api-getway/api" // Import your api package
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// @title Restaurant reservation system API
// @version 1.0
// @description API for managing Restaurant reservation syste resources
// @host localhost:8080
// @BasePath /api/v1
// @in header
// @name Authorization
func main() {
	// Set up gRPC connections
	paymeConn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials())) // Update the address
	if err != nil {
		log.Fatalf("Failed to connect to payment service: %v", err)
	}
	defer paymeConn.Close()

	reserConn, err := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials())) // Update the address
	if err != nil {
		log.Fatalf("Failed to connect to reservation service: %v", err)
	}
	defer reserConn.Close()

	menuConn, err := grpc.NewClient("localhost:50053", grpc.WithTransportCredentials(insecure.NewCredentials())) // Update the address
	if err != nil {
		log.Fatalf("Failed to connect to menu service: %v", err)
	}
	defer menuConn.Close()

	userConn, err := grpc.NewClient("localhost:50054", grpc.WithTransportCredentials(insecure.NewCredentials())) // Update the address
	if err != nil {
		log.Fatalf("Failed to connect to user service: %v", err)
	}
	defer userConn.Close()

	// Create the Gin engine with the gRPC clients
	router := api.NewGin(paymeConn, reserConn, userConn)
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to connect to gin engine: %v", err)
	}
	fmt.Println("API Gateway running on http://localhost:8080")

}
