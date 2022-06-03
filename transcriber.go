package transcriber

import (
	"encoding/csv"
	"log"
	"os"
	"path/filepath"

	nativeregexp "regexp"

	regexp "github.com/scorpionknifes/go-pcre"
)

type Transcriber struct {
	Dict                  map[string][]string
	KnownRuToEnDict       map[string]string
	onlyEnglishRegexp     *nativeregexp.Regexp
	KnownWords            map[string]int
	MultilangRulesRegexps [][]regexp.Regexp
	MultilangRules        [][][]string

	AbbrRules        [][]string
	AbbrRulesRegexps []regexp.Regexp

	DoubleRules        [][]string
	DoubleRulesRegexps []*nativeregexp.Regexp
}

func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
func NewTranscriber() Transcriber {
	newTranscriber := Transcriber{}
	newTranscriber.Dict = map[string][]string{}
	newTranscriber.KnownRuToEnDict = map[string]string{}
	newTranscriber.KnownWords = map[string]int{}

	newTranscriber.onlyEnglishRegexp = nativeregexp.MustCompile(`^[a-z]+$`)

	newTranscriber.MultilangRules = [][][]string{

		////////////////////////////
		//NAIVE

		{
			{"sch", "ш"}, {"ya", "я"}, {`\Aand\Z`, "энд"}, {`\Ay(?=[euioa])`, "й"},
			{`\Ae(?![euioay-])`, "э"}, {`(?<=[euioa])y\Z`, "й"}, {`(?<![euioa])y\Z`, "и"}, {"(?<=[euioa])c(?=[euioay-])", "с"},
			{"a", "а"}, {"e", "е"}, {"o", "о"}, {"b", "б"}, {"c", "к"}, {"d", "д"}, {"f", "ф"}, {"g", "г"}, {"h", "х"}, {"j", "й"}, {"k", "к"}, {"l", "л"}, {"m", "м"}, {"n", "н"}, {"p", "п"}, {"q", "к"}, {"r", "р"}, {"s", "с"}, {"t", "т"}, {"v", "в"}, {"w", "в"}, {"x", "кс"}, {"z", "з"}, {"a", "а"}, {"e", "е"}, {"i", "и"}, {"u", "у"}, {"o", "о"}, {"y", "и"}, {"ß", "сс"}, {"ö", "о"}, {"ä", "а"}, {"ü", "у"}, {"é", "е"}, {"è", "е"}, {"à", "а"}, {"ù", "у"}, {"ê", "е"}, {"â", "а"}, {"ô", "о"}, {"î", "и"}, {"û", "у"}, {"ë", "е"}, {"ï", "и"}, {"ÿ", "и"}, {"ç", "с"},
		},

		//EN
		{
			{`ay`, `ей`}, {`[ao]ught`, `от`}, {`ueu`, `е`},
			{`you`, `ю`}, {`chr`, `кр`}, {`scq`, `ск`},
			{`(?<=[euioay])s(?=[euioay])`, `з`},
			{`ts(?=\w+)`, `ц`},
			{`c(?=[eiy])`, `с`}, {`(?<!g)g(?=[eiy])`, `дж`},
			{`a(?=[uw])`, `о`},
			{`(?<=w)a(?=r)`, `о`},
			{`\Awr(?=[euioay])`, `р`},
			{`j(?![euioay])`, `ж`},
			{`(?<![euioay])th`, `т`},
			{`(?<!e)ew`, `ью`},
			{`ow(?=\w{2}){`, `оу`},
			{`l(?=t)`, `ль`},
			{`ea(?=(d|th|lth|sure|sant))`, `е`},
			{`t(?=(ure|ural|ury))`, `ч`},
			{`(?<=[^euioay])y(?=\w[euioay])`, `ай`},
			{`(?<=[^euioay])a(?=\w[euioay])`, `ей`},
			{`(?<=[^euioay])i(?=\we\Z)`, `ай`},
			{`\Ae(?![euioay-])`, `э`},
			{`\Au(?![euioay-])`, `ю`},
			{`\Ath`, `с`}, {`\Aeu`, `ев`},
			{`\Ax(?![euioay])`, `икс`},
			{`ie\Z`, `и`}, {`ies\Z`, `ис`},
			{`th\Z`, `с`}, {`ue\Z`, `ю`},
			{`ey\Z`, `и`}, {`ai\Z`, `ай`},
			{`au`, `о`},
			{`[ae]i|ey`, `ей`},
			{`(?<=[rdgkzb])h(?!\Z)`, ``},
			{`(?<=\w{3})e\Z`, ``},
			{`qu`, `кв`}, {`ie`, `и`}, {`ue`, `ью`}, {`eu`, `ью`},
			{`ck`, `к`}, {`wh`, `в`}, {`ch`, `ч`}, {`th`, `з`}, {`sh`, `ш`}, {`ph`, `ф`},
			{`ee`, `и`}, {`oar`, `ор`}, {`oo`, `у`},
			{`ya`, `я`}, {`ye`, `е`}, {`yu`, `ю`}, {`yi`, `и`}, {`\Ayo`, `йо`}, {`ea`, `и`},
			{`b`, `б`}, {`c`, `к`}, {`d`, `д`}, {`f`, `ф`}, {`g`, `г`},
			{`h`, `х`}, {`k`, `к`}, {`l`, `л`}, {`m`, `м`}, {`n`, `н`},
			{`p`, `п`}, {`q`, `к`}, {`r`, `р`}, {`s`, `с`}, {`t`, `т`},
			{`v`, `в`}, {`w`, `в`}, {`x`, `кс`}, {`z`, `з`}, {`a`, `а`},
			{`e`, `е`}, {`i+`, `и`}, {`o`, `о`}, {`u+`, `у`}, {`y+`, `и`}, {`j`, `дж`},
		},
		//FRA

		{
			{`eaux\Z`, `о`}, {`beaut`, `бьют`}, {`eau`, `о`}, {`ogne\Z`, `он`},
			{`gnie`, `йн`}, {`agne`, `ейн`}, {`ouge`, `уж`}, {`oix`, `уа`}, {`iei`, `ье`},
			{`oux`, `о`}, {`(?<=\w)qu[eéè]\Z`, `к`}, {`ch`, `ш`},

			{`\Ales\Z`, `ле`}, {`\Ac`, `к`},
			{`\A[eéè](?![euioayéèàù-])`, `э`},
			{`\Aeu`, `ев`}, {`u[eéè]u`, `е`},
			{`ieu\Z`, `ью`},
			{`u[eéè]\Z`, `ью`}, {`gi[eéè]\Z`, `ж`},
			{`nc[eéè]\Z`, `нс`},
			{`g\Z`, `ж`}, {`z\Z`, `ц`}, {`y\Z`, `и`},
			{`tion(?=\w?\Z)`, `шн`},
			{`g(?=[ieyéè])`, `ж`},
			{`ai(?=[^euioayéèàù-]{2})`, `е`},
			{`l[eéè]`, `ле`}, {`l[uù]`, `лю`},
			{`(?<=[euioayéèàù])x{?=[euioayéèàù]\w+}`, `кз`},
			{`(?<=[euioayéèàù])s(?=[euioayéèàù])`, `з`},
			{`(?<=\w{3})[e]\Z`, ``},
			{`(?<=[rdgkzb])h(?!\Z)`, ``},

			{`ph`, `ф`}, {`qu`, `кв`}, {`sc`, `ск`}, {`cs`, `кс`}, {`th`, `т`},
			{`oi`, `уа`}, {`ou`, `у`}, {`ay`, `ей`}, {`ie`, `ье{`},

			{`b`, `б`}, {`c`, `к`}, {`d`, `д`}, {`f`, `ф`}, {`g`, `г`},
			{`h`, `х`}, {`j`, `ж`}, {`k`, `к`}, {`l`, `л`}, {`m`, `м`},
			{`n`, `н`}, {`p`, `п`}, {`q`, `к`}, {`r`, `р`}, {`s`, `с`},
			{`t`, `т`}, {`v`, `в`}, {`w`, `в`}, {`x`, `кс`}, {`z`, `з`},
			{`a`, `а`}, {`e`, `е`}, {`i`, `и`}, {`o`, `о`}, {`u`, `у`},
			{`y`, `и`}, {`é`, `е`}, {`è`, `е`}, {`à`, `а`}, {`ù`, `у`},
			{`ê`, `е`}, {`â`, `а`}, {`ô`, `о`}, {`î`, `и`}, {`û`, `у`},
			{`ë`, `е`}, {`ï`, `и`}, {`ü`, `у`}, {`ÿ`, `и`}, {`ç`, `с`},
		},

		///ITALY

		{
			{`cch`, `чч`}, {`zz`, `цц`}, {`lum`, `люм`},

			{`\Az`, `з`},
			{`(?<!l)l(?![leuioay-])`, `ль`}, {`l\Z`, `ль`},
			{`(?<![euioay-])z`, `ц`},
			{`cc(?=[ei])`, `чч`}, {`(?<!s)c(?=[ei])`, `ч`},
			{`gg?(?=[ei])`, `дж`},
			{`(?<![euioay-])eu(?![euioay-])`, `ью`},
			{`\Ae(?![euioay-])`, `э`},
			{`(?<=[euioa])i(?=\w[euioa])`, `й`},
			{`ue\Z`, `ью`},

			{`cc`, `цц`}, {`ch`, `к`}, {`qu`, `кв`}, {`sh`, `ш`}, {`ts`, `ц`},

			{`b`, `б`}, {`c`, `к`}, {`d`, `д`}, {`f`, `ф`},
			{`g`, `г`}, {`h`, `х`}, {`j`, `ж`}, {`k`, `к`},
			{`l`, `л`}, {`m`, `м`}, {`n`, `н`}, {`p`, `п`},
			{`q`, `к`}, {`r`, `р`}, {`s`, `с`}, {`t`, `т`},
			{`v`, `в`}, {`w`, `в`}, {`x`, `кс`}, {`z`, `з`},
			{`a+`, `а`}, {`e+`, `е`}, {`i+`, `и`}, {`o+`, `о`},
			{`u+`, `у`}, {`y+`, `и`},
		},

		//GERMAN

		{

			{`tsch`, `ч`}, {`zsch`, `ч`}, {`sch`, `ш`}, {`chs`, `кс`}, {`ss`, `сс`},
			{`sch`, `ч`}, {`ch`, `х`}, {`ph`, `ф`}, {`rh`, `р`}, {`th`, `т`}, {`gk`, `г`}, {`tz `, `ц`},

			{`s(?=[euioay])`, `з`},
			{`(?<![eiiouay])z(?![euioay])`, `ц`},
			{`(?<=[euoiay])v(?=[euioay])`, `в`},
			{`(?<![euioay])ä`, `е`},
			{`\Ae(?![euioay-])`, `э`},
			{`(?<!\Ai)st`, `шт`},

			{`s(?=[pt])`, `ш`},
			{`l(?=[euioay])`, `л`},
			{`l(?=[bcdfgkmnpqrstvwxz$])`, `ль`},
			{`ll$`, `лль`},
			{`(?=[bcdfgkmnpqrstvwxz])ll(?=[bcdfgkmnpqrstvwxz])`, `лль`},

			{`t?z\Z`, `ц`}, {`\Ach`, `к`},

			{`ch`, `х`}, {`tz`, `ц`}, {`sp`, `шп`},
			{`ck`, `к`}, {`ph`, `ф`}, {`sh`, `ш`},
			{`eh`, `е`}, {`je`, `е`}, {`ju`, `ю`},
			{`ja`, `я`}, {`qu`, `кв`}, {`ei`, `ай`},
			{`ie`, `и`}, {`eu`, `ой`},

			{`b`, `б`}, {`c`, `к`}, {`d`, `д`}, {`f`, `ф`}, {`g`, `г`}, {`h`, `х`},
			{`j`, `й`}, {`k`, `к`}, {`l`, `л`}, {`m`, `м`}, {`n`, `н`}, {`p`, `п`},
			{`q`, `к`}, {`r`, `р`}, {`s`, `с`}, {`t`, `т`}, {`v`, `ф`}, {`w`, `в`},
			{`x`, `кс`}, {`z`, `з`}, {`a+`, `а`}, {`e+`, `е`}, {`i+`, `и`},
			{`u+`, `у`}, {`o+`, `о`}, {`y+`, `и`},
		},

		//LATIN
		{

			{`\Aa\Z`, `эй`}, {`\Ab\Z`, `би`}, {`\Ac\Z`, `си`}, {`\Ad\Z`, `ди`}, {`\Af\Z`, `эф`}, {`\Ae\Z`, `и`}, {`\Ao\Z`, `оу`}, {`\Ac\Z`, `си`}, {`\Ax\Z`, `экс`}, {`\Aу\Z`, `уай`},
			{`\Ag\Z`, `джи`}, {`\Ah\Z`, `эйч`}, {`\Aj\Z`, `джей`}, {`\Ak\Z`, `кей`},
			{`\Al\Z`, `эль`}, {`\Am\Z`, `эм`}, {`\An\Z`, `эн`}, {`\Ap\Z`, `пи`},
			{`\Aq\Z`, `кью`}, {`\Ar\Z`, `эр`}, {`\As\Z`, `эс`}, {`\At\Z`, `ти`},
			{`\Av\Z`, `ви`}, {`\Aw\Z`, `даблю`}, {`\Ax\Z`, `икс`}, {`\Az\Z`, `зед`},
			{`\Aa\Z`, `эй`}, {`\Ao\Z`, `оу`}, {`\Ai\Z`, `ай`}, {`\Au\Z`, `ю`},
			{`\Ae\Z`, `и`}, {`\Ay\Z`, `вай`},

			{`sch`, `ш`}, {`ya`, `я`}, {`\Aand\Z`, `энд`},

			{`(?<=[euioay])s(?=[euioay])`, `з`},
			{`l(?![euioaylk])`, `ль`},
			{`\Ay(?=[euioa])`, `й`},
			{`\Ae(?![euioay-])`, `э`},
			{`(?<=[euioa])y\Z`, `й`},
			{`(?<![euioa])y\Z`, `и`},
			{`(?<=\w{3})e\Z{`, ``},

			{`qu`, `кв`}, {`ch`, `ч`}, {`sh`, `ш`}, {`ck`, `к`}, {`th`, `т`},
			{`ju`, `ю`}, {`ja`, `я`}, {`je`, `е`}, {`jo`, `е`},
			{`ph`, `ф`}, {`sc`, `ск`}, {`you`, `ю`},

			{`b`, `б`}, {`c`, `к`}, {`d`, `д`}, {`f`, `ф`}, {`g`, `г`}, {`h`, `х`},
			{`j`, `й`}, {`k`, `к`}, {`l`, `л`}, {`m`, `м`}, {`n`, `н`}, {`p`, `п`},
			{`q`, `к`}, {`r`, `р`}, {`s`, `с`}, {`t`, `т`}, {`v`, `в`}, {`w`, `в`},
			{`x`, `кс`}, {`z`, `з`}, {`a+`, `а`}, {`e+`, `е`}, {`i+`, `и`},
			{`u+`, `у`}, {`o+`, `о`}, {`y+`, `и`},
		},
		//JAPANESE

		{

			{`sh`, `ш`}, {`ts`, `ц`}, {`ya`, `я`}, {`yo`, `е`}, {`yu`, `ю`},
			{`aa`, `а`}, {`ee`, `е`}, {`uu`, `у`}, {`ii`, `и`}, {`oo`, `о`},

			{`b`, `б`}, {`d`, `д`}, {`f`, `ф`}, {`g`, `г`},
			{`h`, `х`}, {`j`, `дж`}, {`k`, `к`}, {`l`, `л`},
			{`m`, `м`}, {`n`, `н`}, {`p`, `п`}, {`r`, `р`},
			{`s`, `с`}, {`t`, `т`}, {`w`, `в`}, {`z`, `з`},
			{`a`, `а`}, {`e`, `е`}, {`i`, `и`}, {`u`, `у`},

			{"e", "е"}, {"o", "о"}, {"b", "б"}, {"c", "к"}, {"d", "д"}, {"f", "ф"}, {"g", "г"}, {"h", "х"}, {"j", "й"}, {"k", "к"}, {"l", "л"}, {"m", "м"}, {"n", "н"}, {"p", "п"}, {"q", "к"}, {"r", "р"}, {"s", "с"}, {"t", "т"}, {"v", "в"}, {"w", "в"}, {"x", "кс"}, {"z", "з"}, {"a", "а"}, {"e", "е"}, {"i", "и"}, {"u", "у"}, {"o", "о"}, {"y", "и"}, {"ß", "сс"}, {"ö", "о"}, {"ä", "а"}, {"ü", "у"}, {"é", "е"}, {"è", "е"}, {"à", "а"}, {"ù", "у"}, {"ê", "е"}, {"â", "а"}, {"ô", "о"}, {"î", "и"}, {"û", "у"}, {"ë", "е"}, {"ï", "и"}, {"ÿ", "и"}, {"ç", "с"},
		},

		/////////////////////////

	}
	newTranscriber.MultilangRulesRegexps = make([][]regexp.Regexp, len(newTranscriber.MultilangRules))
	for keyLanguage, oneLangRules := range newTranscriber.MultilangRules {

		newTranscriber.MultilangRulesRegexps[keyLanguage] = make([]regexp.Regexp, len(newTranscriber.MultilangRules[keyLanguage]))
		for keyRule, twoRules := range oneLangRules {
			newTranscriber.MultilangRulesRegexps[keyLanguage][keyRule] = regexp.MustCompile(twoRules[0], 0)
		}
	}

	newTranscriber.AbbrRules = [][]string{
		{`b`, `би`}, {`c`, `си`}, {`d`, `ди`}, {`f`, `эф`},
		{`g`, `джи`}, {`h`, `эйч`}, {`j`, `джей`}, {`k`, `кей`},
		{`l`, `эль`}, {`m`, `эм`}, {`n`, `эн`}, {`p`, `пи`},
		{`q`, `кью`}, {`r`, `эр`}, {`s`, `эс`}, {`t`, `ти`},
		{`v`, `ви`}, {`w`, `даблю`}, {`x`, `икс`}, {`z`, `зед`},
		{`a`, `эй`}, {`o`, `оу`}, {`i`, `ай`}, {`u`, `ю`},
		{`e`, `и`}, {`y`, `вай`},
	}

	newTranscriber.AbbrRulesRegexps = make([]regexp.Regexp, len(newTranscriber.AbbrRules))
	for keyRule, twoRules := range newTranscriber.AbbrRules {
		newTranscriber.AbbrRulesRegexps[keyRule] = regexp.MustCompile(twoRules[0], 0)
	}

	newTranscriber.DoubleRules = [][]string{
		{`а+`, `а`},
		{`б+`, `б`},
		{`в+`, `в`},
		{`г+`, `г`},
		{`д+`, `д`},
		{`е+`, `е`},
		{`ё+`, `ё`},
		{`ж+`, `ж`},
		{`з+`, `з`},
		{`и+`, `и`},
		{`й+`, `й`},
		{`к+`, `к`},
		{`л+`, `л`},
		{`м+`, `м`},
		{`н+`, `н`},
		{`о+`, `о`},
		{`п+`, `п`},
		{`р+`, `р`},
		{`с+`, `с`},
		{`т+`, `т`},
		{`у+`, `у`},
		{`ф+`, `ф`},
		{`х+`, `х`},
		{`ц+`, `ц`},
		{`ч+`, `ч`},
		{`ш+`, `ш`},
		{`щ+`, `щ`},
		{`ъ+`, `ъ`},
		{`ы+`, `ы`},
		{`ь+`, `ь`},
		{`э+`, `э`},
		{`ю+`, `ю`},
		{`я+`, `я`},
	}

	newTranscriber.DoubleRulesRegexps = make([]*nativeregexp.Regexp, len(newTranscriber.DoubleRules))
	for keyRule, twoRules := range newTranscriber.DoubleRules {
		newTranscriber.DoubleRulesRegexps[keyRule] = nativeregexp.MustCompile(twoRules[0])
	}

	return newTranscriber

	//Load Dict
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)

	f, err := os.Open(exPath + "/dict.csv")
	if err != nil {
		log.Fatal("Unable to read input file "+exPath+"/dict.csv", err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+exPath+"/dict.csv", err)
	}

	for _, row := range records {
		if len(row) != 2 {
			continue
		}
		if _, isset := newTranscriber.Dict[row[0]]; !isset {
			newTranscriber.Dict[row[0]] = []string{row[1]}
		} else {
			newTranscriber.Dict[row[0]] = append(newTranscriber.Dict[row[0]], row[1])
		}
	}

	return newTranscriber
}

