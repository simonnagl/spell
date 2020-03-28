package main

import (
	"flag"
	"fmt"
	"github.com/simonnagl/spell/alphabet"
	"sort"
	"strings"
)

func main() {
	flag.CommandLine.Usage = usage
	printHelp, lang := DefineFlags()

	flag.Parse()

	if *printHelp || nothingToSpell() {
		usage()
		return
	}

	a := alphabet.ForLanguageCode(*lang)
	args := strings.Join(flag.Args(), " ")

	fmt.Println(a.Spell(args))
}

func DefineFlags() (printHelp *bool, lang *string) {
	lang = flag.String("l", "en", "Spelling alphabet to use")
	printHelp = flag.Bool("h", false, "Print this usage note")
	return
}

func nothingToSpell() bool {
	return len(flag.Args()) == 0
}

func synopsis() string {
	var allName string
	flag.VisitAll(func(flag *flag.Flag) {
		allName += flag.Name
	})

	return fmt.Sprintf("spell [-%s] <word(s)>", allName)
}

func usage() {
	fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s \n\nOptions:\n", synopsis())
	flag.CommandLine.PrintDefaults()
	fmt.Fprintf(flag.CommandLine.Output(), "\nSpelling alphabets:\n  %s\n", alphabets())
}

func alphabets() string {
	keys := make([]string, 0, len(alphabet.Lang))

	for k := range alphabet.Lang {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return strings.Join(keys, ", ")
}
