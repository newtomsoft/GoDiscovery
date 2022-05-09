package dictionary

import (
	"bufio"
	"fmt"
	"os"
)

func GetFirstWord(firstLetter string, length int) string {
	for _, word := range getFirstsWords() {
		if string(word[0]) == firstLetter && len(word) == length {
			return word
		}
	}
	return ""
}

func GetAllWordsMap() map[string]map[int][]string {
	allWords := make(map[string]map[int][]string)
	for _, letter := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		allWords[string(letter)] = make(map[int][]string)
	}
	for _, word := range getAllWords() {
		var letter = word[0]
		var length = len(word)
		allWords[string(letter)][length] = append(allWords[string(letter)][length], word)
	}
	return allWords
}

func getFirstsWords() []string {
	readFile, err := os.Open("words_4-12.txt")
	if err != nil {
		fmt.Println(err)
	}
	words := readWords(readFile)
	return words
}

func getAllWords() []string {
	readFile, err := os.Open("words_4-12.txt")
	if err != nil {
		fmt.Println(err)
	}
	words := readWords(readFile)
	return words
}

func readWords(readFile *os.File) []string {
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var words []string
	for fileScanner.Scan() {
		words = append(words, fileScanner.Text())
	}
	readFile.Close()
	return words
}
