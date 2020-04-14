package main

import (
	"bytes"
	"flag"
	"github.com/simonnagl/spell/test"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"text/template"
)

const manpageTmpl = `= spell(1)
Simon Nagl
v0.1.0
:doctype: manpage

== Name

spell - spell word(s) using a spelling alphabet.

== Synopsis

{{ .Synopsis }}

== Options
{{ range .Options }}
*-{{ .Name }}* {{ .Type }}:: {{ .Usage }} (Default: {{ .DefValue }}){{ end }}

== Spelling alphabets
{{ range .Alphabets}}
*{{ .Tag }}* :: {{ .EnglishName }} -- {{ .SelfName }}{{ end }}

== Examples

To set a default language you may use an alias:

	alias spell="spell -l de"

== Copyright

Copyright (C) 2020 Simon Nagl. +
Free use of this software is granted under the terms of the MIT License.
`

const readmeTmpl = `= spell

image:https://github.com/simonnagl/spell/workflows/Go/badge.svg[Go,link=https://github.com/simonnagl/spell/actions?query=branch:master]
image:https://coveralls.io/repos/github/simonnagl/spell/badge.svg?branch=master&t=47TqXT[Coverage Status,link=https://coveralls.io/github/simonnagl/spell?branch=master]

spell word(s) using a spelling alphabet.

== Installation

	go install github.com/simonnagl/spell/cmd/spell

== Synopsis

	{{ .Synopsis }}

== Options
{{ range .Options }}
*-{{ .Name }}* {{ .Type }}:: {{ .Usage }} (Default: {{ .DefValue }}){{ end }}

== Spelling alphabets

[cols="h,2*"]
|===
{{ range .Alphabets}}
| {{ .Tag }} | {{ .EnglishName }} | {{ .SelfName }}{{ end }}

|===

== Examples

To set a default language you may use an alias:

	alias spell="spell -l de"

== Copyright

Copyright (C) 2020 Simon Nagl. +
Free use of this software is granted under the terms of the MIT License.
`

const godocTmpl = `// Spell is a tool to spell word(s) using a spelling alphabet.
//
// Usage:
//     {{ .Synopsis }}
// Options:{{ range .Options }}
//     -{{ .Name }}={{ .DefValue }}
//     	{{ .Usage }}{{ end }}
// Spelling alphabets:{{ range .Alphabets }}
//     {{ printf "%-8v" .Tag }}{{ .EnglishName }}{{end}}
package main
`

func TestManpage(t *testing.T) {
	testDoc(t, "../../README.adoc", readmeTmpl)
}

func TestReadme(t *testing.T) {
	testDoc(t, "../../man-page.adoc", manpageTmpl)
}

func TestGodoc(t *testing.T) {
	testDoc(t, "doc.go", godocTmpl)
}

func testDoc(t *testing.T, path string, tmpl string) {
	committedDoc := readDoc(t, path)
	generatedDoc := genDoc(t, tmpl)

	generatedFile := path + ".generated"
	if committedDoc != generatedDoc {
		writeDoc(t, generatedDoc, generatedFile)

		name := filepath.Base(path)
		t.Error("The generated", name, "does not match the committed one.",
			"The", name, "is committed to make it readable.",
			"The", name, "is generated to reduce duplicate information.",
			"This test finds regressions.",
			"Adopt either the generation or the committed file to fix this test.")
	} else {
		err := os.Remove(generatedFile)
		if err != nil && !os.IsNotExist(err) {
			t.Fatal(err)
		}
	}
}

func readDoc(t *testing.T, path string) string {
	commitedManpage, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal("Could not read file", path, err)
	}
	return string(commitedManpage)
}

func genDoc(t *testing.T, tmplString string) string {

	tmpl, err := template.New("doc").Parse(tmplString)
	if err != nil {
		t.Fatal("Could not parse template.", err)
	}

	data := data()

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		t.Fatal("Could not execute template.", err)
	}

	return buf.String()
}

type Data struct {
	Synopsis  string
	Options   []Flag
	Alphabets []alphabetView
}

func data() Data {
	cleanup := test.ClearCommandLine()
	defer cleanup()
	DefineFlags()

	return Data{
		Synopsis:  synopsis(),
		Options:   flags(),
		Alphabets: alphabetViewModel(),
	}
}

type Flag struct {
	Name     string
	Type     string
	Usage    string
	DefValue string
}

func flags() []Flag {
	var r []Flag
	flag.VisitAll(func(f *flag.Flag) {
		fType, fUsage := flag.UnquoteUsage(f)
		r = append(r, Flag{f.Name, fType, fUsage, f.DefValue})
	})
	return r
}

func writeDoc(t *testing.T, generatedManpage string, fName string) {
	f, err := os.Create(fName)
	if err != nil {
		t.Fatal("Could not create file", f, err)
	}
	_, err = f.WriteString(generatedManpage)
	if err != nil {
		t.Fatal("Could not write to file", f, err)
	}
}
