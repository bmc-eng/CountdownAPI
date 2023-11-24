package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func wordsGameHandler(c *gin.Context) {
	// Split the letters into an array
	strLetters := c.Param("letters")
	letters := strings.Split(strLetters, ";")

	c.JSON(http.StatusOK, gin.H{"test": letters})
}

func main() {
	router := gin.Default()
	router.GET("/words/:letters", wordsGameHandler)
	router.Run()
}
