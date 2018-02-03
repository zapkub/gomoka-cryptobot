package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Print("Hello")

	app := gin.Default()

	err := app.Run(":8080")
	if err != nil {
		log.Fatal(err.Error())
	}
}
