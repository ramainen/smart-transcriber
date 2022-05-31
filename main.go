package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"

	//"regexp"
	regexp "github.com/scorpionknifes/go-pcre"
)

type Transpiler struct {
	Dict map[string][]string
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
func NewTranspiler() Transpiler {
	transpiler := Transpiler{}
	transpiler.Dict = map[string][]string{}
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
		if _, isset := transpiler.Dict[row[0]]; !isset {
			transpiler.Dict[row[0]] = []string{row[1]}
		} else {
			transpiler.Dict[row[0]] = append(transpiler.Dict[row[0]], row[1])
		}
	}

	return transpiler
}

func (obj *Transpiler) Transpile(text string) []string {
	result := []string{}

	//1. Got correct word from dictionary
	for _, word := range obj.Dict[text] {
		result = append(result, word)
	}
	if len(result) > 0 {
		//return result
	}

	simpeRules := [][]string{
		{"sch", "ш"}, {"ya", "я"}, {`\Aand\Z`, "энд"}, {`\Ay(?=[euioa])`, "й"},
		{`\Ae(?![euioay-])`, "э"}, {`(?<=[euioa])y\Z`, "й"}, {`(?<![euioa])y\Z`, "и"}, {"(?<=[euioa])c(?=[euioay-])", "с"},
		{"a", "а"}, {"e", "е"}, {"o", "о"}, {"b", "б"}, {"c", "к"}, {"d", "д"}, {"f", "ф"}, {"g", "г"}, {"h", "х"}, {"j", "й"}, {"k", "к"}, {"l", "л"}, {"m", "м"}, {"n", "н"}, {"p", "п"}, {"q", "к"}, {"r", "р"}, {"s", "с"}, {"t", "т"}, {"v", "в"}, {"w", "в"}, {"x", "кс"}, {"z", "з"}, {"a", "а"}, {"e", "е"}, {"i", "и"}, {"u", "у"}, {"o", "о"}, {"y", "и"}, {"ß", "сс"}, {"ö", "о"}, {"ä", "а"}, {"ü", "у"}, {"é", "е"}, {"è", "е"}, {"à", "а"}, {"ù", "у"}, {"ê", "е"}, {"â", "а"}, {"ô", "о"}, {"î", "и"}, {"û", "у"}, {"ë", "е"}, {"ï", "и"}, {"ÿ", "и"}, {"ç", "с"},
	}

	newWord := text

	for _, twoRules := range simpeRules {
		re := regexp.MustCompile(twoRules[0], 0)
		newWord, _ = re.ReplaceAllString(newWord, twoRules[1], 0)

	}
	result = append(result, newWord)

	englishRules := [][]string{
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
	}

	newWordEnglish := text

	for _, twoRules := range englishRules {
		re := regexp.MustCompile(twoRules[0], 0)
		newWordEnglish, _ = re.ReplaceAllString(newWordEnglish, twoRules[1], 0)

	}

	result = append(result, newWordEnglish)

	//Наивный алгорим
	/*
		'sch':'ш', 'ya':'я', '\Aand\Z':'энд'

		'\Ay(?=[euioa])':'й',
		'\Ae(?![euioay-])':'э',
		'(?<=[euioa])y\Z':'й',
		'(?<![euioa])y\Z':'и',
		'(?<=[euioa])c(?=[euioay-])':'с'



		'b':'б', 'c':'к', 'd':'д', 'f':'ф', 'g':'г', 'h':'х',
		'j':'й', 'k':'к', 'l':'л', 'm':'м', 'n':'н', 'p':'п',
		'q':'к', 'r':'р', 's':'с', 't':'т', 'v':'в', 'w':'в',
		'x':'х', 'z':'з', 'a':'а',  'e':'е',  'i':'и',
		'u':'у', 'o':'о',  'y':'и',
		'ß':'сс', 'ö':'о', 'ä':'а', 'ü':'у', 'é':'е', 'è':'е',
		'à':'а', 'ù':'у', 'ê':'е', 'â':'а', 'ô':'о', 'î':'и',
		'û':'у', 'ë':'е', 'ï':'и', 'ÿ':'и', 'ç':'с'

	*/

	return removeDuplicateStr(result)

}
func main() {

	trans := NewTranspiler()
	fmt.Println(trans.Transpile("apple"))
	fmt.Println(trans.Transpile("makita"))
	fmt.Println(trans.Transpile("bosch"))
	fmt.Println(trans.Transpile("xiaomi"))
	fmt.Println(trans.Transpile("bergauf"))
	fmt.Println(trans.Transpile("knauf"))

	fmt.Println(trans.Transpile("megastroy"))
	fmt.Println(trans.Transpile("zelda"))
	fmt.Println(trans.Transpile("argotech"))
	fmt.Println(trans.Transpile("rossinka"))
}
