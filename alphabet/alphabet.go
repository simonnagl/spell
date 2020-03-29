package alphabet

import (
	"fmt"
	"strings"
	"unicode"
)

type SpellingAlphabet struct {
	// Map lower case keys to their phonetic form.
	m map[string]string
	// Language specific case mappings. Can be nil.
	c *unicode.SpecialCase
}

func (sa SpellingAlphabet) Spell(allLetter string) string {
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
	var lowerKey string
	if sa.c == nil {
		lowerKey = strings.ToLower(key)
	} else {
		lowerKey = strings.ToLowerSpecial(*sa.c, key)
	}
	if value, ok := sa.m[lowerKey]; ok {
		return key, value
	} else {
		if len(key) == 1 {
			return key, fmt.Sprintf("'%s'", key)
		} else {
			return sa.spellFirstMatch(key[:len(key)-1])
		}
	}
}

func ForLanguageCode(lang string) SpellingAlphabet {
	a, ok := Lang[lang]
	if ok {
		return a
	} else {
		return Lang["en"]
	}
}

var Lang = map[string]SpellingAlphabet{
	// ICAO / NATO
	"en": {m: map[string]string{
		"a": "Alfa",
		"b": "Bravo",
		"c": "Charlie",
		"d": "Delta",
		"e": "Echo",
		"f": "Foxtrot",
		"g": "Golf",
		"h": "Hotel",
		"i": "India",
		"j": "Juliett",
		"k": "Kilo",
		"l": "Lima",
		"m": "Mike",
		"n": "November",
		"o": "Oscar",
		"p": "Papa",
		"q": "Quebec",
		"r": "Romeo",
		"s": "Sierra",
		"t": "Tango",
		"u": "Uniform",
		"v": "Victor",
		"w": "Whiskey",
		"x": "X-ray",
		"y": "Yankee",
		"z": "Zulu",
	}},
	"fr": {m: map[string]string{
		"a": "Anatole",
		"b": "Berthe",
		"c": "Célestin",
		"d": "Désiré",
		"e": "Eugène",
		"f": "François",
		"g": "Gaston",
		"h": "Henri",
		"i": "Irma",
		"j": "Joseph",
		"k": "Kléber",
		"l": "Louis",
		"m": "Marcel",
		"n": "Nicolas",
		"o": "Oscar",
		"p": "Pierre",
		"q": "Quintal",
		"r": "Raoul",
		"s": "Suzanne",
		"t": "Thérèse",
		"u": "Ursule",
		"v": "Victor",
		"w": "William",
		"x": "Xavier",
		"y": "Yvonne",
		"z": "Zoé",
	}},
	"nl": {m: map[string]string{
		"a": "Anna/Anton",
		"b": "Bernard",
		"c": "Cornelis",
		"d": "Dirk",
		"e": "Eduard",
		"f": "Ferdinand",
		"g": "Gerard",
		"h": "Hendrik",
		"i": "Izaak",
		"j": "Julius",
		"k": "Karel",
		"l": "Lodewijk",
		"m": "Maria",
		"n": "Nico",
		"o": "Otto",
		"p": "Pieter",
		"q": "Quotiënt",
		"r": "Richard",
		"s": "Simon",
		"t": "Theodor",
		"u": "Utrecht",
		"v": "Victor",
		"w": "Willem",
		"x": "Xanthippe",
		"y": "Ypsilon",
		"z": "Zaandam",
	}},
	// DIN5009
	"de": {m: map[string]string{
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
	}},
	// ÖNORM A 1081
	"at": {m: map[string]string{
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
		"k":   "Konrad",
		"l":   "Ludwig",
		"m":   "Martha",
		"n":   "Nordpol",
		"o":   "Otto",
		"ö":   "Österreich",
		"p":   "Paula",
		"q":   "Quelle",
		"r":   "Richard",
		"s":   "Siegfried",
		"sch": "Schule",
		"ß":   "scharfes S",
		"t":   "Theodor",
		"u":   "Ulrich",
		"ü":   "Übel",
		"v":   "Viktor",
		"w":   "Wilhelm",
		"x":   "Xaver",
		"y":   "Ypsilon",
		"z":   "Zürich",
	}},
	"ch": {m: map[string]string{
		"a":  "Anna",
		"ä":  "Äsch",
		"b":  "Berta",
		"c":  "Cäsar",
		"ch": "Chiasso",
		"d":  "Daniel",
		"e":  "Emil",
		"f":  "Friedrich",
		"g":  "Gustav",
		"h":  "Heinrich",
		"i":  "Ida",
		"j":  "Jakob",
		"k":  "Kaiser",
		"l":  "Leopold",
		"m":  "Marie",
		"n":  "Niklaus",
		"o":  "Otto",
		"ö":  "Örlikon",
		"p":  "Peter",
		"q":  "Quasi",
		"r":  "Rosa",
		"s":  "Sophie",
		"t":  "Theodor",
		"u":  "Ulrich",
		"ü":  "Übermut",
		"v":  "Viktor",
		"w":  "Wilhelm",
		"x":  "Xaver",
		"y":  "Yverdon",
		"z":  "Zürich",
	}},
	"it": {m: map[string]string{
		"a": "Ancona",
		"b": "Bari",
		"c": "Como",
		"d": "Domodossola",
		"e": "Empoli",
		"f": "Firenze",
		"g": "Genova",
		"h": "Hotel",
		"i": "Imola",
		"j": "Juventus",
		"k": "Kilometro",
		"l": "Livorno",
		"m": "Milano",
		"n": "Napoli",
		"o": "Otranto",
		"p": "Pisa",
		"q": "Quadro",
		"r": "Roma",
		"s": "Savona",
		"t": "Torino",
		"u": "Udine",
		"v": "Venezia",
		"w": "Vu Doppia",
		"x": "Xilofono",
		"y": "Ipsilon",
		"z": "Zara",
	}},
	"tr": {m: map[string]string{
		"a": "Adana",
		"b": "Bolu",
		"c": "Ceyhan",
		"ç": "Çanakkale",
		"d": "Denizli",
		"e": "Edirne",
		"f": "Fatsa",
		"g": "Giresun",
		"ğ": "yumuşak G",
		"h": "Hatay",
		"i": "İzmir",
		"ı": "Isparta",
		"j": "jandarma",
		"k": "Kars",
		"l": "Lüleburgaz",
		"m": "Muş",
		"n": "Niğde",
		"o": "Ordu",
		"ö": "Ödemiş",
		"p": "Polatlı",
		"r": "Rize",
		"s": "Sinop",
		"ş": "Şırnak",
		"t": "Tokat",
		"u": "Uşak",
		"ü": "Ünye",
		"v": "Van",
		"w": "duble V",
		"y": "Yozgat",
		"z": "Yozgat"},
		c: &unicode.TurkishCase},
}
