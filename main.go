package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
)

var dictionary map[string]interface{}
var words []string

func findWords(letters []string) (ret []string) {
	// go through each of the letters and see which words contain
	// the letters

	// step 1 - filter all the words with letter[i] in them
	var filteredWords []string
	filteredWords = words

	for _, letter := range letters {
		var newFilteredWords []string
		for _, s := range filteredWords {
			if strings.Contains(s, letter) {
				// exclude the word if its longer than the number of letters
				if len(letters) >= len(s) {
					newFilteredWords = append(newFilteredWords, s)
				}

			}
		}
		filteredWords = newFilteredWords
	}

	// step 2 - find the word with the longest number of letters

	// check that the word has

	// return the word
	return filteredWords
}

func wordsGameHandler(c *gin.Context) {
	// Split the letters into an array
	strLetters := c.Param("letters")
	letters := strings.Split(strLetters, ";")
	filteredWords := findWords(letters)
	c.JSON(http.StatusOK, gin.H{"test": letters, "dictionary": filteredWords})
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

	// sort the words largest to smallest
	sort.Slice(words, func(i, j int) bool {
		l1, l2 := len(words[i]), len(words[j])
		if l1 != l2 {
			return l1 > l2
		}
		return words[i] > words[j]
	})
	fmt.Println("setup complete")
}

func main() {
	router := gin.Default()
	router.GET("/words/:letters", wordsGameHandler)
	router.Run()
}
