package word

import (
	"math"
)

func GetBestWord(wordsToParse []string, maxInformation bool) string {
	if maxInformation {
		return wordGivingMaxInformation(wordsToParse, wordsToParse)
	}
	return ""
}

func wordGivingMaxInformation(candidateWords []string, wordsToParse []string) (bestWord string) {
	minSum := math.MaxInt
	for _, word := range wordsToParse {
		realWord := Word{word}
		print(word + " ")
		sum := 0
		for _, w := range candidateWords {
			status := realWord.ComputeStatus(w)
			parse := status.GetReducedWordsToParse(wordsToParse)
			sum += len(parse)
		}
		if sum < minSum {
			minSum = sum
			bestWord = word
		}
	}
	println(bestWord)
	return
}
