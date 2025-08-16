package handler

import (
	"countdownapi/model"
	"net/http"
	"strconv"
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

// Handle the GET request for numbers game in format n,n,n,n,n,n/target
func NumbersHandler(c *gin.Context) {
	// Get numbers and target from URL parameters
	numbersParam := c.Param("numbers")
	targetParam := c.Param("target")
	
	// Parse target
	target, err := strconv.Atoi(targetParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid target number"})
		return
	}
	
	// Parse numbers (comma-separated)
	numberStrings := strings.Split(numbersParam, ",")
	if len(numberStrings) != 6 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Must provide exactly 6 numbers"})
		return
	}
	
	numbers := make([]int, 6)
	for i, numStr := range numberStrings {
		num, err := strconv.Atoi(strings.TrimSpace(numStr))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid number: " + numStr})
			return
		}
		numbers[i] = num
	}
	
	// Solve the numbers game
	result := model.SolveNumbersEnhanced(numbers, target)
	
	// Return the JSON response
	c.JSON(http.StatusOK, result)
}

// Handler for AWS health checks when running in ECS
func HealthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"healthcheck": "OK"})
}
