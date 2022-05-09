package word

import (
	"reflect"
	"strings"
	"testing"
)

func TestWord_ComputeStatus(t *testing.T) {
	type fields struct {
		Value string
	}
	type args struct {
		wordToFind Word
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantWordStatus Status
	}{
		{
			name:           "test GG",
			fields:         fields{"TO"},
			args:           args{Word{Value: "TO"}},
			wantWordStatus: Status{Statuses: []LetterStatus{GoodPlace, GoodPlace}, Letters: strings.Split("TO", "")},
		},
		{
			name:           "test GN",
			fields:         fields{"TO"},
			args:           args{Word{Value: "TA"}},
			wantWordStatus: Status{Statuses: []LetterStatus{GoodPlace, NotPresent}, Letters: strings.Split("TO", "")},
		},
		{
			name:           "test BB",
			fields:         fields{"TO"},
			args:           args{Word{Value: "OT"}},
			wantWordStatus: Status{Statuses: []LetterStatus{BadPlace, BadPlace}, Letters: strings.Split("TO", "")},
		},
		{
			name:           "test",
			fields:         fields{"THOMAS"},
			args:           args{Word{Value: "TARTES"}},
			wantWordStatus: Status{Statuses: []LetterStatus{GoodPlace, NotPresent, NotPresent, NotPresent, BadPlace, GoodPlace}, Letters: strings.Split("THOMAS", "")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Word{
				Value: tt.fields.Value,
			}
			if gotWordStatus := w.ComputeStatus(tt.args.wordToFind.Value); !reflect.DeepEqual(gotWordStatus, tt.wantWordStatus) {
				t.Errorf("ComputeStatus() = %v, want %v", gotWordStatus, tt.wantWordStatus)
			}
		})
	}
}

func Test_removeElement(t *testing.T) {
	type args struct {
		elements        []string
		elementToRemove string
	}
	tests := []struct {
		name         string
		args         args
		wantNewArray []string
	}{
		{
			name:         "toto",
			args:         args{elements: []string{"TOTO", "TATA"}, elementToRemove: "TOTO"},
			wantNewArray: []string{"TATA"},
		},
		{
			name:         "toto",
			args:         args{elements: []string{"TOTO", "TATA", "TITI"}, elementToRemove: "TATA"},
			wantNewArray: []string{"TOTO", "TITI"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewArray := removeElement(tt.args.elements, tt.args.elementToRemove); !reflect.DeepEqual(gotNewArray, tt.wantNewArray) {
				t.Errorf("removeElement() = %v, want %v", gotNewArray, tt.wantNewArray)
			}
		})
	}
}

