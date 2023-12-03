package model

import (
	"encoding/json"
	"os"
	"sort"
	"strings"
)

var words []string
var dictionary map[string]interface{}
var definitions map[string]interface{}

type Dictionary struct {
	word       string
	definition string
}

// New search function - previous contains function not working for longer strings
func bruteForceSearch(letters []string) (ret []string) {

	var matchedWords []string
	var isLetterInWord bool

	// loop through all words in the dictionary
	for _, word := range words {
		// filter the words that are over number of letters provided by user
		if len(word) <= len(letters) {
			// check each word to see if it contains the letter
			//convert string to array
			remainingLetters := append([]string(nil), letters...)
			lettersInWord := strings.Split(word, "")

			//fmt.Println(remainingLetters)
			isLetterInWord = false

			// loop through the word and remove from the lettersinWord if it appears
			for i := 0; i < len(lettersInWord); i++ {
				// check each letter in the word
				isLetterInWord = checkLetter(lettersInWord[i], remainingLetters)
				if !isLetterInWord {
					break
				} else {
					// remove letters from the remaining
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

func FindWords(letters []string) ([]string, []string) {
	// go through each of the letters and see which words contain
	// the letters
	matchedWords := bruteForceSearch(letters)

	// Change the dictionary so that only top answers are returned
	filteredWords := matchedWords[:5]

	// send the definitions back to the handler
	var definitions []string
	for _, word := range filteredWords {
		definition := dictionary[word].(string)
		definitions = append(definitions, definition)
	}

	return filteredWords, definitions
}

// Load the dictionary and sort from largest to smallest
func LoadDictionary() {
	//unpack the json file
	//file, _ := os.ReadFile("data/dict-test.json")
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

}

// ########################
// ### HELPER FUNCTIONS ###
// ########################

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
