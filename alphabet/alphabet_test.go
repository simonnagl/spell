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
		t.Run(a.lang.String(), func(t *testing.T) {
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
		lang               string
		expected           language.Tag
		expectedConfidence Exactness
	}
	var allTestCase = []TestCase{
		{"default", language.MustParse("en"), Default},
		{"de-DE", language.MustParse("de-DE"), Exact},
		{"en", language.MustParse("en"), Exact},
		{"fr", language.MustParse("fr"), Exact},
		{"nl", language.MustParse("nl"), Exact},
		{"fr-CH", language.MustParse("fr"), Guess},
		{"zh", language.MustParse("en"), Default},
		{"ru", language.MustParse("ru"), Exact},
		{"uk", language.MustParse("uk"), Exact},
	}

	for _, test := range allTestCase {
		t.Run(test.lang, func(t *testing.T) {
			alphabet, confidence := Lookup(test.lang)
			if test.expected != alphabet.lang {
				t.Error("Code", test.lang, "should return\n", test.expected, "but was\n", alphabet.lang)
			}
			if test.expectedConfidence != confidence {
				t.Error("Code", test.lang, "should return", test.expected,
					"with", test.expectedConfidence, "confidence, but confidence was", confidence)
			}
		})
	}
}

func BenchmarkSpellingAlphabet_Spell(b *testing.B) {
	for i := 0; i < b.N; i++ {
		All[0].Spell("Donaudampfschiffahrtsgesellschaftskapitänsmützenspitze")
	}
}
