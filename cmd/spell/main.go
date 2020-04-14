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
	Tag         string
	EnglishName string
	SelfName    string
}

func alphabetViewModel() []alphabetView {

	allAlphabetView := make([]alphabetView, 0, len(alphabet.All))
	var displayEnglish = display.English.Tags()

	for _, a := range alphabet.All {
		lang := a.Lang
		allAlphabetView = append(allAlphabetView, alphabetView{lang.String(), displayEnglish.Name(lang), display.Self.Name(lang)})
	}

	sort.Slice(allAlphabetView, func(i int, j int) bool {
		return allAlphabetView[i].Tag < allAlphabetView[j].Tag
	})

	return allAlphabetView
}

func printAlphabets() {

	allAlphabet := alphabetViewModel()

	var maxKeyLen, maxEnglisLen int
	for _, f := range allAlphabet {
		if maxKeyLen < len(f.Tag) {
			maxKeyLen = len(f.Tag)
		}
		if maxEnglisLen < len(f.EnglishName) {
			maxEnglisLen = len(f.EnglishName)
		}
	}

	for _, f := range allAlphabet {
		fmt.Fprintf(flag.CommandLine.Output(), "  %-*v%-*v%s\n", maxKeyLen+1, f.Tag, maxEnglisLen+1, f.EnglishName, f.SelfName)
	}
}
