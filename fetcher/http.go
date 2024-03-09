package fetcher

import (
	"io"
	"log"
	"net/http"
)

// Gets a page from SEC.gov. Injects the required User-Agent header
// so that the request does not fail.
func GetPage(url string) io.ReadCloser {
	client := http.Client{}
	req, _ := http.NewRequest("GET", url, nil)

	req.Header = http.Header{
		"User-Agent": {"GoEdgar cole@cole.dev"},
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Query to get page ", url, "failed: ", err)
		return nil
	}
	return resp.Body
}
