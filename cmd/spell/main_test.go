package main

import (
	"bytes"
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
