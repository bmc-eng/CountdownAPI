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
	router.StaticFile("/", "./index.html")
	router.StaticFile("/styles.css", "./styles.css")
	router.StaticFile("/script.js", "./script.js")
	router.GET("/words/:letters", handler.GameHandler)
	router.GET("/numbers/:numbers/:target", handler.NumbersHandler)
	router.GET("/health", handler.HealthCheckHandler)
	router.Run(":3000")
}
