package main

import (
	"bytes"
	"github.com/simonnagl/spell/test"
	"io"
	"os"
	"strings"
	"testing"
)

func TestMain_EmtpyArgs(t *testing.T) {
	output, err := captureOutput(main)
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(output, "Usage") {
		t.Error("main() without args should print Usage, not", output)
	}
}

func testMain(t *testing.T, arg string, expected string) {
	cleanup := test.ClearCommandLine()
	defer cleanup()

	os.Args = append(os.Args, arg)

	o, err := captureOutput(main)
	if err != nil {
		t.Fatal("Could not capture output of main().", err)
	}

	if expected != o {
		t.Errorf("Expected usage does not match.\ngot:\n%s\nwant:\n%s", o, expected)
	}
}

func captureOutput(f func()) (string, error) {
	r, w, err := os.Pipe()
	if err != nil {
		return "", err
	}

	var stderr = os.Stderr
	os.Stderr = w
	defer func() {
		os.Stderr = stderr
	}()

	f()

	_ = w.Close()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, r)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func TestMain_Usage(t *testing.T) {
	e := `Usage: spell [-hlv] <word(s)> 

Options:
  -h	Print this usage note
  -l alphabet
    	Spelling alphabet to use (default "en")
  -v	Print version info

Spelling alphabets:
  cs    Czech
  da    Danish
  de-AT Austrian German, ÖNORM A 1081
  de-CH Swiss High German
  de-DE German (Germany), DIN 5009
  en    English, ICAO, NATO
  en-GB British English
  es    Spanish
  fi    Finnish
  fr    French
  it    Italian
  nl    Dutch
  no    Norwegian Bokmål
  pt-BR Brazilian Portuguese
  pt-PT European Portuguese
  ro    Romanian
  sl    Slovenian
  sv    Swedish
  tr    Turkish
`
	testMain(t, "-h", e)
}

func TestMain_Version(t *testing.T) {
	testMain(t, "-v", "spell 0.2.0\n")
}
