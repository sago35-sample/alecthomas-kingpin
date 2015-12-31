package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	verbose = kingpin.Flag("verbose", "Set verbose mode").Bool()
	count   = kingpin.Flag("count", "counter").Int()
	name    = kingpin.Arg("name", "Input name").Required().String()
)

func main() {
	kingpin.Parse()
	fmt.Printf("verbose mode: %v\n", *verbose)
	fmt.Printf("count       : %d\n", *count)
	fmt.Printf("name        : %s\n", *name)
}

