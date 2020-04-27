package alphabet

import (
	"golang.org/x/text/language"
	"testing"
)

func TestExactness_String(t *testing.T) {
	tests := []struct {
		c    Exactness
		want string
	}{
		{Default, "Default"},
		{Guess, "Guess"},
		{Exact, "Exact"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := tt.c.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExactness_FromLangConfidence(t *testing.T) {
	tests := []struct {
		c    language.Confidence
		want Exactness
	}{
		{language.No, Default},
		{language.Low, Guess},
		{language.High, Guess},
		{language.Exact, Exact},
	}
	for _, tt := range tests {
		t.Run(tt.c.String(), func(t *testing.T) {
			if got := FromLangConfidence(tt.c); got != tt.want {
				t.Errorf("FromLangConfidence() = %v, want %v", got, tt.want)
			}
		})
	}
}