func TestStatus_GetReducedWordsToParse(t *testing.T) {
	type fields struct {
		Statuses  []LetterStatus
		Letters   []string
		Compliant bool
	}
	type args struct {
		wordsToParse []string
	}
	tests := []struct {
		name                    string
		fields                  fields
		args                    args
		wantReducedWordsToParse []string
	}{
		{
			name:                    "toto",
			fields:                  fields{Statuses: []LetterStatus{GoodPlace}, Letters: []string{"A", "B"}, Compliant: false},
			args:                    args{wordsToParse: []string{"AB", "AC", "AD"}},
			wantReducedWordsToParse: nil,
		},
		{
			name:                    "toto",
			fields:                  fields{Statuses: []LetterStatus{GoodPlace, GoodPlace}, Letters: []string{"A", "B"}, Compliant: false},
			args:                    args{wordsToParse: []string{"AB", "AC", "AD"}},
			wantReducedWordsToParse: []string{"AB"},
		},
		{
			name:                    "toto",
			fields:                  fields{Statuses: []LetterStatus{GoodPlace, NotPresent}, Letters: []string{"A", "B"}, Compliant: false},
			args:                    args{wordsToParse: []string{"AB", "AC", "AD"}},
			wantReducedWordsToParse: []string{"AC", "AD"},
		},
		{
			name:                    "toto",
			fields:                  fields{Statuses: []LetterStatus{GoodPlace, BadPlace, NotPresent}, Letters: []string{"A", "B", "C"}, Compliant: false},
			args:                    args{wordsToParse: []string{"ABC", "ACB", "ADC", "ADB", "AEB"}},
			wantReducedWordsToParse: []string{"ADB", "AEB"},
		},
		{
			name:                    "toto",
			fields:                  fields{Statuses: []LetterStatus{GoodPlace, BadPlace, NotPresent, BadPlace, BadPlace}, Letters: []string{"T", "O", "T", "O", "S"}, Compliant: false},
			args:                    args{wordsToParse: []string{"TSOBO", "TSOCO", "TOTOS"}},
			wantReducedWordsToParse: []string{"TSOBO", "TSOCO"},
		},
		{
			name:                    "toto",
			fields:                  fields{Statuses: []LetterStatus{GoodPlace, GoodPlace, GoodPlace, GoodPlace, GoodPlace, GoodPlace, GoodPlace, NotPresent, NotPresent}, Letters: []string{"E", "B", "R", "A", "N", "C", "H", "A", "T"}, Compliant: false},
			args:                    args{wordsToParse: []string{"EBRAISONS", "EBRANCHAI", "EBRANCHAS", "EBRANCHAT", "EBRANCHEE", "EBRANCHER", "EBRANCHES", "EBRANCHEZ", "EBRANLAIS", "EBRANLAIT", "EBRANLANT", "EBRANLEES", "EBRANLENT", "EBRANLERA", "EBRANLIEZ", "EBRANLONS", "EBRASAMES", "EBRASASSE", "EBRASATES", "EBRASERAI", "EBRASERAS", "EBRASEREZ", "EBRASIONS", "EBRASURES", "EBRECHAIS", "EBRECHAIT", "EBRECHANT", "EBRECHEES", "EBRECHENT", "EBRECHERA", "EBRECHIEZ", "EBRECHONS", "EBRECHURE", "EBROUAMES", "EBROUASSE", "EBROUATES", "EBROUDIES", "EBROUDIRA", "EBROUERAI", "EBROUERAS", "EBROUEREZ", "EBROUIONS", "EBRUITAIS", "EBRUITAIT", "EBRUITANT", "EBRUITEES", "EBRUITENT", "EBRUITERA", "EBRUITIEZ", "EBRUITONS", "EBRUTAMES", "EBRUTASSE", "EBRUTATES", "EBRUTERAI", "EBRUTERAS", "EBRUTEREZ", "EBRUTIONS", "EBURNEENS", "EBURONNES", "ECACHAMES", "ECACHASSE", "ECACHATES", "ECACHERAI", "ECACHERAS", "ECACHEREZ", "ECACHIONS", "ECAFFAMES", "ECAFFASSE", "ECAFFATES", "ECAFFERAI", "ECAFFERAS", "ECAFFEREZ", "ECAFFIONS", "ECAILLAGE", "ECAILLAIS", "ECAILLAIT", "ECAILLANT", "ECAILLEES", "ECAILLENT", "ECAILLERA", "ECAILLERE", "ECAILLERS", "ECAILLEUX", "ECAILLIEZ", "ECAILLONS", "ECAILLURE", "ECALAIENT", "ECALASSES", "ECALERAIS", "ECALERAIT", "ECALERENT", "ECALERIEZ", "ECALERONS", "ECALERONT", "ECANGUAIS", "ECANGUAIT", "ECANGUANT", "ECANGUEES", "ECANGUENT", "ECANGUERA", "ECANGUEUR", "ECANGUIEZ", "ECANGUONS", "ECARLATES", "ECARTAMES", "ECARTASSE", "ECARTATES", "ECARTELAI", "ECARTELAS", "ECARTELAT", "ECARTELEE", "ECARTELER", "ECARTELES", "ECARTELEZ", "ECARTERAI", "ECARTERAS", "ECARTEREZ", "ECARTEURS", "ECARTIONS", "ECATIRAIS", "ECATIRAIT", "ECATIRENT", "ECATIRIEZ", "ECATIRONS", "ECATIRONT", "ECATISSES", "ECATISSEZ", "ECBALLIUM", "ECCHYMOSE", "ECCLESIAL", "ECDYSONES", "ECERVELEE", "ECERVELES", "ECHAFAUDA", "ECHAFAUDE"}},
			wantReducedWordsToParse: []string{"EBRANCHEE", "EBRANCHER", "EBRANCHES", "EBRANCHEZ"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Status{
				Statuses:  tt.fields.Statuses,
				Letters:   tt.fields.Letters,
				Compliant: tt.fields.Compliant,
			}
			if gotReducedWordsToParse := s.GetReducedWordsToParse(tt.args.wordsToParse); !reflect.DeepEqual(gotReducedWordsToParse, tt.wantReducedWordsToParse) {
				t.Errorf("GetReducedWordsToParse() = %v, want %v", gotReducedWordsToParse, tt.wantReducedWordsToParse)
			}
		})
	}
}

