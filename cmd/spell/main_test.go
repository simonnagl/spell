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

	o, err := captureOutput(usage)
	if err != nil {
		t.Fatal(err)
	}

	e := `Usage: spell [-hl] <word(s)> 

Options:
  -h	Print this usage note
  -l string
    	Spelling alphabet to use (default "en")

Spelling alphabets:
  at, ch, de, en, fr, it, nl
`
	if o != e {
		t.Errorf("Expected usage does not match.\ngot:\n%s\nwant:\n%s", o, e)
	}
}
