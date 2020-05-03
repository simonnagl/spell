// Package alphabet implements word-spelling alphabets. Clients should not use this internal package, used by github.com/simonnagl/spell/cmd/spell.
package alphabet

import (
	"fmt"
	"golang.org/x/text/language"
	"golang.org/x/text/language/display"
	"strings"
	"unicode"
)

// SpellingAlphabet represents a word-spelling alphabet.
//
// SpellingAlphabet is a set of words used to pronounce the letters of an alphabet in oral communication.
type SpellingAlphabet struct {
	// BCP 47 language tag, describing where this SpellingAlphabet is used.
	lang language.Tag
	// Additional names of organisations or standards, defining or using this SpellingAlphabet.
	names []string
	// Map lower case keys to their phonetic form.
	m map[string]string
	// Language specific case mappings. Can be nil.
	c *unicode.SpecialCase
}

// Names of organisations or standards, defining or using this SpellingAlphabet.
func (sa SpellingAlphabet) Names() []string {
	return sa.names
}

// LangTag returns a BCP 47 tag, describing where SpellingAlphabet is used.
func (sa SpellingAlphabet) LangTag() string {
	return sa.lang.String()
}

// LangEnglishName returns the full description in English, which language uses SpellingAlphabet.
func (sa SpellingAlphabet) LangEnglishName() string {
	return display.English.Tags().Name(sa.lang)
}

// LangEnglishName returns the full description which language uses SpellingAlphabet, in this language.
func (sa SpellingAlphabet) LangSelfName() string {
	return display.Self.Name(sa.lang)
}

// Spell generates the text to speak for spelling text.
func (sa SpellingAlphabet) Spell(text string) string {
	maxKeyLen := sa.maxMKeyLen()

	var sb strings.Builder
	for i := 0; i < len(text); {

		matchGroup := text[i:]
		if len(matchGroup) > maxKeyLen {
			matchGroup = matchGroup[:maxKeyLen]
		}

		key, value := sa.spellFirstMatch(matchGroup)
		i += len(key)
		sb.WriteString(value)
		if i != len(text) {
			sb.WriteRune(' ')
		}
	}
	return sb.String()
}