func Benchmark_GetReducedWordsToParse(b *testing.B) {
	s := Status{
		Statuses:  []LetterStatus{GoodPlace, GoodPlace, GoodPlace, GoodPlace, GoodPlace, GoodPlace, GoodPlace, NotPresent, NotPresent},
		Letters:   []string{"E", "B", "R", "A", "N", "C", "H", "A", "T"},
		Compliant: false,
	}
	for i := 0; i < b.N; i++ {
		words := []string{"EBRAISONS", "EBRANCHAI", "EBRANCHAS", "EBRANCHAT", "EBRANCHEE", "EBRANCHER", "EBRANCHES", "EBRANCHEZ", "EBRANLAIS", "EBRANLAIT", "EBRANLANT", "EBRANLEES", "EBRANLENT", "EBRANLERA", "EBRANLIEZ", "EBRANLONS", "EBRASAMES", "EBRASASSE", "EBRASATES", "EBRASERAI", "EBRASERAS", "EBRASEREZ", "EBRASIONS", "EBRASURES", "EBRECHAIS", "EBRECHAIT", "EBRECHANT", "EBRECHEES", "EBRECHENT", "EBRECHERA", "EBRECHIEZ", "EBRECHONS", "EBRECHURE", "EBROUAMES", "EBROUASSE", "EBROUATES", "EBROUDIES", "EBROUDIRA", "EBROUERAI", "EBROUERAS", "EBROUEREZ", "EBROUIONS", "EBRUITAIS", "EBRUITAIT", "EBRUITANT", "EBRUITEES", "EBRUITENT", "EBRUITERA", "EBRUITIEZ", "EBRUITONS", "EBRUTAMES", "EBRUTASSE", "EBRUTATES", "EBRUTERAI", "EBRUTERAS", "EBRUTEREZ", "EBRUTIONS", "EBURNEENS", "EBURONNES", "ECACHAMES", "ECACHASSE", "ECACHATES", "ECACHERAI", "ECACHERAS", "ECACHEREZ", "ECACHIONS", "ECAFFAMES", "ECAFFASSE", "ECAFFATES", "ECAFFERAI", "ECAFFERAS", "ECAFFEREZ", "ECAFFIONS", "ECAILLAGE", "ECAILLAIS", "ECAILLAIT", "ECAILLANT", "ECAILLEES", "ECAILLENT", "ECAILLERA", "ECAILLERE", "ECAILLERS", "ECAILLEUX", "ECAILLIEZ", "ECAILLONS", "ECAILLURE", "ECALAIENT", "ECALASSES", "ECALERAIS", "ECALERAIT", "ECALERENT", "ECALERIEZ", "ECALERONS", "ECALERONT", "ECANGUAIS", "ECANGUAIT", "ECANGUANT", "ECANGUEES", "ECANGUENT", "ECANGUERA", "ECANGUEUR", "ECANGUIEZ", "ECANGUONS", "ECARLATES", "ECARTAMES", "ECARTASSE", "ECARTATES", "ECARTELAI", "ECARTELAS", "ECARTELAT", "ECARTELEE", "ECARTELER", "ECARTELES", "ECARTELEZ", "ECARTERAI", "ECARTERAS", "ECARTEREZ", "ECARTEURS", "ECARTIONS", "ECATIRAIS", "ECATIRAIT", "ECATIRENT", "ECATIRIEZ", "ECATIRONS", "ECATIRONT", "ECATISSES", "ECATISSEZ", "ECBALLIUM", "ECCHYMOSE", "ECCLESIAL", "ECDYSONES", "ECERVELEE", "ECERVELES", "ECHAFAUDA", "ECHAFAUDE"}
		s.GetReducedWordsToParse(words)
	}
}

