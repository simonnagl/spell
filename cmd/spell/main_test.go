package main

import (
	"bytes"
	"flag"
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

func TestUsage(t *testing.T) {
	flags := flag.CommandLine
	defer func() {
		flag.CommandLine = flags
	}()
	flag.CommandLine = &flag.FlagSet{}
	DefineFlags()

	o, err := captureOutput(printUsage)
	if err != nil {
		t.Fatal(err)
	}

	e := `Usage: spell [-hl] <word(s)> 

Options:
  -h	Print this usage note
  -l alphabet
    	Spelling alphabet to use (default "en")

Spelling alphabets:
  cs    Czech                čeština
  da    Danish               dansk
  de-AT Austrian German      Österreichisches Deutsch
  de-CH Swiss High German    Schweizer Hochdeutsch
  de-DE German (Germany)     Deutsch
  en    English              English
  en-GB British English      British English
  es    Spanish              español
  fi    Finnish              suomi
  fr    French               français
  it    Italian              italiano
  nl    Dutch                Nederlands
  no    Norwegian Bokmål     norsk bokmål
  pt-BR Brazilian Portuguese português
  pt-PT European Portuguese  português europeu
  ro    Romanian             română
  sl    Slovenian            slovenščina
  sv    Swedish              svenska
  tr    Turkish              Türkçe
`
	if o != e {
		t.Errorf("Expected usage does not match.\ngot:\n%s\nwant:\n%s", o, e)
	}
}
