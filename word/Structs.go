package word

type LetterStatus int8

const (
	NotPresent LetterStatus = 0
	BadPlace   LetterStatus = 1
	GoodPlace  LetterStatus = 2
)

type Word struct {
	Value string
}
