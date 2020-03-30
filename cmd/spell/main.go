package main

import (
	"flag"
	"fmt"
	"github.com/simonnagl/spell/alphabet"
	"golang.org/x/text/language/display"
	"sort"
	"strings"
)

func main() {
	flag.CommandLine.Usage = printUsage
	printHelp, lang := DefineFlags()

	flag.Parse()

	if *printHelp || nothingToSpell() {
		printUsage()
		return
	}

	a := alphabet.ForLanguageCode(*lang)
	args := strings.Join(flag.Args(), " ")

	fmt.Println(a.Spell(args))
}

func DefineFlags() (printHelp *bool, lang *string) {
	lang = flag.String("l", "en", "Spelling `alphabet` to use")
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

func printUsage() {
	fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s \n\nOptions:\n", synopsis())
	flag.CommandLine.PrintDefaults()
	fmt.Fprintf(flag.CommandLine.Output(), "\nSpelling alphabets:\n")
	printAlphabets()
}

func printAlphabets() {
	type Alphabet struct {
		tag         string
		englishName string
		selfName    string
	}

	result := make([]Alphabet, 0, len(alphabet.Lang))
	var displayEnglish = display.English.Tags()

	for k := range alphabet.Lang {
		result = append(result, Alphabet{k.String(), displayEnglish.Name(k), display.Self.Name(k)})
	}

	sort.Slice(result, func(i int, j int) bool {
		return result[i].tag < result[j].tag
	})

	var maxKeyLen, maxEnglisLen int
	for _, f := range result {
		if maxKeyLen < len(f.tag) {
			maxKeyLen = len(f.tag)
		}
		if maxEnglisLen < len(f.englishName) {
			maxEnglisLen = len(f.englishName)
		}
	}

	for _, f := range result {
		fmt.Fprintf(flag.CommandLine.Output(), "  %-*v%-*v%s\n", maxKeyLen+1, f.tag, maxEnglisLen+1, f.englishName, f.selfName)
	}
}
