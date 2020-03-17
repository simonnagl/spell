package main

import (
	"strings"
	"testing"
)

func testSpell(t *testing.T, allInputLetters string, expectedResult string) bool {
	return t.Run("Spell "+allInputLetters, func(t *testing.T) {
		result := DIN5009.spell(allInputLetters)
		if expectedResult != result {
			t.Errorf("'%s' should be spelled as \n'%s', but was \n'%s'", allInputLetters, expectedResult, result)
		}
	})
}

func TestSpell_Alphabet(t *testing.T) {
	for inputLetter, expectedResult := range DIN5009 {
		testSpell(t, inputLetter, expectedResult)
		testSpell(t, strings.Title(inputLetter), expectedResult)
	}
}

func TestSpell_Words(t *testing.T) {
	testSpell(t, "es", "Emil Samuel")
	testSpell(t, "Simon", "Samuel Ida Martha Otto Nordpol")
	testSpell(t, "SCHULE", "Schule Ulrich Ludwig Emil")
	testSpell(t, "Der Satz", "Dora Emil Richard Leerzeichen Samuel Anton Theodor Zacharias")
}

func TestSpell_SpecialCharacter(t *testing.T) {
	testSpell(t, "?", "?")
}
