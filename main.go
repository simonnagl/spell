package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var commandLineArguments = strings.Join(os.Args[1:], " ")
	fmt.Println(DIN5009.spell(commandLineArguments))
}

var DIN5009 = SpellingAlphabet{
	"a":   "Anton",
	"ä":   "Ärger",
	"b":   "Berta",
	"c":   "Cäsar",
	"ch":  "Charlotte",
	"d":   "Dora",
	"e":   "Emil",
	"f":   "Friedrich",
	"g":   "Gustav",
	"h":   "Heinrich",
	"i":   "Ida",
	"j":   "Julius",
	"k":   "Kaufmann",
	"l":   "Ludwig",
	"m":   "Martha",
	"n":   "Nordpol",
	"o":   "Otto",
	"ö":   "Ökonom",
	"p":   "Paula",
	"q":   "Quelle",
	"r":   "Richard",
	"s":   "Samuel",
	"sch": "Schule",
	"ß":   "Eszett",
	"t":   "Theodor",
	"u":   "Ulrich",
	"ü":   "Übermut",
	"v":   "Viktor",
	"w":   "Wilhelm",
	"x":   "Xanthippe",
	"y":   "Ypsilon",
	"z":   "Zacharias",
	" ":   "Leerzeichen",
}

type SpellingAlphabet map[string]string

func (sa SpellingAlphabet) spell(allLetter string) string {
	var sb strings.Builder
	for i := 0; i < len(allLetter); {
		key, value := sa.spellFirstMatch(allLetter[i:])
		i += len(key)
		sb.WriteString(value)
		if i != len(allLetter) {
			sb.WriteRune(' ')
		}
	}
	return sb.String()
}

func (sa SpellingAlphabet) spellFirstMatch(key string) (string, string) {
	if value, ok := sa[strings.ToLower(key)]; ok {
		return key, value
	} else {
		if len(key) == 1 {
			return key, key
		} else {
			return sa.spellFirstMatch(key[:len(key)-1])
		}
	}
}
