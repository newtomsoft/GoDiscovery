package word

import "testing"

func Test_wordGivingMaxInformation(t *testing.T) {
	type args struct {
		candidateWords []string
		wordsToParse   []string
	}
	tests := []struct {
		name         string
		args         args
		wantBestWord string
	}{
		{
			name: "test",
			args: args{
				candidateWords: []string{"QANUN", "QIBLA", "QINGS", "QUADO", "QUADS", "QUAIS", "QUAND", "QUANT", "QUARK", "QUART", "QUASI", "QUBIT", "QUEER", "QUELS", "QUENA", "QUETA", "QUETE", "QUEUE", "QUEUX", "QUICK", "QUIET", "QUILT", "QUINE", "QUINT", "QUIPO", "QUIPU", "QUOTA"},
				wordsToParse:   []string{"QANUN", "QIBLA", "QINGS", "QUADO", "QUADS", "QUAIS", "QUAND", "QUANT", "QUARK", "QUART", "QUASI", "QUBIT", "QUEER", "QUELS", "QUENA", "QUETA", "QUETE", "QUEUE", "QUEUX", "QUICK", "QUIET", "QUILT", "QUINE", "QUINT", "QUIPO", "QUIPU", "QUOTA"},
			},
			wantBestWord: "QUIET",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotBestWord := wordGivingMaxInformation(tt.args.candidateWords, tt.args.wordsToParse); gotBestWord != tt.wantBestWord {
				t.Errorf("wordGivingMaxInformation() = %v, want %v", gotBestWord, tt.wantBestWord)
			}
		})
	}
}
