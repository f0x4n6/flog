// Log forensic artifacts as JSON in ECS format.
//
// Usage:
//
//	flog [-pqhv] [-D DIR] [FILE ...]
//
// The flags are:
//
//	 -D directory
//	    The log directory.
//	 -p
//		Pretty JSON.
//	 -q
//		Quiet mode.
//	 -h
//		Show usage.
//	 -v
//		Show version.
//
// The arguments are:
//
//	 file
//		The event log file(s) to process.
//		Defaults to STDIN if not given.
package main

import (
	"flag"
	"io"

	"github.com/f0x4n6/futils/pkg/sys"

	"github.com/f0x4n6/flog/internal/flog"
)

func main() {
	D := flag.String("D", "", "Log directory")
	p := flag.Bool("p", false, "Pretty JSON")
	q := flag.Bool("q", false, "Quiet mode")
	h := flag.Bool("h", false, "Show usage")
	v := flag.Bool("v", false, "Show version")

	flag.CommandLine.SetOutput(io.Discard)
	flag.Parse()

	args, _ := sys.Args()

	files := flog.StripHash(args)

	if *v {
		sys.Final("flog", flog.Version)
	}

	if *h || len(files) == 0 {
		sys.Usage("flog [-pqhv] [-D DIR] [FILE ...]")
	}

	if *q {
		sys.Progress = nil
	}

	err := flog.Log(files, *D, *p)

	if err != nil {
		sys.Fatal(err)
	}
}
