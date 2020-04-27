package alphabet

import (
	"fmt"
	"golang.org/x/text/language"
)

// Exactness indicates the level of certainty for a given return value.
// The confidence level indicates whether a value was explicitly specified,
// whether it is a educated guess, or whether it is the default value.
type Exactness int

const (
	Default Exactness = iota // full confidence that there was no match
	Guess                    // most likely value picked out of a set of alternatives
	Exact                    // exact match or explicitly specified value
)

var confName = []string{"Default", "Guess", "Exact"}

func (c Exactness) String() string {
	return confName[c]
}

func FromLangConfidence(c language.Confidence) Exactness {
	switch c {
	case language.No:
		return Default
	case language.Low, language.High:
		return Guess
	case language.Exact:
		return Exact
	}
	panic(fmt.Sprint("FromLangConfidence not implemented for Confidence", c))
}
