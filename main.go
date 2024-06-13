package main

import (
	"log"

	"github.com/Project_Restaurant/api-gateway/api"
)

func main() {
	engine := api.NewGin()
	err := engine.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
