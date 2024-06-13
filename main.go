package main

import (
	// Import your api package
	"fmt"
	"log"

	"github.com/Project_Restaurant/api-gateway/api"
	_ "github.com/Project_Restaurant/api-gateway/api/docs"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Set up gRPC connections
	paymeConn, err := grpc.NewClient("localhost:8083", grpc.WithTransportCredentials(insecure.NewCredentials())) // Update the address
	if err != nil {
		log.Fatalf("Failed to connect to payment service: %v", err)
	}
	defer paymeConn.Close()

	reserConn, err := grpc.NewClient("localhost:8082", grpc.WithTransportCredentials(insecure.NewCredentials())) // Update the address
	if err != nil {
		log.Fatalf("Failed to connect to reservation service: %v", err)
	}
	defer reserConn.Close()

	// Create the Gin engine with the gRPC clients
	router := api.NewGin(paymeConn, reserConn)
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to connect to gin engine: %v", err)
	}
	fmt.Println("API Gateway running on http://localhost:8080")

}
