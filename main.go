// lsr lists directories recursively
//
// For each directory argument, lsr recursively lists the directory's contents;
// for each file argument, lsr repeats its name. Given no arguments, the current
// directory is listed.
//
// Usage:
//
//	lsr [-d] [-f] [name ...]
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
)

var (
	flagD = flag.Bool("d", false, "print directories")
	flagF = flag.Bool("f", false, "print files")
)

var usageString = `usage: lsr [-d] [-f] [name ...]
options:
`

func usage() {
	fmt.Fprint(os.Stderr, usageString)
	flag.PrintDefaults()
}

func prname(p string, f os.FileInfo) error {
	if f.IsDir() && p[len(p)-1] != '/' {
		if !*flagD {
			return nil
		}
		p += "/"
	} else if !*flagF {
		return nil
	}
	fmt.Println(path.Clean(p))
	return nil
}

func pr(p string, f os.FileInfo, err error) error {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil
	}
	return prname(p, f)
}

func main() {
	log.SetPrefix("lsr: ")
	flag.Usage = usage
	flag.Parse()

	if !(*flagD || *flagF) {
		*flagF = true
	}

	if flag.NArg() == 0 {
		if err := filepath.Walk(".", pr); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}

	for i := range flag.Args() {
		if err := filepath.Walk(flag.Arg(i), pr); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
