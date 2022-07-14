package main

import (
	"os"

	"server/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())

	router.Use(cors.Default())

	// these are the endpoints
	//C
	router.POST("/transport", routes.AddTransport)
	//R
	router.GET("/transports", routes.GetTransports)
	router.GET("/transport/:id", routes.GetTransportById)
	//U
	router.PUT("/transport/:id", routes.UpdateTransport)
	//D
	router.DELETE("/transport/:id", routes.DeleteTransport)

	//this runs the server and allows it to listen to requests.
	router.Run(":" + port)
}
