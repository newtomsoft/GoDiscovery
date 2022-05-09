package main

import (
	"discovery/dictionary"
	"discovery/word"
)

func main() {
	firstLetter := "W"
	length := 4

	_ = dictionary.GetAllWordsMap()[firstLetter][length]
	_ = dictionary.GetFirstWord(firstLetter, length)

	const allLetters = "A"
	for _, letter := range allLetters {
		var wordsToParse = dictionary.GetAllWordsMap()[string(letter)][length]
		word.GetBestWord(wordsToParse, true)
	}

}
