package main

import (
	"countdownapi/handler"
	"countdownapi/model"
	"fmt"

	"github.com/gin-gonic/gin"
)

// load required models before running
func init() {
	model.LoadDictionary()
	fmt.Println("setup complete")
}

// set up router and run
func main() {
	router := gin.Default()
	router.GET("/words/:letters", handler.GameHandler)
	router.GET("/", handler.HealthCheckHandler)
	router.Run(":3000")
}
