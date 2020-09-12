// lsr lists directories recursively
//
// Usage:
//		lsr [-d | -f ] [ n ...]
//
// For each directory argument, lsr recursively lists the contents of the
// directory; for each file argument, lsr repeats its name. When no argument is
// given, the current directory is listed.
//
// Options:
//		-d print directories
//		-f print files
//
// TODO: print resolved, absolute paths instead. For example,
// lsr -d ../ | xargs realpath
//
// Output:
// /home/bwasd/go/src/github.com/bwasd/p9/cmd/fortune
// /home/bwasd/go/src/github.com/bwasd/p9/cmd/2fa
// /home/bwasd/go/src/github.com/bwasd/p9/cmd/lsr
// ...
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var (
	flagD = flag.Bool("d", false, "print directories")
	flagF = flag.Bool("f", false, "print files")
)

var usageString = `usage: lsr [ -d | -f ] [name ...]
	Options:`

func usage() {
	fmt.Fprint(os.Stderr, usageString)
	flag.PrintDefaults()
	os.Exit(1)
}

func prname(path string, f os.FileInfo) error {
	if f.IsDir() && path[len(path)-1] != '/' {
		if !*flagD {
			return nil
		}
		path = path + "/"
	} else if !*flagF {
		return nil
	}
	fmt.Println(path)
	return nil
}

func pr(path string, f os.FileInfo, err error) error {
	if err != nil {
		fmt.Fprintf(os.Stderr, "lsr: %s", err)
		return nil
	}
	return prname(path, f)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if !(*flagD || *flagF) {
		*flagF = true
	}

	if flag.NArg() == 0 {
		filepath.Walk(".", pr)
		return
	}

	for _, v := range flag.Args() {
		filepath.Walk(v, pr)
	}
}
