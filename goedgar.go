package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/coleaeason/goedgar/dates"
)

// Globals
var (
	// CLI Flags
	flagDate  = flag.String("date", "2023-11-01", "Date to check")
	flagDebug = flag.Bool("debug", false, "Print some debug information")
)

func main() {
	// Create a usage message that contains examples instead of just default args.
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Example usages:\n")
		fmt.Fprintf(os.Stderr, "  Generate a default, valid token:\n")
		fmt.Fprintf(os.Stderr, "    jwtgen\n")
		fmt.Fprintf(os.Stderr, "  Generate a default, valid token, and pretty-print debug information:\n")
		fmt.Fprintf(os.Stderr, "    jwtgen --debug -pp\n")
		fmt.Fprintf(os.Stderr, "  Generate an expired token for cole@test.com:\n")
		fmt.Fprintf(os.Stderr, "    jwtgen --expired --email=cole@test.com\n")
		fmt.Fprintf(os.Stderr, "\nOptions: \n")

		flag.PrintDefaults()
	}

	// First things first, parse our CLI flags.
	flag.Parse()

	result, err := dates.IsSECHoliday(*flagDate)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println(*flagDate, "is a holiday?", result)
}