func (sa SpellingAlphabet) maxMKeyLen() int {
	var maxKeyLen int
	for key := range sa.m {
		if maxKeyLen < len(key) {
			maxKeyLen = len(key)
		}
	}
	return maxKeyLen
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

// Lookup returns the best matching SpellingAlphabet of All together with a confidence score.
//
// First, Lookup searches for a SpellingAlphabet with a matching name.
// Second, Lookup tries to interpret key as a BCP 47 language tag and finds the best match for the SpellingAlphabets lang.
// golang.org/x/text/language is used for finding the best match.
// If there is no match, 'en' is used as the default SpellingAlphabet.
func Lookup(key string) (SpellingAlphabet, Exactness) {
	if a, ok := lookupName(key); ok {
		return a, Exact
	}
	return lookupLang(key)
}

func lookupName(key string) (alphabet SpellingAlphabet, ok bool) {
	for _, alphabet := range All {
		for _, name := range alphabet.Names() {
			if key == name {
				return alphabet, true
			}
		}
	}
	return SpellingAlphabet{}, false
}

func lookupLang(lang string) (SpellingAlphabet, Exactness) {

	tags := make([]language.Tag, 0, len(All))
	for _, alphabet := range All {
		tags = append(tags, alphabet.lang)
	}

	matcher := language.NewMatcher(tags)

	tag, err := language.Parse(lang)
	if err != nil {
		return All[0], Default
	}
	_, i, c := matcher.Match(tag)

	return All[i], FromLangConfidence(c)
}

// All SpellingAlphabet.
var All = []SpellingAlphabet{
	English,
	BritishEnglish,
	French,
	Dutch,
	German,
	AustrianGerman,
	SwissHighGerman,
	Italian,
	Spanish,
	Turkish,
	Norwegian,
	Swedish,
	Finnish,
	Danish,
	Czech,
	EuropeanPortuguese,
	BrazilianPortuguese,
	Romanian,
	Slovenian,
}

var (
	English = SpellingAlphabet{
		lang:  language.English,
		names: []string{"ICAO", "NATO"},
		m: map[string]string{
			"a":  "Alfa",
			"b":  "Bravo",
			"c":  "Charlie",
			"d":  "Delta",
			"e":  "Echo",
			"f":  "Foxtrot",
			"g":  "Golf",
			"h":  "Hotel",
			"i":  "India",
			"j":  "Juliett",
			"k":  "Kilo",
			"l":  "Lima",
			"m":  "Mike",
			"n":  "November",
			"o":  "Oscar",
			"p":  "Papa",
			"q":  "Quebec",
			"r":  "Romeo",
			"s":  "Sierra",
			"t":  "Tango",
			"u":  "Uniform",
			"v":  "Victor",
			"w":  "Whiskey",
			"x":  "X-ray",
			"y":  "Yankee",
			"z":  "Zulu",
			"0":  "Zero",
			"1":  "One",
			"2":  "Two",
			"3":  "Three",
			"4":  "Four",
			"5":  "Five",
			"6":  "Six",
			"7":  "Seven",
			"8":  "Eight",
			"9":  "Nine",
			" ":  "Space",
			".":  "Dot",
			",":  "Comma",
			";":  "Semicolon",
			":":  "Colon",
			"?":  "Question Mark",
			"!":  "Exclamation Mark",
			"@":  "At Sign",
			"&":  "Ampersand",
			"\"": "Double Quotation Mark",
			"'":  "Singel Quotation Mark",
			"-":  "Dash",
			"/":  "Forward Slash",
			"\\": "Backslash",
			"(":  "Left Round Bracket",
			")":  "Right Round Bracket",
			"[":  "Left Square Bracket",
			"]":  "Right Square Bracket",
			"{":  "Left Curly Bracket",
			"}":  "Right Curly Bracket",
			"<":  "Left Angle Bracket",
			">":  "Right Angle Bracket",
			"|":  "Vertical Bar",
			"°":  "Degree Symbol",
			"*":  "Asterisk",
			"+":  "Plus Sign",
			"=":  "Equal Sign",
			"#":  "Number Sign",
			"§":  "Section Sign",
			"$":  "Dollar Sign",
			"€":  "Euro Sign",
			"~":  "Tilde",
			"_":  "Underscore",
			"%":  "Percent Sign",
			"^":  "Caret",
		},
	}
	BritishEnglish = SpellingAlphabet{
		lang: language.BritishEnglish,
		m: map[string]string{
			"a": "Alfred",
			"b": "Benjamin",
			"c": "Charles",
			"d": "David",
			"e": "Edward",
			"f": "Frederick",
			"g": "George",
			"h": "Harry",
			"i": "Isaac",
			"j": "Jack",
			"k": "King",
			"l": "London",
			"m": "Mary",
			"n": "Nellie",
			"o": "Oliver",
			"p": "Peter",
			"q": "Queen",
			"r": "Robert",
			"s": "Samuel",
			"t": "Tommy",
			"u": "Uncle",
			"v": "Victor",
			"w": "William",
			"x": "X-ray",
			"y": "Yellow",
			"z": "Zebra",
		},
	}
	French = SpellingAlphabet{
		lang: language.French,
		m: map[string]string{
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
		},
	}
	Dutch = SpellingAlphabet{
		lang: language.Dutch,
		m: map[string]string{
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
	}
	German = SpellingAlphabet{
		lang:  language.MustParse("de-DE"),
		names: []string{"DIN 5009"},
		m: map[string]string{
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
			"0":   "Null",
			"1":   "Eins",
			"2":   "Zwo",
			"3":   "Drei",
			"4":   "Vier",
			"5":   "Fünf",
			"6":   "Sechs",
			"7":   "Sieben",
			"8":   "Acht",
			"9":   "Neun",
			" ":   "Leerzeichen",
			".":   "Punkt",
			",":   "Komma",
			";":   "Semikolon",
			":":   "Doppelpunkt",
			"?":   "Fragezeichen",
			"!":   "Ausrufezeichen",
			"@":   "Klammeraffe",
			"&":   "kaufmännisches Und",
			"\"":  "Anführungszeichen",
			"'":   "Apostroph",
			"-":   "Bindestrich",
			"/":   "Schrägstrich",
			"\\":  "Umgekehrter Schrägstrich",
			"(":   "Runde Klammer links",
			")":   "Runde Klammer rechts",
			"[":   "Eckige Klammer links",
			"]":   "Eckige Klammer rechts",
			"{":   "Geschweifte Klammer links",
			"}":   "Geschweifte Klammer rechts",
			"<":   "Spitze Klammer links",
			">":   "Spitze Klammer rechts",
			"|":   "Senkrechter Strich",
			"°":   "Gradzeichen",
			"*":   "Asterisk",
			"+":   "Pluszeichen",
			"=":   "Gleichheitszeichen",
			"#":   "Rautenzeichen",
			"§":   "Paragraphenzeichen",
			"$":   "Dollarzeichen",
			"€":   "Eurozeichen",
			"~":   "Tilde",
			"_":   "Unterstrich",
			"%":   "Prozentzeichen",
			"^":   "Zirkumflex",
		},
	}
	AustrianGerman = SpellingAlphabet{
		lang:  language.MustParse("de-AT"),
		names: []string{"ÖNORM A 1081"},
		m: map[string]string{
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
			"0":   "Null",
			"1":   "Eins",
			"2":   "Zwo",
			"3":   "Drei",
			"4":   "Vier",
			"5":   "Fünf",
			"6":   "Sechs",
			"7":   "Sieben",
			"8":   "Acht",
			"9":   "Neun",
			" ":   "Leerzeichen",
			".":   "Punkt",
			",":   "Komma",
			";":   "Semikolon",
			":":   "Doppelpunkt",
			"?":   "Fragezeichen",
			"!":   "Ausrufezeichen",
			"@":   "Klammeraffe",
			"&":   "kaufmännisches Und",
			"\"":  "Anführungszeichen",
			"'":   "Apostroph",
			"-":   "Bindestrich",
			"/":   "Schrägstrich",
			"\\":  "Umgekehrter Schrägstrich",
			"(":   "Runde Klammer links",
			")":   "Runde Klammer rechts",
			"[":   "Eckige Klammer links",
			"]":   "Eckige Klammer rechts",
			"{":   "Geschweifte Klammer links",
			"}":   "Geschweifte Klammer rechts",
			"<":   "Spitze Klammer links",
			">":   "Spitze Klammer rechts",
			"|":   "Senkrechter Strich",
			"°":   "Gradzeichen",
			"*":   "Asterisk",
			"+":   "Pluszeichen",
			"=":   "Gleichheitszeichen",
			"#":   "Rautenzeichen",
			"§":   "Paragraphenzeichen",
			"$":   "Dollarzeichen",
			"€":   "Eurozeichen",
			"~":   "Tilde",
			"_":   "Unterstrich",
			"%":   "Prozentzeichen",
			"^":   "Zirkumflex",
		},
	}
	SwissHighGerman = SpellingAlphabet{
		lang: language.MustParse("de-CH"),
		m: map[string]string{
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
			"0":  "Null",
			"1":  "Eins",
			"2":  "Zwo",
			"3":  "Drei",
			"4":  "Vier",
			"5":  "Fünf",
			"6":  "Sechs",
			"7":  "Sieben",
			"8":  "Acht",
			"9":  "Neun",
			" ":  "Leerzeichen",
			".":  "Punkt",
			",":  "Komma",
			";":  "Semikolon",
			":":  "Doppelpunkt",
			"?":  "Fragezeichen",
			"!":  "Ausrufezeichen",
			"@":  "Klammeraffe",
			"&":  "kaufmännisches Und",
			"\"": "Anführungszeichen",
			"'":  "Apostroph",
			"-":  "Bindestrich",
			"/":  "Schrägstrich",
			"\\": "Umgekehrter Schrägstrich",
			"(":  "Runde Klammer links",
			")":  "Runde Klammer rechts",
			"[":  "Eckige Klammer links",
			"]":  "Eckige Klammer rechts",
			"{":  "Geschweifte Klammer links",
			"}":  "Geschweifte Klammer rechts",
			"<":  "Spitze Klammer links",
			">":  "Spitze Klammer rechts",
			"|":  "Senkrechter Strich",
			"°":  "Gradzeichen",
			"*":  "Asterisk",
			"+":  "Pluszeichen",
			"=":  "Gleichheitszeichen",
			"#":  "Rautenzeichen",
			"§":  "Paragraphenzeichen",
			"$":  "Dollarzeichen",
			"€":  "Eurozeichen",
			"~":  "Tilde",
			"_":  "Unterstrich",
			"%":  "Prozentzeichen",
			"^":  "Zirkumflex",
		},
	}
	Italian = SpellingAlphabet{
		lang: language.Italian,
		m: map[string]string{
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
		},
	}
	Spanish = SpellingAlphabet{
		lang: language.Spanish,
		m: map[string]string{
			"a":  "Antonio",
			"b":  "Burgos",
			"c":  "Carmen",
			"ch": "Chocolate",
			"d":  "David",
			"e":  "España",
			"f":  "Francia",
			"g":  "Granada",
			"h":  "Historia",
			"i":  "Inés",
			"j":  "José",
			"k":  "Kilo",
			"l":  "Lorenzo",
			"ll": "Llave",
			"m":  "Madrid",
			"n":  "Navidad",
			"ñ":  "Ñoño",
			"o":  "Oviedo",
			"p":  "París",
			"q":  "Queso",
			"r":  "Ramón",
			"s":  "Sábado",
			"t":  "Toledo",
			"u":  "Ulises",
			"v":  "Valencia",
			"w":  "Washington",
			"x":  "Xilófono",
			"y":  "Yolanda",
			"z":  "Zaragoza",
		},
	}
	Turkish = SpellingAlphabet{
		lang: language.Turkish,
		m: map[string]string{
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
			"z": "Yozgat",
		},
		c: &unicode.TurkishCase,
	}
	Norwegian = SpellingAlphabet{
		lang: language.Norwegian,
		m: map[string]string{
			"a": "Anna",
			"å": "Åse",
			"æ": "Ægir",
			"b": "Bernhard",
			"c": "Caesar",
			"d": "David",
			"e": "Edith",
			"f": "Frederik",
			"g": "Gustav",
			"h": "Harald",
			"i": "Ivar",
			"j": "Johan",
			"k": "Karin",
			"l": "Ludvig",
			"m": "Martin",
			"n": "Nils",
			"o": "Olivia",
			"ø": "Østen",
			"p": "Petter",
			"q": "Quintus",
			"r": "Rikard",
			"s": "Sigrid",
			"t": "Teodor",
			"u": "Ulrik",
			"v": "Enkelt-V",
			"w": "Dobbelt-W",
			"x": "Xerxes",
			"y": "Yngling",
			"z": "Zakarias"},
	}
	Swedish = SpellingAlphabet{
		lang: language.Swedish,
		m: map[string]string{
			"a": "Adam",
			"å": "Åke",
			"ä": "Ärlig",
			"b": "Bertil",
			"c": "Cesar",
			"d": "David",
			"e": "Erik",
			"f": "Filip",
			"g": "Gustav",
			"h": "Helge",
			"i": "Ivar",
			"j": "Johan",
			"k": "Kalle",
			"l": "Ludvig",
			"m": "Martin",
			"n": "Niklas",
			"o": "Olof",
			"ö": "Östen",
			"p": "Petter",
			"q": "Qvintus",
			"r": "Rudolf",
			"s": "Sigurd",
			"t": "Tore",
			"u": "Urban",
			"ü": "Übel",
			"v": "Viktor",
			"w": "Wilhelm",
			"x": "Xerxes",
			"y": "Yngve",
			"z": "Zäta",
		},
	}
	Finnish = SpellingAlphabet{
		lang: language.Finnish,
		m: map[string]string{
			"a": "Aarne",
			"ä": "Äiti",
			"å": "Åke",
			"b": "Bertta",
			"c": "Celsius",
			"d": "Daavid",
			"e": "Eemeli",
			"f": "Faarao",
			"g": "Gideon",
			"h": "Heikki",
			"i": "Iivari",
			"j": "Jussi",
			"k": "Kalle",
			"l": "Lauri",
			"m": "Matti",
			"n": "Niilo",
			"o": "Otto",
			"ö": "Öljy",
			"p": "Paavo",
			"q": "Kuu",
			"r": "Risto",
			"s": "Sakari",
			"t": "Tyyne",
			"u": "Urho",
			"v": "Vihtori",
			"w": "Wiski",
			"x": "Äksä",
			"y": "Yrjö",
			"z": "Tseta",
		},
	}
	Danish = SpellingAlphabet{
		lang: language.Danish,
		m: map[string]string{
			"a": "Anna",
			"å": "Åse",
			"æ": "Ægir",
			"b": "Bernhard",
			"c": "Cecilie",
			"d": "David",
			"e": "Erik",
			"f": "Frederik",
			"g": "Georg",
			"h": "Hans",
			"i": "Ida",
			"j": "Johan",
			"k": "Karen",
			"l": "Ludvig",
			"m": "Mari",
			"n": "Nikolaj",
			"o": "Odin",
			"ø": "Øresund",
			"p": "Peter",
			"q": "Quintus",
			"r": "Rasmus",
			"s": "Søren",
			"t": "Theodor",
			"u": "Ulla",
			"v": "Viggo",
			"w": "William",
			"x": "Xerxes",
			"y": "Yrsa",
			"z": "Zackarias",
		},
	}
	Czech = SpellingAlphabet{
		lang: language.Czech,
		m: map[string]string{
			"a":  "Adam",
			"á":  "a s čárkou",
			"b":  "Božena",
			"c":  "Cyril",
			"č":  "Čeněk",
			"d":  "David",
			"ď":  "Ďáblice",
			"e":  "Emil",
			"é":  "e s čárkou",
			"ě":  "e s háčkem",
			"f":  "František",
			"g":  "Gustav",
			"h":  "Helena",
			"ch": "Chrudim",
			"i":  "Ivan",
			"í":  "i s čárkou",
			"j":  "Josef",
			"k":  "Karel",
			"l":  "Ludvík",
			"m":  "Marie",
			"n":  "Neruda",
			"ň":  "n s háčkem",
			"o":  "Oto",
			"ó":  "o s čárkou",
			"p":  "Petr",
			"q":  "Quido",
			"r":  "Rudolf",
			"ř":  "Řehoř",
			"s":  "Svatopluk",
			"š":  "Šimon",
			"t":  "Tomáš",
			"ť":  "Těšnov",
			"u":  "Urban",
			"ú":  "u s čárkou",
			"ů":  "u s kroužkem",
			"v":  "Václav",
			"w":  "Dvojité v",
			"x":  "Xaver",
			"y":  "Ypsilon",
			"ý":  "y s čárkou",
			"z":  "Zuzana",
			"ž":  "Žofie",
		},
	}
	EuropeanPortuguese = SpellingAlphabet{
		lang: language.EuropeanPortuguese,
		m: map[string]string{
			"a": "Aveiro",
			"b": "Braga",
			"c": "Coimbra",
			"d": "Dafundo",
			"e": "Évora",
			"f": "Faro",
			"g": "Guarda",
			"h": "Horta",
			"i": "Itália",
			"j": "José",
			"k": "Kodak",
			"l": "Lisboa",
			"m": "Maria",
			"n": "Nazaré",
			"o": "Ovar",
			"p": "Porto",
			"q": "Queluz",
			"r": "Rossio",
			"s": "Setúbal",
			"t": "Tavira",
			"u": "Unidade",
			"v": "Vidago",
			"w": "Waldemar",
			"x": "Xavier",
			"y": "York",
			"z": "Zulmira",
		},
	}
	BrazilianPortuguese = SpellingAlphabet{
		lang: language.BrazilianPortuguese,
		m: map[string]string{
			"a": "Amor",
			"b": "Bandeira",
			"c": "Cobra",
			"d": "Dado",
			"e": "Estrela",
			"f": "Feira",
			"g": "Goiaba",
			"h": "Hotel",
			"i": "Índio",
			"j": "José",
			"k": "Kiwi",
			"l": "Lua",
			"m": "Maria",
			"n": "Navio",
			"o": "Ouro",
			"p": "Pipa",
			"q": "Quilombo",
			"r": "Raiz",
			"s": "Saci",
			"t": "Tatu",
			"u": "Uva",
			"v": "Vitória",
			"w": "Wilson",
			"x": "Xadrez",
			"y": "Yolanda",
			"z": "Zebra",
		},
	}
	Romanian = SpellingAlphabet{
		lang: language.Romanian,
		m: map[string]string{
			"a": "Ana",
			"b": "Barbu",
			"c": "Constantin",
			"d": "Dumitru",
			"e": "Elena",
			"f": "Florea",
			"g": "Gheorghe",
			"h": "Haralambie",
			"i": "Ion",
			"j": "Jean",
			"k": "Kilogram",
			"l": "Lazăr",
			"m": "Maria",
			"n": "Nicolae",
			"o": "Olga",
			"p": "Petre",
			"q": "Qu (Chiu)",
			"r": "Radu",
			"s": "Sandu",
			"t": "Tudor",
			"u": "Udrea",
			"v": "Vasile",
			"w": "dublu v",
			"x": "Xenia",
			"y": "I grec",
			"z": "Zahăr",
		},
	}
	Slovenian = SpellingAlphabet{
		lang: language.Slovenian,
		m: map[string]string{
			"a": "Ankaran",
			"b": "Bled",
			"c": "Celje",
			"č": "Čatež",
			"d": "Drava",
			"e": "Evropa",
			"f": "Fala",
			"g": "Gorica",
			"h": "Hrastnik",
			"i": "Izola",
			"j": "Jadran",
			"k": "Kamnik",
			"l": "Ljubljana",
			"m": "Maribor",
			"n": "Nanos",
			"o": "Ormož",
			"p": "Piran",
			"q": "Queen",
			"r": "Ravne",
			"s": "Soča",
			"š": "Šmarje",
			"t": "Triglav",
			"u": "Unec",
			"v": "Velenje",
			"w": "Dvojni v",
			"x": "Iks",
			"y": "Ipsilon",
			"z": "Zalog",
			"ž": "Žalec",
		},
	}
)