func (obj *Transcriber) Transcribe(text string) []string {
	result := []string{}

	//1. Got correct word from dictionary

	result = append(result, obj.Dict[text]...)

	/*
		if len(result) > 0 {
			//return result
		}
	*/
	for keyLanguage, oneLangRules := range obj.MultilangRules {
		newWordMultilang := text

		for keyRule, twoRules := range oneLangRules {

			newWordMultilang, _ = obj.MultilangRulesRegexps[keyLanguage][keyRule].ReplaceAllString(newWordMultilang, twoRules[1], 0)

		}

		result = append(result, newWordMultilang)
	}

	//////////////////////////
	if len([]rune(text)) <= 4 {

		newWordAbbr := text

		for keyRule, twoRules := range obj.AbbrRules {

			newWordAbbr, _ = obj.AbbrRulesRegexps[keyRule].ReplaceAllString(newWordAbbr, twoRules[1], 0)

		}

		result = append(result, newWordAbbr)
	}

	result = removeDuplicateStr(result)

	//убираем дубли букв
	for _, word := range result {
		newWordWithoutDoubles := word
		for ruleKey, twoRules := range obj.DoubleRules {

			newWordWithoutDoubles = obj.DoubleRulesRegexps[ruleKey].ReplaceAllString(newWordWithoutDoubles, twoRules[1])

		}

		if newWordWithoutDoubles != word {
			result = append(result, newWordWithoutDoubles)
		}
	}
	return removeDuplicateStr(result)
}
