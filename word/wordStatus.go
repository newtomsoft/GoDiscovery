package word

import (
	"strings"
)

type Status struct {
	Statuses  []LetterStatus
	Letters   []string
	Compliant bool
}

func (w Word) ComputeStatus(wordToFind string) (wordStatus Status) {
	wordStatus.Statuses = make([]LetterStatus, len(w.Value))
	allGoodIndexes := make([]int, len(w.Value))
	for i := 0; i < len(allGoodIndexes); i++ {
		allGoodIndexes[i] = -1
	}
	allGoodLetters := make([]string, len(w.Value))
	for i := 0; i < len(wordToFind); i++ {
		if w.Value[i] == wordToFind[i] {
			wordStatus.Statuses[i] = GoodPlace
			allGoodIndexes = append(allGoodIndexes, i)
			allGoodLetters = append(allGoodLetters, string(w.Value[i]))
		}
	}
	for i := 0; i < len(wordToFind); i++ {
		if isExist(allGoodIndexes, i) {
			continue
		}
		letter := string(w.Value[i])
		inWordCount := strings.Count(wordToFind, letter)
		foundCount := count(allGoodLetters, letter)
		if foundCount == inWordCount {
			wordStatus.Statuses[i] = NotPresent
		} else {
			wordStatus.Statuses[i] = BadPlace
			allGoodLetters = append(allGoodLetters, letter)
		}
	}
	wordStatus.Letters = strings.Split(w.Value, "")
	return
}

func (s Status) GetReducedWordsToParse(wordsToParse []string) (reducedWordsToParse []string) {
	if len(s.Statuses) != len(wordsToParse[0]) {
		return
	}

	var lettersFound []string
	tempWordsToParse := wordsToParse
	for i, letterStatus := range s.Statuses {
		if letterStatus != GoodPlace {
			continue
		}
		for _, w := range wordsToParse {
			if string(w[i]) != s.Letters[i] {
				tempWordsToParse = removeElement(tempWordsToParse, w)
			}
		}
		lettersFound = append(lettersFound, s.Letters[i])
	}

	reducedWordsToParse = tempWordsToParse
	for i, letterStatus := range s.Statuses {
		if letterStatus != BadPlace {
			continue
		}
		minNumberToCount := 1
		for _, letter := range lettersFound {
			if letter == s.Letters[i] {
				minNumberToCount++
			}
		}
		for _, w := range reducedWordsToParse {
			if string(w[i]) == s.Letters[i] || strings.Count(w, s.Letters[i]) < minNumberToCount {
				tempWordsToParse = removeElement(tempWordsToParse, w)
			}
		}
		lettersFound = append(lettersFound, s.Letters[i])
	}

	reducedWordsToParse = tempWordsToParse
	for i, letterStatus := range s.Statuses {
		if letterStatus != NotPresent {
			continue
		}
		numberToCount := 0
		for _, letter := range lettersFound {
			if letter == s.Letters[i] {
				numberToCount++
			}
		}
		for _, w := range reducedWordsToParse {
			if string(w[i]) == s.Letters[i] || strings.Count(w, s.Letters[i]) != numberToCount {
				tempWordsToParse = removeElement(tempWordsToParse, w)
			}
		}
	}
	reducedWordsToParse = tempWordsToParse
	return
}

func isExist(values []int, value int) bool {
	for _, v := range values {
		if v == value {
			return true
		}
	}
	return false
}

func count(values []string, string string) (number int) {
	for _, s := range values {
		if s == string {
			number++
		}
	}
	return
}

func removeElement(elements []string, elementToRemove string) (newArray []string) {
	newLength := 0
	newArray = make([]string, len(elements))
	for i, element := range elements {
		if element == elementToRemove {
			continue
		}
		newArray[newLength] = elements[i]
		newLength++
	}
	newArray = newArray[:newLength]
	return
}
