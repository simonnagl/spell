package alphabet

import (
	"fmt"
	"golang.org/x/text/language"
	"strings"
	"unicode"
)

type SpellingAlphabet struct {
	// Language which uses this spelling alphabet.
	Lang language.Tag
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

	tags := make([]language.Tag, 0, len(List))
	for _, alphabet := range List {
		tags = append(tags, alphabet.Lang)
	}

	matcher := language.NewMatcher(tags)

	tag, err := language.Parse(lang)
	if err != nil {
		return List[0]
	}
	_, i, _ := matcher.Match(tag)

	return List[i]
}

// List of all supported spelling alphabets.
var List = []SpellingAlphabet{
	// ICAO / NATO
	{
		Lang: language.English,
		m: map[string]string{
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
	}, {
		Lang: language.BritishEnglish,
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
	}, {
		Lang: language.French,
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
	}, {
		Lang: language.Dutch,
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
	},
	// DIN5009
	{
		Lang: language.MustParse("de-DE"),
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
			" ":   "Leerzeichen",
		},
	},
	// ÖNORM A 1081
	{
		Lang: language.MustParse("de-AT"),
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
		},
	}, {
		Lang: language.MustParse("de-CH"),
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
		},
	}, {
		Lang: language.Italian,
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
	}, {
		Lang: language.Spanish,
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
	}, {Lang: language.Turkish,
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
	}, {
		Lang: language.Norwegian,
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
	}, {
		Lang: language.Swedish,
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
	}, {
		Lang: language.Finnish,
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
	}, {
		Lang: language.Danish,
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
	},
	{
		Lang: language.Czech,
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
	}, {
		Lang: language.EuropeanPortuguese,
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
	}, {
		Lang: language.BrazilianPortuguese,
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
	}, {
		Lang: language.Romanian,
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
	}, {
		Lang: language.Slovenian,
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
	},
}
