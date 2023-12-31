package handler

import (
	"countdownapi/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Handle the GET request from the user in format l;l;l;l...
func GameHandler(c *gin.Context) {
	// Split the letters into an array
	strLetters := c.Param("letters")
	letters := strings.Split(strLetters, ";")

	// Find the words and return a list of potential words
	filteredWords, filteredDefinitions := model.FindWords(letters)

	// return the JSON file to the user
	c.JSON(http.StatusOK, gin.H{"userLetters": letters,
		"dictionary":  filteredWords,
		"definitions": filteredDefinitions})
}

// Handler for AWS health checks when running in ECS
func HealthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"healthcheck": "OK"})
}
