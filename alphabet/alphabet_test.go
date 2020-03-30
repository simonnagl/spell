package alphabet

import (
	"golang.org/x/text/language"
	"reflect"
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
	for key, alphabet := range Lang {
		t.Run(key.String(), func(t *testing.T) {
			testSpellAlphabet(t, alphabet)
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
		expected SpellingAlphabet
	}
	var allTestCase = []TestCase{
		{"default", Lang[language.MustParse("en")]},
		{"de-DE", Lang[language.MustParse("de-DE")]},
		{"en", Lang[language.MustParse("en")]},
		{"fr", Lang[language.MustParse("fr")]},
		{"nl", Lang[language.MustParse("nl")]},
	}

	for _, test := range allTestCase {
		t.Run(test.lang, func(t *testing.T) {
			var alphabet = ForLanguageCode(test.lang)
			if !reflect.DeepEqual(test.expected, alphabet) {
				t.Error("Code", test.lang, "should return\n", test.expected, "but was\n", alphabet)
			}
		})
	}

}
