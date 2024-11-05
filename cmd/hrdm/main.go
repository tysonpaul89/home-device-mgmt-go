package main

import (
	"hrdm/pkg/config" // <-- Don't change this. The config should be loaded first
	"hrdm/pkg/log"

	"github.com/gin-gonic/gin"
)

// Main function initalizes, logger, config files, web servers, etc.
func main() {
	log.Info("Started execution of the main function")

	// ============ Logger configuration ============
	// To write the gin log to both console and to the file
	multiWriter := log.GetMultiWriter()
	gin.DefaultWriter = multiWriter      // Sets the writer for Gin's logger
	gin.DefaultErrorWriter = multiWriter // If you also want to log recovery (panic) to the file

	// ======================== Getting configurations =========================

	log.Info(config.GetEnv("logLevel"))

	// ============ Gin configuration ============
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {

		log.Debug("Debug")
		log.Info("Info")
		log.Warn("Warn")
		log.Error("Error")

		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
