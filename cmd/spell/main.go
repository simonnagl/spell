package main

import (
	"flag"
	"fmt"
	"github.com/simonnagl/spell/alphabet"
	"os"
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

	a, e := alphabet.Lookup(*lang)
	args := strings.Join(flag.Args(), " ")

	switch e {
	case alphabet.Guess:
		fmt.Fprintf(os.Stderr, "Info: Guess alphabet '%s' for input '%s':\n", a.LangTag(), *lang)
	case alphabet.Default:
		fmt.Fprintf(os.Stderr, "Warning: Found no spelling alphabet for '%s'. Using default '%s':\n", *lang, a.LangTag())
	}

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
	LangTag         string
	LangEnglishName string
	LangSelfName    string
	AltNames        string
}

func alphabetViewModel() []alphabetView {

	allAlphabetView := make([]alphabetView, 0, len(alphabet.All))

	for _, a := range alphabet.All {
		allAlphabetView = append(allAlphabetView, alphabetView{
			LangTag:         a.LangTag(),
			LangEnglishName: a.LangEnglishName(),
			LangSelfName:    a.LangSelfName(),
			AltNames:        strings.Join(a.Names(), ", "),
		})
	}

	sort.Slice(allAlphabetView, func(i int, j int) bool {
		return allAlphabetView[i].LangTag < allAlphabetView[j].LangTag
	})

	return allAlphabetView
}

func printAlphabets() {
	allAlphabet := alphabetViewModel()

	for _, f := range allAlphabet {
		if "" != f.AltNames {
			fmt.Fprintf(flag.CommandLine.Output(), "  %-6v%v, %v\n", f.LangTag, f.LangEnglishName, f.AltNames)
		} else {
			fmt.Fprintf(flag.CommandLine.Output(), "  %-6v%v\n", f.LangTag, f.LangEnglishName)
		}
	}
}
