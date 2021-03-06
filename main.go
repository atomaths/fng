// Copyright 2014 Jongmin Kim. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package main

import (
	"flag"
	"fmt"
	"go/scanner"
	"os"

	"github.com/atomaths/fng/pkg/term"
)

var (
	exitCode   = 0
	recursive  bool
	binary     bool
	ignoreCase bool
	word       bool
)

func init() {
	flag.BoolVar(&recursive, "r", false, "recursive directory search")
	flag.BoolVar(&binary, "b", false, "binary file search")
	flag.BoolVar(&ignoreCase, "i", false, "ignore case distinctions")
	flag.BoolVar(&word, "w", false, "force PATTERN to match only whole words")
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: fng [OPTION]... PATTERN [FILE]...\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func report(err error) {
	scanner.PrintError(os.Stderr, err)
	exitCode = 2
}

func gofmtMain() {
	flag.Usage = usage
	flag.Parse()

	f := new(finder)

	// TODO(atomaths): skip directory or not
	for i := 0; i < flag.NArg(); i++ {
		f.files = append(f.files, flag.Arg(i))
	}

	fmt.Println(f)

	fmt.Println(term.Width())
}

func main() {
	// call gofmtMain in a separate function
	// so that it can use defer and have them
	// run before the exit.
	gofmtMain()
	os.Exit(exitCode)
}
