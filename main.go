package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

var dictionary map[string]interface{}
var words []string

func wordsGameHandler(c *gin.Context) {
	// Split the letters into an array
	strLetters := c.Param("letters")
	letters := strings.Split(strLetters, ";")

	c.JSON(http.StatusOK, gin.H{"test": letters, "dictionary": words[0]})
}

func init() {
	//unpack the json file
	file, _ := os.ReadFile("data/dictionary.json")
	_ = json.Unmarshal([]byte(file), &dictionary)

	// pull all the keys into a single words array list
	words = make([]string, 0, len(dictionary))
	for k := range dictionary {
		words = append(words, k)
	}
	fmt.Println("setup complete")
}

func main() {
	router := gin.Default()
	router.GET("/words/:letters", wordsGameHandler)
	router.Run()
}
