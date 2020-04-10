package main

import (
	"flag"
	"fmt"
	"github.com/simonnagl/spell/alphabet"
	"golang.org/x/text/language/display"
	"sort"
	"strings"
)

var (
	printHelp    *bool
	printVersion *bool
	lang         *string
)

func main() {
	flag.CommandLine.Usage = printUsage
	DefineFlags()

	flag.Parse()

	if *printHelp {
		printUsage()
		return
	}
	if *printVersion {
		fmt.Fprintln(flag.CommandLine.Output(), "spell", Version)
		return
	}
	if nothingToSpell() {
		printUsage()
		return
	}

	a := alphabet.ForLanguageCode(*lang)
	args := strings.Join(flag.Args(), " ")

	fmt.Println(a.Spell(args))
}

func DefineFlags() {
	lang = flag.String("l", "en", "Spelling `alphabet` to use")
	printHelp = flag.Bool("h", false, "Print this usage note")
	printVersion = flag.Bool("v", false, "Print version info")
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

type alphabetView struct {
	tag         string
	englishName string
	selfName    string
}

func alphabetViewModel() []alphabetView {

	allAlphabetView := make([]alphabetView, 0, len(alphabet.List))
	var displayEnglish = display.English.Tags()

	for _, a := range alphabet.List {
		lang := a.Lang
		allAlphabetView = append(allAlphabetView, alphabetView{lang.String(), displayEnglish.Name(lang), display.Self.Name(lang)})
	}

	sort.Slice(allAlphabetView, func(i int, j int) bool {
		return allAlphabetView[i].tag < allAlphabetView[j].tag
	})

	return allAlphabetView
}

func printAlphabets() {

	allAlphabet := alphabetViewModel()

	var maxKeyLen, maxEnglisLen int
	for _, f := range allAlphabet {
		if maxKeyLen < len(f.tag) {
			maxKeyLen = len(f.tag)
		}
		if maxEnglisLen < len(f.englishName) {
			maxEnglisLen = len(f.englishName)
		}
	}

	for _, f := range allAlphabet {
		fmt.Fprintf(flag.CommandLine.Output(), "  %-*v%-*v%s\n", maxKeyLen+1, f.tag, maxEnglisLen+1, f.englishName, f.selfName)
	}
}
