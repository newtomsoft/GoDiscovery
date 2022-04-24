package main

func main() {
	const allLetters = "VUTSR"
	const length = 8
	for _, letter := range allLetters {
		var wordsToParse = getAllWordsMap()[string(letter)][length]

		getBestWord(wordsToParse, true)
	}

}

func getBestWord(wordsToParse []string, b bool) {
	wordGivingMaxInformation(wordsToParse, wordsToParse)
}

func wordGivingMaxInformation(bestWords []string, wordsToParse []string) {

	var stayingWordsSumByWord = make(map[string]int)
	for _, word := range wordsToParse {
		print(word)
		sum := 0
		for _, w := range bestWords {
			sum += getReducedWordsToParse(word.GetStatusWith(w), wordsToParse).Count()
		}
	}
	var stayingWordsSumMin = stayingWordsSumByWord.Min(item => item.Value);
	var candidatesWords = stayingWordsSumByWord.Where(item => item.Value == stayingWordsSumMin).Select(item => item.Key).ToList();
	var firstCandidate = candidatesWords.Find(bestWordsArray.Contains);
	return firstCandidate ?? candidatesWords[0];
}

func getReducedWordsToParse(i interface{}, parse []string) []string {
	value := []string {"toto", "tata"}
	return value
}
