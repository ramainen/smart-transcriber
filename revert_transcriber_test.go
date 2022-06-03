package transcriber

import "testing"

func TestTranscriber_AddKnownPhrase(t *testing.T) {

	transciber := NewTranscriber()

	cases := []struct {
		phrase  string
		russian string
		english string
	}{
		{"перфоратор makita", "мейкита", "makita"},
		{"пылесос  Bosch", "бош", "bosch"},
		{"смартфон-apple", "аппл", "apple"},
	}
	for _, oneCase := range cases {
		transciber.AddKnownPhrase(oneCase.phrase)
		if transciber.KnownRuToEnDict[oneCase.russian] != oneCase.english {
			t.Error("Word not found", oneCase.russian)
		}
	}
	if _, isset := transciber.KnownWords["пылесос"]; !isset {
		t.Error("Original word not found", "пылесос")
	}

}

func TestTranscriber_BeautifyQuery(t *testing.T) {

	transciber := NewTranscriber()

	database := []struct {
		phrase string
	}{
		{"перфоратор makita"},
		{"пылесос  Bosch"},
		{"смартфон-apple"},
		{"Зелда буги"},
		{"zelda буги"},
		{"плита OSB"}, //Uppercase
	}
	for _, oneCase := range database {
		transciber.AddKnownPhrase(oneCase.phrase)
	}

	cases := []struct {
		phrase       string
		betterPhrase string
	}{
		{"макита перфоратор", "makita перфоратор"},
		{"ковыряка бош", "ковыряка bosch"},
		{"ковыряка бош буги", "ковыряка bosch буги"},
		{"смартфон-аппл", "смартфон apple"},
		{"макита", "makita"},
		{"зелда буги", "зелда буги"},  //Do not change known words
		{"зельда буги", "zelda буги"}, //... but change if mistake
		{"плита осб", "плита osb"},
		{"плита осбулька", "плита осбулька"},                                   //Do not change part of word
		{"Рамаммбахарум мамбурум 1233-B12", "Рамаммбахарум мамбурум 1233-B12"}, //Do not change unknown words
	}
	for _, oneCase := range cases {
		if result := transciber.BeautifyQuery(oneCase.phrase); result != oneCase.betterPhrase {
			t.Error("Phrase failed", oneCase.phrase, "GOT: ", result, "EXPECTS: ", oneCase.betterPhrase)
		}
	}

}
