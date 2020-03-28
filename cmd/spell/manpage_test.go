package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestManpage(t *testing.T) {
	commitedManpage := readCommitedManpage(t)
	generatedManpage := generateManpage()

	if commitedManpage != generatedManpage {
		writeGeneratedManpage(t, generatedManpage)

		t.Error("The manpage generated by this test does not match the commited manpage of this project.",
			"The manpage is generated to reduce duplicate information.",
			"The manpage is also commited because we include parts of it into the README.",
			"Fix the generated manpage or adopt the commited manpage to fix this test.")
	}
}

func readCommitedManpage(t *testing.T) string {
	const commitedManpageName = "../../man-page.adoc"

	commitedManpage, err := ioutil.ReadFile(commitedManpageName)
	if err != nil {
		t.Fatal("Could not read file", commitedManpageName, err)
	}
	return string(commitedManpage)
}

func generateManpage() string {
	flags := flag.CommandLine
	defer func() {
		flag.CommandLine = flags
	}()
	flag.CommandLine = &flag.FlagSet{}
	DefineFlags()

	return fmt.Sprintf(`= spell(1)
Simon Nagl
v0.1.0
:doctype: manpage

== Name

spell - spell word(s) using a spelling alphabet.

== Synopsis

%s

== Examples

To set a default language you may use an alias:

	alias spell="spell -l de"

== Options

%s
== Copyright

Copyright (C) 2020 Simon Nagl. +
Free use of this software is granted under the terms of the MIT License.
`,
		synopsis(),
		options())
}

func options() string {
	buf := bytes.Buffer{}
	flag.VisitAll(func(f *flag.Flag) {
		fType, fUsage := flag.UnquoteUsage(f)
		buf.WriteString(fmt.Sprintf("*-%s* %s:: %s (Default: %s)\n", f.Name, fType, fUsage, f.DefValue))
	})
	return buf.String()
}

func writeGeneratedManpage(t *testing.T, generatedManpage string) {
	fName := "man-page.test.adoc"
	f, err := os.Create(fName)
	if err != nil {
		t.Fatal("Could not create file", f, err)
	}
	_, err = f.WriteString(generatedManpage)
	if err != nil {
		t.Fatal("Could not write to file", f, err)
	}
}

func TestReadme(t *testing.T) {
	const commitedReadmeName = "../../README.adoc"

	commitedReadme, err := ioutil.ReadFile(commitedReadmeName)
	if err != nil {
		t.Fatal("Could not read file", commitedReadmeName, err)
	}

	flags := flag.CommandLine
	defer func() {
		flag.CommandLine = flags
	}()
	flag.CommandLine = &flag.FlagSet{}
	DefineFlags()

	generatedManpage := fmt.Sprintf(`= spell

image:https://github.com/simonnagl/spell/workflows/Go/badge.svg[Go,link=https://github.com/simonnagl/spell/actions?query=branch:master]
image:https://coveralls.io/repos/github/simonnagl/spell/badge.svg?branch=master&t=47TqXT[Coverage Status,link=https://coveralls.io/github/simonnagl/spell?branch=master]

spell word(s) using a spelling alphabet.

== Synopsis

	%s

== Examples

To set a default language you may use an alias:

	alias spell="spell -l de"

== Options

%s
== Copyright

Copyright (C) 2020 Simon Nagl. +
Free use of this software is granted under the terms of the MIT License.
`,
		synopsis(),
		options())

	if string(commitedReadme) != generatedManpage {
		fName := "readme.test.adoc"
		f, err := os.Create(fName)
		if err != nil {
			t.Fatal("Could not create file", f, err)
		}
		_, err = f.WriteString(generatedManpage)
		if err != nil {
			t.Fatal("Could not write to file", f, err)
		}
		t.Error("The README generated by this test does not match the commited REAMDME of this project.",
			"The README is generated to reduce duplicate information.",
			"The README is also commited because we want to make it readable.",
			"Fix the generated README or adopt the commited README to fix this test.")
	}
}
