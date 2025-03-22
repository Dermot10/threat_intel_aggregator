package main

import (
	"log"

	"github.com/dermot10/threat_intel_aggregator/api/handlers"
	"github.com/dermot10/threat_intel_aggregator/internal/config"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnvironmentVariables()

	db := config.InitDB()

	router := gin.Default() //adds logging and recovery middleware, router

	handlers.RegisterRoutes(router, db)

	for _, route := range router.Routes() {
		log.Printf("Registered Route: %s %s\n", route.Method, route.Path)
	}

	err := router.Run(":8080") // API on port 8080
	if err != nil {
		log.Fatal(err)
	}
}
