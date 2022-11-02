package main

import (
	"log"
	"net/http"
	route "servers/restAPI/route"

	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	app := gin.New()

	port := os.Getenv("port")

	app.Use(gin.Logger())
	app.Use(gin.Recovery())
	app.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		AllowWildcard: true,
	}))

	app.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"info":   "404 page not found",
		})
	})

	app.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"status":  http.StatusMethodNotAllowed,
			"message": "405 method not allowed",
		})
	})

	app.GET("/error", func(c *gin.Context) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"info":   "Internal Server Error",
		})
	})

	route.V1(app)

	app.Run(":" + port)
}
