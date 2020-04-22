package alphabet

import (
	"golang.org/x/text/language"
	"strings"
	"testing"
)

var alphabet = SpellingAlphabet{m: map[string]string{
	"a":   "Anton",
	"ä":   "Ärger",
	"l":   "Ludwig",
	"sch": "Schule",
	"ch":  "Charlotte",
	"t":   "Theodor",
}}

func testSpell(t *testing.T, alphabet SpellingAlphabet, allInputLetters string, expectedResult string) bool {
	return t.Run("Spell "+allInputLetters, func(t *testing.T) {
		result := alphabet.Spell(allInputLetters)
		if expectedResult != result {
			t.Errorf("'%s' should be spelled as \n'%s', but was \n'%s'", allInputLetters, expectedResult, result)
		}
	})
}

func TestSpell_Alphabet(t *testing.T) {
	testSpellAlphabet(t, alphabet)
}

func testSpellAlphabet(t *testing.T, alphabet SpellingAlphabet) {
	for inputLetter, expectedResult := range alphabet.m {
		testSpell(t, alphabet, inputLetter, expectedResult)
		var titleLetter string
		if alphabet.c == nil {
			titleLetter = strings.ToTitle(inputLetter)
		} else {
			titleLetter = strings.ToTitleSpecial(*alphabet.c, inputLetter)
		}
		if inputLetter != titleLetter {
			testSpell(t, alphabet, titleLetter, expectedResult)
		}
	}
}

func TestSpell_Lang(t *testing.T) {
	for _, a := range All {
		t.Run(a.Lang.String(), func(t *testing.T) {
			testSpellAlphabet(t, a)
		})
	}
}

func TestSpell_Words(t *testing.T) {
	testSpell(t, alphabet, "aä", "Anton Ärger")
	testSpell(t, alphabet, "Schlacht alt", "Schule Ludwig Anton Charlotte Theodor ' ' Anton Ludwig Theodor")
}

func TestSpell_SpecialCharacter(t *testing.T) {
	testSpell(t, alphabet, "?", "'?'")
}

func TestForLanguageCode(t *testing.T) {
	type TestCase struct {
		lang     string
		expected language.Tag
	}
	var allTestCase = []TestCase{
		{"default", language.MustParse("en")},
		{"de-DE", language.MustParse("de-DE")},
		{"en", language.MustParse("en")},
		{"fr", language.MustParse("fr")},
		{"nl", language.MustParse("nl")},
		{"fr-CH", language.MustParse("fr")},
		{"zh", language.MustParse("en")},
		{"DIN 5009", language.MustParse("de-DE")},
	}

	for _, test := range allTestCase {
		t.Run(test.lang, func(t *testing.T) {
			var alphabet = Lookup(test.lang)
			if test.expected != alphabet.Lang {
				t.Error("Code", test.lang, "should return\n", test.expected, "but was\n", alphabet.Lang)
			}
		})
	}
}

func BenchmarkSpellingAlphabet_Spell(b *testing.B) {
	for i := 0; i < b.N; i++ {
		All[0].Spell("Donaudampfschiffahrtsgesellschaftskapitänsmützenspitze")
	}
}
