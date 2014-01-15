// Copyright 2014 Jongmin Kim. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package main

import (
	"flag"
	"fmt"
	"go/scanner"
	"os"

	"github.com/atomaths/fng/term"
)

var (
	exitCode = 0
	recursive bool
)

func init() {
	flag.BoolVar(&recursive, "r", false, "recursive")
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: fng [file name patterns...] [match string]\n")
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

	for i := 0; i < flag.NArg(); i++ {
		fmt.Println(flag.Arg(i))
	}

	fmt.Println(term.Width())
}

func main() {
	// call gofmtMain in a separate function
	// so that it can use defer and have them
	// run before the exit.
	gofmtMain()
	os.Exit(exitCode)
}
