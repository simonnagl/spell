// Package test provides internal helper functions for tests. Clients should not use this internal package.
package test

import (
	"flag"
	"os"
)

// ClearCommandLine clears flag.CommandLine and os.Args
//
// os.Args is set to []string{"spell"}.
// flag.CommandLine is set to an empty FlagSet
//
// Call reset to restore flag.CommandLine and os.Args to their state before calling this function.
func ClearCommandLine() (reset func()) {
	flags := flag.CommandLine
	flag.CommandLine = &flag.FlagSet{}
	args := os.Args
	os.Args = []string{"spell"}
	return func() {
		flag.CommandLine = flags
		os.Args = args
	}
}
