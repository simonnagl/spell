package alphabet

import (
	"fmt"
	"strings"
)

type SpellingAlphabet map[string]string

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
	if value, ok := sa[strings.ToLower(key)]; ok {
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
	a := Lang[lang]
	if a != nil {
		return a
	} else {
		return Lang["en"]
	}
}

var Lang = map[string]SpellingAlphabet{
	// ICAO
	"en": {
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
	},
	"fr": {
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
	}, "nl": {
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
	},
	// DIN5009
	"de": {
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
	},
}
