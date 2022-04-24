package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	firstLetter := "B"
	length := 7
	var _ = getAllWordsMap()[firstLetter][length]
	word := getFirstWord(firstLetter, length)
	print(word)
}

func getFirstWord(firstLetter string, length int) string {
	for _, word := range getFirstWords() {
		if string(word[0]) == firstLetter && len(word) == length {
			return word
		}
	}
	return ""
}

func getFirstWords() []string {
	readFile, err := os.Open("words_4-12.txt")
	if err != nil {
		fmt.Println(err)
	}
	words := readWords(readFile)
	return words
}

func getAllWordsMap() map[string]map[int][]string {
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
