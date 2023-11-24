package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func wordsGameHandler(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"test": "ok"})
}

func main() {
	router := gin.Default()
	router.GET("/words/:letters", wordsGameHandler)
	router.Run()
}
