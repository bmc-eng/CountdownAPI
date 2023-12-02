package model

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
)

var words []string
var dictionary map[string]interface{}

func removeIndex(s []string, index string) []string {
	//var indexInt int
	if len(s) == 0 {
		return s
	}
	for i := 0; i < len(s); i++ {
		if s[i] == index {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s

}

func checkLetter(letter string, lettersFromUser []string) (ret bool) {

	for i := 0; i < len(lettersFromUser); i++ {
		if letter == lettersFromUser[i] {
			return true
		}
	}

	return false
}

// New search function - previous contains function not working for longer strings
func bruteForceSearch(letters []string) (ret []string) {
	// loop through all words in the dictionary
	var matchedWords []string
	var isLetterInWord bool
	for _, word := range words {
		// filter the words that are over number of letters provided by user
		if len(word) <= len(letters) {
			// check each word to see if it contains the letter
			//convert string to array
			var remainingLetters []string
			lettersInWord := strings.Split(word, "")
			remainingLetters = letters

			fmt.Println(lettersInWord)
			isLetterInWord = false

			// loop through the word and remove from the lettersinWord if it appears
			for i := 0; i < len(lettersInWord); i++ {
				// check each letter in the word
				isLetterInWord = checkLetter(lettersInWord[i], remainingLetters)
				if !isLetterInWord {
					break
				} else {
					remainingLetters = removeIndex(remainingLetters, lettersInWord[i])
					isLetterInWord = true
				}
			}

			if isLetterInWord {
				matchedWords = append(matchedWords, word)
			}

		}
	}
	return matchedWords
}

// Filter all of the words initially
func initialFilter(letters []string) (ret []string) {
	// step 1 - filter all the words with letter[i] in them
	var filteredWords []string
	filteredWords = words

	// Loop through each user input letter
	for _, letter := range letters {
		var newFilteredWords []string

		// Loop through each word in the dictionary
		for _, s := range filteredWords {
			// If the word contains the letter then add to the filtered words
			if strings.Contains(s, letter) {
				// exclude the word if its longer than the number of letters
				if len(letters) >= len(s) {
					newFilteredWords = append(newFilteredWords, s)
				}

			}
		}
		filteredWords = newFilteredWords
	}
	return filteredWords
}

func FindWords(letters []string) (ret []string) {
	// go through each of the letters and see which words contain
	// the letters

	//filteredWords := initialFilter(letters)
	filteredWords := bruteForceSearch(letters)

	return filteredWords
}

// Load the dictionary and sort from largest to smallest
func LoadDictionary() {
	//unpack the json file
	file, _ := os.ReadFile("data/dict-test.json")
	//file, _ := os.ReadFile("data/dictionary.json")
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
}
