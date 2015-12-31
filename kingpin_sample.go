package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	verbose = kingpin.Flag("verbose", "Set verbose mode").Short('v').OverrideDefaultFromEnvar("KINGPIN_SAMPLE_VERBOSE").Bool()
	count   = kingpin.Flag("count", "counter").Default("10").Int()
	name    = kingpin.Arg("name", "Input name").Required().String()
)

func main() {
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	fmt.Printf("verbose mode: %v\n", *verbose)
	fmt.Printf("count       : %d\n", *count)
	fmt.Printf("name        : %s\n", *name)
}

