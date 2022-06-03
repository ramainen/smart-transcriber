package transcriber

import (
	"testing"
)

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func TestTranscriber_Transcribe(t *testing.T) {
	transciber := NewTranscriber()

	cases := []struct {
		english string
		russian string
	}{
		{"makita", "мейкита"},
		{"makita", "макита"},
		{"bosch", "бош"},
		{"osb", "осб"},
		{"bergauf", "бергауф"},
		{"apple", "аппле"},
		{"philips", "пхилипс"},
		{"philips", "филипс"},
		{"zelda", "зельда"},
		{"bmw", "биэмдаблю"},
		{"bmw", "бмв"},
		{"knauf", "кнауф"},
		{"wolfenstein", "волфенстеин"},
		{"wolfenstein", "волфенстейн"},
		{"wolfenstein", "вольфенстеин"},
		{"wolfenstein", "вольфенштайн"},

		//Double letters cases
		{"rossinka", "росинка"},
		{"sheffilton", "шефилтон"},
	}
	for _, oneCase := range cases {
		if result := transciber.Transcribe(oneCase.english); !contains(result, oneCase.russian) {

			t.Error("Phrase failed", oneCase.english, "GOT: ", result, "EXPECTS: ", oneCase.russian)
		}
	}

}
