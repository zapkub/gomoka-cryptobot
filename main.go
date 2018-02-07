package main

import (
	"fmt"
	"gomoka-cryptobot/config"
	"gomoka-cryptobot/connector"
	"gomoka-cryptobot/core"
	"log"

	"github.com/gin-gonic/gin"
)

var app *gin.Engine

func addMessageConnectorToRoute(routeName string, handler core.MessengerConnector) {

}

func createMessengerHandler(r *gin.Engine, config config.Config) {
	line := connector.CreateLineConnector(config)
	app.POST("/line", line.RequestHandler)
}

func main() {
	fmt.Print("Hello")

	config := config.Config{}

	app = gin.Default()
	createMessengerHandler(app, config)
	err := app.Run(":8080")
	if err != nil {
		log.Fatal(err.Error())
	}
}
