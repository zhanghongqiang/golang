package main

import (
	"log"
	"os"

	"task4/config"
	"task4/database"
	"task4/routes"
	"task4/utils"

	"github.com/gin-gonic/gin"
)

func main() {

	if err := config.LoadConfig("./configs/config.yaml"); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	if err := utils.InitLogger(); err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}

	if err := database.ConnectDB(); err != nil {
		utils.Logger.Fatalf("Failed to connect to database: %v", err)
	}

	if err := database.AutoMigrate(); err != nil {
		utils.Logger.Fatalf("Failed to migrate database: %v", err)
	}

	if config.AppConfig.Server.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	router := routes.SetupRouter(database.DB)

	port := config.AppConfig.Server.Port
	if port == "" {
		port = "8080"
	}

	utils.Logger.Infof("Starting server on port %s", port)
	if err := router.Run(":" + port); err != nil {
		utils.Logger.Fatalf("Failed to start server: %v", err)
		os.Exit(1)
	}
}
