package alphabet

import (
	"reflect"
	"strings"
	"testing"
)

var alphabet = SpellingAlphabet{
	"a":   "Anton",
	"ä":   "Ärger",
	"l":   "Ludwig",
	"sch": "Schule",
	"ch":  "Charlotte",
	"t":   "Theodor",
}

func testSpell(t *testing.T, alphabet SpellingAlphabet, allInputLetters string, expectedResult string) bool {
	return t.Run("Spell "+allInputLetters, func(t *testing.T) {
		result := alphabet.Spell(allInputLetters)
		if expectedResult != result {
			t.Errorf("'%s' should be spelled as \n'%s', but was \n'%s'", allInputLetters, expectedResult, result)
		}
	})
}

func TestSpell_Alphabet(t *testing.T) {
	for inputLetter, expectedResult := range alphabet {
		testSpell(t, alphabet, inputLetter, expectedResult)
		testSpell(t, alphabet, strings.Title(inputLetter), expectedResult)
	}
}

func TestSpell_Lang(t *testing.T) {
	for key, alphabet := range Lang {
		t.Run(key, func(t *testing.T) {
			for inputLetter, expectedResult := range alphabet {
				testSpell(t, alphabet, inputLetter, expectedResult)
				testSpell(t, alphabet, strings.Title(inputLetter), expectedResult)
			}
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
		{"default", Lang["en"]},
		{"de", Lang["de"]},
		{"en", Lang["en"]},
		{"fr", Lang["fr"]},
		{"nl", Lang["nl"]},
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
