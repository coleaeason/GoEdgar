package filings

import (
	"time"

	"github.com/coleaeason/goedgar/dates"
)

type CIKFilingJson struct {
	Cik                               string   `json:"cik"`
	EntityType                        string   `json:"entityType"`
	Sic                               string   `json:"sic"`
	SicDescription                    string   `json:"sicDescription"`
	InsiderTransactionForOwnerExists  int      `json:"insiderTransactionForOwnerExists"`
	InsiderTransactionForIssuerExists int      `json:"insiderTransactionForIssuerExists"`
	Name                              string   `json:"name"`
	Tickers                           []string `json:"tickers"`
	Exchanges                         []string `json:"exchanges"`
	Ein                               string   `json:"ein"`
	Description                       string   `json:"description"`
	Website                           string   `json:"website"`
	InvestorWebsite                   string   `json:"investorWebsite"`
	Category                          string   `json:"category"`
	FiscalYearEnd                     string   `json:"fiscalYearEnd"`
	StateOfIncorporation              string   `json:"stateOfIncorporation"`
	StateOfIncorporationDescription   string   `json:"stateOfIncorporationDescription"`
	Addresses                         struct {
		Mailing struct {
			Street1                   string `json:"street1"`
			Street2                   string `json:"street2"`
			City                      string `json:"city"`
			StateOrCountry            string `json:"stateOrCountry"`
			ZipCode                   string `json:"zipCode"`
			StateOrCountryDescription string `json:"stateOrCountryDescription"`
		} `json:"mailing"`
		Business struct {
			Street1                   string `json:"street1"`
			Street2                   string `json:"street2"`
			City                      string `json:"city"`
			StateOrCountry            string `json:"stateOrCountry"`
			ZipCode                   string `json:"zipCode"`
			StateOrCountryDescription string `json:"stateOrCountryDescription"`
		} `json:"business"`
	} `json:"addresses"`
	Phone       string   `json:"phone"`
	Flags       string   `json:"flags"`
	FormerNames []string `json:"formerNames"`
	Filings     struct {
		Recent Filings  `json:"recent"`
		Files  []string `json:"files"`
	} `json:"filings"`
}

type Filings struct {
	AccessionNumber       []string     `json:"accessionNumber"`
	FilingDate            []dates.Date `json:"filingDate"`
	ReportDate            []dates.Date `json:"reportDate"`
	AcceptanceDateTime    []time.Time  `json:"acceptanceDateTime"`
	Act                   []string     `json:"act"`
	Form                  []string     `json:"form"`
	FileNumber            []string     `json:"fileNumber"`
	FilmNumber            []string     `json:"filmNumber"`
	Items                 []string     `json:"items"`
	Size                  []int        `json:"size"`
	IsXBRL                []int        `json:"isXBRL"`
	IsInlineXBRL          []int        `json:"isInlineXBRL"`
	PrimaryDocument       []string     `json:"primaryDocument"`
	PrimaryDocDescription []string     `json:"primaryDocDescription"`
}

type Filing struct {
	AccessionNumber       string     `json:"accessionNumber"`
	FilingDate            dates.Date `json:"filingDate"`
	ReportDate            dates.Date `json:"reportDate"`
	AcceptanceDateTime    time.Time  `json:"acceptanceDateTime"`
	Act                   string     `json:"act"`
	Form                  string     `json:"form"`
	FileNumber            string     `json:"fileNumber"`
	FilmNumber            string     `json:"filmNumber"`
	Items                 string     `json:"items"`
	Size                  int        `json:"size"`
	IsXBRL                int        `json:"isXBRL"`
	IsInlineXBRL          int        `json:"isInlineXBRL"`
	PrimaryDocument       string     `json:"primaryDocument"`
	PrimaryDocDescription string     `json:"primaryDocDescription"`
}