func Benchmark_GetReducedWordsToParse4A(b *testing.B) {
	s := Status{
		Statuses:  []LetterStatus{GoodPlace, BadPlace, NotPresent, NotPresent},
		Letters:   []string{"A", "C", "N", "E"},
		Compliant: false,
	}
	for i := 0; i < b.N; i++ {
		words := []string{
			"ABLE",
			"ABOI",
			"ABOT",
			"ABRI",
			"ABUS",
			"ACAI",
			"ACCU",
			"ACES",
			"ACHE",
			"ACME",
			"ACNE",
			"ACON",
			"ACRA",
			"ACRE",
			"ACTA",
			"ACTE",
			"ACTU",
			"ACUL",
			"ADAC",
			"ADAS",
			"ADAV",
			"ADNE",
			"ADON",
			"ADOS",
			"AEDE",
			"AERA",
			"AERE",
			"AFAR",
			"AFAT",
			"AFIN",
			"AFRO",
			"AGAS",
			"AGEE",
			"AGES",
			"AGHA",
			"AGIE",
			"AGIO",
			"AGIR",
			"AGIS",
			"AGIT",
			"AGNI",
			"AGUI",
			"AHAN",
			"AIDA",
			"AIDE",
			"AIDS",
			"AIES",
			"AIGU",
			"AILE",
			"AILS",
			"AIMA",
			"AIME",
			"AINE",
			"AIRA",
			"AIRE",
			"AIRS",
			"AISE",
			"AISY",
			"AJUT",
			"AKAN",
			"ALEA",
			"ALEM",
			"ALES",
			"ALFA",
			"ALLA",
			"ALLE",
			"ALLO",
			"ALOI",
			"ALPA",
			"ALPE",
			"ALTI",
			"ALTO",
			"ALUN",
			"ALUS",
			"ALYA",
			"AMAN",
			"AMAS",
			"AMEN",
			"AMER",
			"AMES",
			"AMIE",
			"AMIS",
			"AMMI",
			"AMOK",
			"AMUI",
			"ANAL",
			"ANAR",
			"ANAS",
			"ANEE",
			"ANEL",
			"ANGE",
			"ANIL",
			"ANIS",
			"ANKH",
			"ANON",
			"ANSE",
			"ANTE",
			"ANUS",
			"AOUT",
			"APAX",
			"APEX",
			"APIS",
			"APPS",
			"APRE",
			"APTE",
			"ARAC",
			"ARAK",
			"ARAS",
			"ARCH",
			"ARCS",
			"ARDU",
			"AREC",
			"AREG",
			"ARES",
			"AREU",
			"ARGH",
			"ARIA",
			"ARMA",
			"ARME",
			"AROL",
			"ARTS",
			"ARTY",
			"ARUM",
			"ASES",
			"ASIC",
			"ASIN",
			"ASPE",
			"ASPI",
			"ASSE",
			"ASSO",
			"ASTE",
			"ASTI",
			"ATRE",
			"AUBE",
			"AUGE",
			"AULA",
			"AULX",
			"AUNA",
			"AUNE",
			"AURA",
			"AUTO",
			"AVAL",
			"AVEC",
			"AVEN",
			"AVES",
			"AVEU",
			"AVEZ",
			"AVIS",
			"AXAI",
			"AXAS",
			"AXAT",
			"AXEE",
			"AXEL",
			"AXER",
			"AXES",
			"AXEZ",
			"AXIS",
			"AXOA",
			"AYEZ",
			"AZUR"}
		reduced := s.GetReducedWordsToParse(words)
		println(len(reduced))
	}
}
