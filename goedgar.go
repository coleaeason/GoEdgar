package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/coleaeason/goedgar/fetcher"
	"github.com/coleaeason/goedgar/filings"
)

// Globals
var (
	// CLI Flags
	flagDate  = flag.String("date", "2023-11-01", "Date to check")
	flagCIK   = flag.String("cik", "", "CIK to process")
	flagDebug = flag.Bool("debug", false, "Print some debug information")
)

func main() {
	// Create a usage message that contains examples instead of just default args.
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Example usages:\n")
		fmt.Fprintf(os.Stderr, "\nOptions: \n")

		flag.PrintDefaults()
	}

	// First things first, parse our CLI flags.
	flag.Parse()

	r := fetcher.GetPage(fmt.Sprintf("https://data.sec.gov/submissions/CIK%s.json", *flagCIK))
	b, _ := io.ReadAll(r)

	var response filings.CIKFilingJson
	err := json.Unmarshal(b, &response)
	if err != nil {
		log.Fatal("JSON Unmarshal failed ", err)
	}

	// dates.IsSECHoliday(*flagDate)
	var forms []filings.Filing
	recents := response.Filings.Recent
	for i, _ := range recents.FilingDate {
		var form filings.Filing
		form.AcceptanceDateTime = recents.AcceptanceDateTime[i]
		form.AccessionNumber = recents.AccessionNumber[i]
		form.Act = recents.Act[i]
		form.FileNumber = recents.FileNumber[i]
		form.FilingDate = recents.FilingDate[i]
		form.FilmNumber = recents.FilmNumber[i]
		form.Form = recents.Form[i]
		form.IsInlineXBRL = recents.IsInlineXBRL[i]
		form.IsXBRL = recents.IsXBRL[i]
		form.Items = recents.Items[i]
		form.PrimaryDocDescription = recents.PrimaryDocDescription[i]
		form.PrimaryDocument = recents.PrimaryDocDescription[i]
		form.ReportDate = recents.ReportDate[i]
		form.Size = recents.Size[i]
		forms = append(forms, form)
	}
	for i, _ := range forms {
		fmt.Println(forms[i])
	}
}
