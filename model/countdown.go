package model

import (
	"encoding/json"
	"os"
	"sort"
	"strings"

	"golang.org/x/exp/slices"
)

var words []string
var dictionary map[string]interface{}

func removeIndex(s []string, index string) []string {
	var indexInt int
	for i := 0; i <= len(s); i++ {
		if s[i] == index {
			return append(s[:indexInt], s[indexInt+1:]...)
		}
	}
	return nil

}

func initialFilter(letters []string) (ret []string) {
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
	return filteredWords
}

func FindWords(letters []string) (ret []string) {
	// go through each of the letters and see which words contain
	// the letters

	filteredWords := initialFilter(letters)

	// step 2 - These are all possibilities of words. Confirm that they are correct
	// for each word in the list, remove the letter one at a time
	var returnedList []string
	for _, word := range filteredWords {
		// remove the letters from the word 1 by one
		var letterTest []string
		letterTest = letters

		// Go through the letter in the word one by one
		lettersInWord := strings.Split(word, "")

		for _, letter := range lettersInWord {
			if slices.Contains(letterTest, letter) {
				//remove from letterTest
				letterTest = removeIndex(letterTest, letter)
			}
		}
		if letterTest != nil {
			returnedList = append(returnedList, word)
		}

	}

	// return the word
	return returnedList
}

// Load the dictionary and sort from largest to smallest
func LoadDictionary() {
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
}
