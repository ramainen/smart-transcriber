package transcriber

import (
	"regexp"
	"strings"
)

// Product: "перфоратор makita" -> ["перфоратор"] + ["макита"=>"makita", "мейкита"=>"makita"].,
// Search: "перфоратор макита" -> "перфоратор makita".
// Search: "фм-трансмиттер" -> "fm трансмиттер".

// AddKnownPhrase adds another phrase with Russian and English words, for example, name of product
// and stores only russian words (without letters) in internal memory.
// also saves English trasliterations
func (obj *Transcriber) AddKnownPhrase(phrase string) {

	words := Splitter(strings.ToLower(phrase), " -,")
	isset := false
	for _, word := range words {
		if res := obj.onlyEnglishRegexp.Match([]byte(word)); res {

			if _, isset = obj.KnownRuToEnDict[word]; isset {
				continue
			}

			for _, newRussianWord := range obj.Transcribe(word) {
				obj.KnownRuToEnDict[newRussianWord] = word
			}
		}
		obj.KnownWords[word] = 1 //We know about every word
	}
}

func (obj *Transcriber) BeautifyQuery(phrase string) string {

	resultPhrase := phrase
	words := Splitter(strings.ToLower(phrase), " -")
	isset := false
	for _, word := range words {
		//Known words do not changes
		wordLower := strings.ToLower(word)
		if _, isset = obj.KnownWords[wordLower]; isset {
			continue
		}

		if english, isset := obj.KnownRuToEnDict[wordLower]; isset {

			var re = regexp.MustCompile(`[, ^\-]?` + word + `[$, \-]?`)
			resultPhrase = re.ReplaceAllString(resultPhrase, ` `+english+` `)

		}

	}
	return strings.Trim(resultPhrase, " ")
}

func Splitter(s string, splits string) []string {
	m := make(map[rune]int)
	for _, r := range splits {
		m[r] = 1
	}

	splitter := func(r rune) bool {
		return m[r] == 1
	}

	return strings.FieldsFunc(s, splitter)
}
